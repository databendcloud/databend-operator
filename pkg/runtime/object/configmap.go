package object

import (
	corev1 "k8s.io/api/core/v1"

	databendv1alpha1 "github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1"
	"github.com/databendcloud/databend-operator/pkg/runtime/config/query"
)

func BuildQueryConfigMap(tenant *databendv1alpha1.Tenant, warehouse *databendv1alpha1.Warehouse) (*corev1.ConfigMap, error) {
	tomlBuilder := query.NewQueryTomlBuilder(tenant, warehouse)
	return tomlBuilder.BuildConfigMap()
}
