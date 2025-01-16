package runtime

import (
	corev1 "k8s.io/api/core/v1"

	v1alpha1 "github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1"
	"github.com/databendcloud/databend-operator/pkg/runtime/configmap/query"
)

func BuildQueryConfigMap(tenant *v1alpha1.Tenant, warehouse *v1alpha1.Warehouse) (*corev1.ConfigMap, error) {
	tomlBuilder := query.NewQueryTomlBuilder(tenant, warehouse)
	return tomlBuilder.BuildConfigMap()
}
