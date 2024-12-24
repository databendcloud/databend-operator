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

	apierrors "k8s.io/apimachinery/pkg/api/errors"
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

	_ = ctx
	setCondition(&warehouse, createSucceeded)

	return ctrl.Result{}, nil
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
		Named("warehouse").
		Complete(r)
}
