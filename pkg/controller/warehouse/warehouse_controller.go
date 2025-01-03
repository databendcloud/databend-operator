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
	"sigs.k8s.io/controller-runtime/pkg/client"

	databendv1alpha1 "github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1"
	"github.com/databendcloud/databend-operator/pkg/common"
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

func (r *WarehouseReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := ctrl.LoggerFrom(ctx)

	var warehouse databendv1alpha1.Warehouse
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

	opState, err := r.reconcileStatefulSet(ctx, tenant, &warehouse)
	if opState == createSucceeded {
		opState, err = r.updateReplicas(ctx, &warehouse)
	}
	setCondition(&warehouse, opState)
	if !equality.Semantic.DeepEqual(warehouse.Status, originStatus) {
		return ctrl.Result{}, errors.Join(err, r.Status().Update(ctx, &warehouse))
	}

	return ctrl.Result{}, err
}

func (r *WarehouseReconciler) reconcileStatefulSet(ctx context.Context, tenant *databendv1alpha1.Tenant, warehouse *databendv1alpha1.Warehouse) (opState, error) {
	log := ctrl.LoggerFrom(ctx)

	// Build and reconcile StatefulSet
	ss, err := buildStatefulSet(ctx, tenant, warehouse)
	if err != nil {
		log.V(5).Error(err, "Failed to build StatefulSet", "namespace", ss.Namespace, "name", ss.Name)
		return buildFailed, err
	}

	// Non-empty resourceVersion indicates UPDATE operation.
	var creationErr error
	var created bool
	if warehouse.GetResourceVersion() == "" {
		creationErr = r.Create(ctx, ss)
		created = creationErr == nil
	}
	switch {
	case created:
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

	return createSucceeded, nil
}

func (r *WarehouseReconciler) updateReplicas(ctx context.Context, warehouse *databendv1alpha1.Warehouse) (opState, error) {
	log := ctrl.LoggerFrom(ctx)

	var ss appsv1.StatefulSet
	ssNN := types.NamespacedName{
		Namespace: warehouse.Namespace,
		Name:      warehouse.Name,
	}
	if err := r.Get(ctx, ssNN, &ss); err != nil {
		log.V(5).Error(err, "Failed to get StatefulSet", "namespacedName", ssNN)
		return buildFailed, err
	}

	warehouse.Status.ReadyReplicas = int(ss.Status.ReadyReplicas)
	if ss.Status.ReadyReplicas == *ss.Spec.Replicas {
		return running, nil
	}

	return createSucceeded, nil
}

func (r *WarehouseReconciler) getTenant(ctx context.Context, nn types.NamespacedName) (*databendv1alpha1.Tenant, error) {
	log := ctrl.LoggerFrom(ctx)

	var tenant databendv1alpha1.Tenant
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

func buildStatefulSet(ctx context.Context, tenant *databendv1alpha1.Tenant, warehouse *databendv1alpha1.Warehouse) (*appsv1.StatefulSet, error) {
	_, _, _ = ctx, tenant, warehouse
	return nil, nil
}

func setCondition(warehouse *databendv1alpha1.Warehouse, opState opState) {
	var newCond metav1.Condition
	switch opState {
	case createSucceeded:
		newCond = metav1.Condition{
			Type:    databendv1alpha1.WarehouseCreated,
			Status:  metav1.ConditionTrue,
			Reason:  databendv1alpha1.WarehouseCreatedReason,
			Message: common.WarehouseCreatedMessage,
		}
	case running:
		newCond = metav1.Condition{
			Type:    databendv1alpha1.WarehouseRunning,
			Status:  metav1.ConditionTrue,
			Reason:  databendv1alpha1.WarehouseRunningReason,
			Message: common.WarehouseRunningMessage,
		}
	case buildFailed:
		newCond = metav1.Condition{
			Type:    databendv1alpha1.WarehouseFailed,
			Status:  metav1.ConditionFalse,
			Reason:  databendv1alpha1.WarehouseBuildFailedReason,
			Message: common.WarehouseBuildFailedMessage,
		}
	case updateFailed:
		newCond = metav1.Condition{
			Type:    databendv1alpha1.WarehouseFailed,
			Status:  metav1.ConditionFalse,
			Reason:  databendv1alpha1.WarehouseBuildFailedReason,
			Message: common.WarehouseUpdateFailedMessage,
		}
	case runFailed:
		newCond = metav1.Condition{
			Type:    databendv1alpha1.WarehouseFailed,
			Status:  metav1.ConditionFalse,
			Reason:  databendv1alpha1.WarehouseRunFailedReason,
			Message: common.WarehouseRunFailedMessage,
		}
	}
	meta.SetStatusCondition(&warehouse.Status.Conditions, newCond)
}

// SetupWithManager sets up the controller with the Manager.
func (r *WarehouseReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&databendv1alpha1.Warehouse{}).
		Owns(&appsv1.StatefulSet{}).
		Named("warehouse").
		Complete(r)
}
