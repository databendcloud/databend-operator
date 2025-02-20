package runtime

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"

	v1alpha1 "github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1"
	"github.com/databendcloud/databend-operator/pkg/runtime/configmap/query"
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
	return nil, nil
}

func BuildTenantServiceAccount(tenant *v1alpha1.Tenant) (*corev1.ServiceAccount, error) {
	return nil, nil
}
