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

package tenant

import (
	"context"
	"errors"

	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	databendv1alpha1 "github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1"
	"github.com/databendcloud/databend-operator/pkg/common"
)

type opState int

const (
	creationSucceeded opState = iota
	storageError	  opState = iota
	metaError	  	  opState = iota
	builtinUserError  opState = iota
)

// TenantReconciler reconciles a Tenant object
type TenantReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=databendlabs.io,resources=tenants,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=databendlabs.io,resources=tenants/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=databendlabs.io,resources=tenants/finalizers,verbs=update

func (r *TenantReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	var tenant databendv1alpha1.Tenant
	if err := r.Get(ctx, req.NamespacedName, &tenant); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	log := ctrl.LoggerFrom(ctx).WithValues("tenant", klog.KObj(&tenant))
	ctx = ctrl.LoggerInto(ctx, log)
	log.V(2).Info("Reconciling Tenant")

	var err error
	originStatus := tenant.Status.DeepCopy()

	// Reconcile storage
	opState, storageErr := r.ReconcileStorage(ctx, &tenant)
	setCondition(&tenant, opState)
	err = errors.Join(err, storageErr)

	// Reconcile meta
	opState, metaErr := r.ReconcileMeta(ctx, &tenant)
	setCondition(&tenant, opState)
	err = errors.Join(err, metaErr)

	// Reconcile built-in users
	opState, userErr := r.ReconcileBuiltinUsers(ctx, &tenant)
	setCondition(&tenant, opState)
	err = errors.Join(err, userErr)

	if !equality.Semantic.DeepEqual(&tenant.Status, originStatus) {
		return ctrl.Result{}, errors.Join(err, r.Status().Update(ctx, &tenant))
	}
	return ctrl.Result{}, err
}

func (r *TenantReconciler) ReconcileStorage(ctx context.Context, tenant *databendv1alpha1.Tenant) (opState, error) {
	return creationSucceeded, nil
}

func (r *TenantReconciler) ReconcileMeta(ctx context.Context, tenant *databendv1alpha1.Tenant) (opState, error) {
	return creationSucceeded, nil
}

func (r *TenantReconciler) ReconcileBuiltinUsers(ctx context.Context, tenant *databendv1alpha1.Tenant) (opState, error) {
	return creationSucceeded, nil
}

func setCondition(tenant *databendv1alpha1.Tenant, opState opState) {
	var newCond metav1.Condition
	switch opState {
	case creationSucceeded:
		newCond = metav1.Condition{
			Type: databendv1alpha1.TenantCreated,
			Status: metav1.ConditionTrue,
			Message: common.TenantCreationSucceededMessage,
			Reason: databendv1alpha1.TenantCreationSucceededReason,
		}
	case storageError:
		newCond = metav1.Condition{
			Type: databendv1alpha1.TenantError,
			Status: metav1.ConditionFalse,
			Message: common.TenantStorageErrorMessage,
			Reason: databendv1alpha1.TenantStorageErrorReason,
		}
	case metaError:
		newCond = metav1.Condition{
			Type: databendv1alpha1.TenantError,
			Status: metav1.ConditionFalse,
			Message: common.TenantMetaErrorMessage,
			Reason: databendv1alpha1.TenantMetaErrorReason,
		}
	case builtinUserError:
		newCond = metav1.Condition{
			Type: databendv1alpha1.TenantError,
			Status: metav1.ConditionFalse,
			Message: common.TenantUserErrorMessage,
			Reason: databendv1alpha1.TenantUserErrorReason,
		}
	default:
		return
	}
	meta.SetStatusCondition(&tenant.Status.Conditions, newCond)
}

// SetupWithManager sets up the controller with the Manager.
func (r *TenantReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&databendv1alpha1.Tenant{}).
		Named("tenant").
		Complete(r)
}
