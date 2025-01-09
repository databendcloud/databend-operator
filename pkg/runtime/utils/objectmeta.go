package utils

import (
	"fmt"

	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	databendv1alpha1 "github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1"
	"github.com/databendcloud/databend-operator/pkg/common"
)

var (
	WarehouseGVK = schema.GroupVersionKind{
		Group:   databendv1alpha1.SchemeGroupVersion.Group,
		Version: databendv1alpha1.SchemeGroupVersion.Version,
		Kind:    databendv1alpha1.WarehouseKind,
	}
	TenantGVK = schema.GroupVersionKind{
		Group:   databendv1alpha1.SchemeGroupVersion.Group,
		Version: databendv1alpha1.SchemeGroupVersion.Version,
		Kind:    databendv1alpha1.TenantKind,
	}
)

func IsWarehouseRef(ownerRef metav1.OwnerReference) bool {
	if ownerRef.Kind == WarehouseGVK.Kind && ownerRef.APIVersion == WarehouseGVK.GroupVersion().String() {
		return true
	}
	return false
}

func IsTenantRef(ownerRef metav1.OwnerReference) bool {
	if ownerRef.Kind == TenantGVK.Kind && ownerRef.APIVersion == TenantGVK.GroupVersion().String() {
		return true
	}
	return false
}

func CheckOwnerRef(ref []metav1.OwnerReference) (bool, *metav1.OwnerReference) {
	if len(ref) == 0 {
		return false, nil
	}
	for _, ownerRef := range ref {
		if IsWarehouseRef(ownerRef) {
			return true, &ownerRef
		} else if IsTenantRef(ownerRef) {
			return true, &ownerRef
		}
	}
	return false, nil
}

func OwnedByTenant(r []metav1.OwnerReference, tenant *databendv1alpha1.Tenant) error {
	has, ref := CheckOwnerRef(r)
	if !has {
		return common.OwnerNotFound
	}
	if IsTenantRef(*ref) && ref.Name == tenant.Name {
		return nil
	}

	return errors.Wrap(common.OwnedByOtherIdentity, ref.String())
}

func OwnedByWarehouse(r []metav1.OwnerReference, wh *databendv1alpha1.Warehouse) error {
	has, ref := CheckOwnerRef(r)
	if !has {
		return common.OwnerNotFound
	}
	if IsWarehouseRef(*ref) && ref.Name == wh.Name {
		return nil
	}

	return errors.Wrap(common.OwnedByOtherIdentity, ref.String())
}

func LabelsFromTenant(tenant *databendv1alpha1.Tenant) map[string]string {
	ll := make(map[string]string)

	ll[common.KeyTenant] = tenant.Name

	return ll
}

func LabelsFromWarehouse(wh *databendv1alpha1.Warehouse) map[string]string {
	ll := make(map[string]string)

	ll[common.KeyTenant] = wh.Spec.Tenant.Name
	ll[common.KeyWarehouse] = wh.Name
	ll[common.KeyWarehouseSize] = fmt.Sprint(wh.Spec.Replicas)
	ll[common.KeyApp] = common.ValueAppWarehouse

	return ll
}

func BuildAnnotationsFromTenant(tenant *databendv1alpha1.Tenant) map[string]string {
	annotations := make(map[string]string)

	annotations[common.KeyTenant] = tenant.Name

	return annotations
}

func BuildAnnotationsFromWarehouse(wh *databendv1alpha1.Warehouse) map[string]string {
	annotations := make(map[string]string)

	annotations[common.KeyTenant] = wh.Spec.Tenant.Name
	annotations[common.KeyWarehouse] = wh.Name
	annotations[common.KeyWarehouseSize] = fmt.Sprint(wh.Spec.Replicas)

	return annotations
}

func BuildOwnerReferencesByTenant(tenant *databendv1alpha1.Tenant) []metav1.OwnerReference {
	var apiVersion, kind string
	if len(tenant.APIVersion) == 0 || len(tenant.Kind) == 0 {
		apiVersion = TenantGVK.GroupVersion().String()
		kind = TenantGVK.Kind
	} else {
		apiVersion = tenant.APIVersion
		kind = tenant.Kind
	}
	return []metav1.OwnerReference{
		{
			APIVersion: apiVersion,
			Kind:       kind,
			Name:       tenant.Name,
			UID:        tenant.UID,
			//// tenant CRD cannot be deleted unless all reference got deleted
			//BlockOwnerDeletion: ptr.Bool(true),
		},
	}
}

func BuildOwnerReferencesByWarehouse(wh *databendv1alpha1.Warehouse) []metav1.OwnerReference {
	var apiVersion, kind string
	if len(wh.APIVersion) == 0 || len(wh.Kind) == 0 {
		apiVersion = WarehouseGVK.GroupVersion().String()
		kind = WarehouseGVK.Kind
	} else {
		apiVersion = wh.APIVersion
		kind = wh.Kind
	}
	return []metav1.OwnerReference{
		{
			APIVersion: apiVersion,
			Kind:       kind,
			Name:       wh.Name,
			UID:        wh.UID,
		},
	}
}

func BuildObjectMetaUnderTenant(tenant *databendv1alpha1.Tenant, name string) *metav1.ObjectMeta {
	ll := LabelsFromTenant(tenant)
	annotations := BuildAnnotationsFromTenant(tenant)
	meta := &metav1.ObjectMeta{
		Name:            name,
		Namespace:       tenant.Namespace,
		Labels:          ll,
		Annotations:     annotations,
		OwnerReferences: BuildOwnerReferencesByTenant(tenant),
	}
	return meta
}

// initialize object meta for the workloads like configMap, statefulset, service, etc to be created under the warehouse.
func BuildObjectMetaUnderWarehouse(wh *databendv1alpha1.Warehouse, name string) *metav1.ObjectMeta {
	ll := LabelsFromWarehouse(wh)
	annotations := BuildAnnotationsFromWarehouse(wh)
	meta := &metav1.ObjectMeta{
		Name:            name,
		Namespace:       wh.Namespace,
		Labels:          ll,
		Annotations:     annotations,
		OwnerReferences: BuildOwnerReferencesByWarehouse(wh),
	}
	return meta
}
