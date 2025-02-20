/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package warehouse

import (
	"context"
	"errors"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/equality"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	v1alpha1 "github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1"
	"github.com/databendcloud/databend-operator/pkg/common"
	databendruntime "github.com/databendcloud/databend-operator/pkg/runtime"
)

type opState int

const (
	createSucceeded opState = iota
	running         opState = iota
	buildFailed     opState = iota
	runFailed       opState = iota
	updateFailed    opState = iota
)

// WarehouseReconciler reconciles a Warehouse object
type WarehouseReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=databendlabs.io,resources=warehouses,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=databendlabs.io,resources=warehouses/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=databendlabs.io,resources=warehouses/finalizers,verbs=update
// +kubebuilder:rbac:groups=databendlabs.io,resources=tenants,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=secrets,verbs=get;list
// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apps,resources=statefulsets,verbs=get;list;watch;create;update;patch;delete

func (r *WarehouseReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := ctrl.LoggerFrom(ctx)

	var warehouse v1alpha1.Warehouse
	if err := r.Get(ctx, req.NamespacedName, &warehouse); err != nil {
		if apierrors.IsNotFound(err) {
			log.V(2).Info("Warehouse has been deleted", "namespacedName", req.NamespacedName)
		}
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	log = log.WithValues("warehouse", klog.KObj(&warehouse))
	ctx = ctrl.LoggerInto(ctx, log)
	log.V(2).Info("Reconciling Warehouse")

	// Get corresponding tenant and retrieve configurations
	tenantNN := types.NamespacedName{
		Namespace: warehouse.Namespace,
		Name:      warehouse.Spec.Tenant.Name,
	}
	tenant, err := r.getTenant(ctx, tenantNN)
	if err != nil {
		log.V(5).Error(err, "Failed to get tenant", "namespacedName", tenantNN)
		setCondition(&warehouse, buildFailed)
		return ctrl.Result{}, errors.Join(err, r.Status().Update(ctx, &warehouse))
	}

	originStatus := warehouse.Status.DeepCopy()

	// Reconcile ConfigMap
	cmOpState, err := r.reconcileConfigMap(ctx, tenant, &warehouse)
	setCondition(&warehouse, cmOpState)
	if err != nil {
		return ctrl.Result{}, errors.Join(err, r.Status().Update(ctx, &warehouse))
	}

	// Reconcile StatefulSet
	ssOpState, err := r.reconcileStatefulSet(ctx, tenant, &warehouse)
	setCondition(&warehouse, ssOpState)
	if err != nil {
		return ctrl.Result{}, errors.Join(err, r.Status().Update(ctx, &warehouse))
	}

	if !equality.Semantic.DeepEqual(warehouse.Status, originStatus) {
		return ctrl.Result{}, r.Status().Update(ctx, &warehouse)
	}

	return ctrl.Result{}, err
}

func (r *WarehouseReconciler) reconcileConfigMap(ctx context.Context, tenant *v1alpha1.Tenant, warehouse *v1alpha1.Warehouse) (opState, error) {
	log := ctrl.LoggerFrom(ctx)

	// Build and reconcile ConfigMap
	cm, err := databendruntime.BuildQueryConfigMap(tenant, warehouse)
	if err != nil {
		log.V(5).Error(err, "Failed to build ConfigMap", "namespace", cm.Namespace, "name", cm.Name)
		return buildFailed, err
	}

	creationErr := r.Create(ctx, cm)
	if creationErr == nil {
		log.V(5).Info("Succeeded to create ConfigMap", "namespace", cm.Namespace, "name", cm.Name)
	} else if client.IgnoreAlreadyExists(creationErr) != nil {
		log.V(5).Error(err, "Failed to create ConfigMap", "namespace", cm.Namespace, "name", cm.Name)
		return buildFailed, creationErr
	} else {
		if err := r.Update(ctx, cm); err != nil {
			return updateFailed, err
		}
		log.V(5).Info("Succeeded to update ConfigMap", "namespace", cm.Namespace, "name", cm.Name)
	}

	return createSucceeded, nil
}

func (r *WarehouseReconciler) reconcileService(ctx context.Context, tenant *v1alpha1.Tenant, warehouse *v1alpha1.Warehouse) (opState, error) {
	return createSucceeded, nil
}

func (r *WarehouseReconciler) reconcileStatefulSet(ctx context.Context, tenant *v1alpha1.Tenant, warehouse *v1alpha1.Warehouse) (opState, error) {
	log := ctrl.LoggerFrom(ctx)

	// Build and reconcile StatefulSet
	ss, err := databendruntime.BuildQueryStatefulSet(tenant, warehouse)
	if err != nil {
		log.V(5).Error(err, "Failed to build StatefulSet", "namespace", ss.Namespace, "name", ss.Name)
		return buildFailed, err
	}
	// log.V(5).Info("New Statefulset", "ss", *ss)

	// Non-empty resourceVersion indicates UPDATE operation.
	creationErr := r.Create(ctx, ss)
	switch {
	case creationErr == nil:
		log.V(5).Info("Succeeded to create StatefulSet", "namespace", ss.Namespace, "name", ss.Name)
	case client.IgnoreAlreadyExists(creationErr) != nil:
		log.V(5).Error(err, "Failed to create StatefulSet", "namespace", ss.Namespace, "name", ss.Name)
		return runFailed, creationErr
	default:
		if err := r.Update(ctx, ss); err != nil {
			return updateFailed, err
		}
		log.V(5).Info("Succeeded to update StatefulSet", "namespace", ss.Namespace, "name", ss.Name)
	}

	return r.updateReplicas(ctx, warehouse)
}

func (r *WarehouseReconciler) updateReplicas(ctx context.Context, warehouse *v1alpha1.Warehouse) (opState, error) {
	log := ctrl.LoggerFrom(ctx)

	var ss appsv1.StatefulSet
	ssNN := types.NamespacedName{
		Namespace: warehouse.Namespace,
		Name:      common.GetQueryStatefulSetName(warehouse.Spec.Tenant.Name, warehouse.Name),
	}
	if err := r.Get(ctx, ssNN, &ss); err != nil {
		if apierrors.IsNotFound(err) {
			log.V(5).Info("StatefulSet has not been created yet", "namespacedName", ssNN)
			return createSucceeded, nil
		}
		log.V(5).Error(err, "Failed to get StatefulSet", "namespacedName", ssNN)
		return buildFailed, err
	}

	warehouse.Status.ReadyReplicas = int(ss.Status.ReadyReplicas)
	if ss.Status.ReadyReplicas == *ss.Spec.Replicas {
		return running, nil
	}

	return createSucceeded, nil
}

func (r *WarehouseReconciler) getTenant(ctx context.Context, nn types.NamespacedName) (*v1alpha1.Tenant, error) {
	log := ctrl.LoggerFrom(ctx)

	var tenant v1alpha1.Tenant
	if err := r.Get(ctx, nn, &tenant, &client.GetOptions{}); err != nil {
		return nil, err
	}

	// Retrieve storage configurations
	s3Config := tenant.Spec.Storage.S3
	if s3Config.S3Auth.SecretRef != nil {
		log.V(5).Info("Getting credentials from Secret")
		var secret corev1.Secret
		nn := types.NamespacedName{
			Namespace: s3Config.S3Auth.SecretRef.Namespace,
			Name:      s3Config.S3Auth.SecretRef.Name,
		}
		if err := r.Get(ctx, nn, &secret, &client.GetOptions{}); err != nil {
			return nil, fmt.Errorf("failed to get secret %v", nn)
		}
		tenant.Spec.Storage.S3.AccessKey = string(secret.Data["accessKey"])
		tenant.Spec.Storage.S3.SecretKey = string(secret.Data["secretKey"])
	}

	// Retrieve meta configurations
	metaConfig := tenant.Spec.Meta
	if metaConfig.MetaAuth.PasswordSecretRef != nil {
		log.V(5).Info("Getting meta password from secret")
		var secret corev1.Secret
		nn := types.NamespacedName{
			Namespace: metaConfig.MetaAuth.PasswordSecretRef.Namespace,
			Name:      metaConfig.MetaAuth.PasswordSecretRef.Name,
		}
		if err := r.Get(ctx, nn, &secret, &client.GetOptions{}); err != nil {
			return nil, fmt.Errorf("failed to get secret %v", nn)
		}
		tenant.Spec.Meta.Password = string(secret.Data["password"])
	}

	// Retrieve user configurations
	for idx, user := range tenant.Spec.Users {
		if user.AuthStringSecretRef != nil {
			var secret corev1.Secret
			nn := types.NamespacedName{
				Namespace: user.AuthStringSecretRef.Namespace,
				Name:      user.AuthStringSecretRef.Name,
			}
			if err := r.Get(ctx, nn, &secret, &client.GetOptions{}); err != nil {
				return nil, fmt.Errorf("failed to get secret %v", nn)
			}
			tenant.Spec.Users[idx].AuthString = string(secret.Data["authString"])
		}
	}

	return &tenant, nil
}

func setCondition(warehouse *v1alpha1.Warehouse, opState opState) {
	var newCond metav1.Condition
	switch opState {
	case createSucceeded:
		newCond = metav1.Condition{
			Type:    v1alpha1.WarehouseCreated,
			Status:  metav1.ConditionTrue,
			Reason:  v1alpha1.WarehouseCreatedReason,
			Message: common.WarehouseCreatedMessage,
		}
	case running:
		newCond = metav1.Condition{
			Type:    v1alpha1.WarehouseRunning,
			Status:  metav1.ConditionTrue,
			Reason:  v1alpha1.WarehouseRunningReason,
			Message: common.WarehouseRunningMessage,
		}
	case buildFailed:
		newCond = metav1.Condition{
			Type:    v1alpha1.WarehouseFailed,
			Status:  metav1.ConditionFalse,
			Reason:  v1alpha1.WarehouseBuildFailedReason,
			Message: common.WarehouseBuildFailedMessage,
		}
	case updateFailed:
		newCond = metav1.Condition{
			Type:    v1alpha1.WarehouseFailed,
			Status:  metav1.ConditionFalse,
			Reason:  v1alpha1.WarehouseBuildFailedReason,
			Message: common.WarehouseUpdateFailedMessage,
		}
	case runFailed:
		newCond = metav1.Condition{
			Type:    v1alpha1.WarehouseFailed,
			Status:  metav1.ConditionFalse,
			Reason:  v1alpha1.WarehouseRunFailedReason,
			Message: common.WarehouseRunFailedMessage,
		}
	}
	meta.SetStatusCondition(&warehouse.Status.Conditions, newCond)
}

// SetupWithManager sets up the controller with the Manager.
func (r *WarehouseReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.Warehouse{}).
		Watches(
			&appsv1.StatefulSet{},
			handler.EnqueueRequestForOwner(
				mgr.GetScheme(), mgr.GetRESTMapper(), &v1alpha1.Warehouse{},
			),
			builder.WithPredicates(predicate.ResourceVersionChangedPredicate{}),
		).
		Named("warehouse").
		Complete(r)
}
