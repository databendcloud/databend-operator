package runtime

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1alpha1 "github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1"
	"github.com/databendcloud/databend-operator/pkg/common"
	"github.com/databendcloud/databend-operator/pkg/runtime/configmap/query"
	"github.com/databendcloud/databend-operator/pkg/runtime/ingress"
	"github.com/databendcloud/databend-operator/pkg/runtime/objectmeta"
	"github.com/databendcloud/databend-operator/pkg/runtime/service"
	"github.com/databendcloud/databend-operator/pkg/runtime/statefulset"
)

func BuildQueryStatefulSet(tenant *v1alpha1.Tenant, warehouse *v1alpha1.Warehouse) (*appsv1.StatefulSet, error) {
	builder := statefulset.NewStatefulSetBuilder(tenant, warehouse)
	return builder.Build(), nil
}

func BuildQueryConfigMap(tenant *v1alpha1.Tenant, warehouse *v1alpha1.Warehouse) (*corev1.ConfigMap, error) {
	tomlBuilder := query.NewQueryTomlBuilder(tenant, warehouse)
	return tomlBuilder.BuildConfigMap()
}

func BuildQueryService(tenant *v1alpha1.Tenant, warehosue *v1alpha1.Warehouse) (*corev1.Service, error) {
	return service.BuildService(tenant, warehosue), nil
}

func BuildQueryIngress(tenant *v1alpha1.Tenant, warehosue *v1alpha1.Warehouse) (*networkingv1.Ingress, error) {
	return ingress.BuildIngress(tenant, warehosue), nil
}

// BuildServiceAccount Provision iam policy for service account
func BuildTenantServiceAccount(tenant *v1alpha1.Tenant) (*corev1.ServiceAccount, error) {
	sa := &corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      common.GetTenantServiceAccountName(tenant.Name),
			Namespace: tenant.Namespace,
			Labels: map[string]string{
				common.KeyTenant: tenant.Name,
			},
			OwnerReferences: objectmeta.BuildOwnerReferencesByTenant(tenant),
		},
	}
	return sa, nil
}
