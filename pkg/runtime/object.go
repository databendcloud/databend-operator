package runtime

import (
	appsv1 "k8s.io/api/apps/v1"

	v1alpha1 "github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1"
	"github.com/databendcloud/databend-operator/pkg/runtime/statefulset"
)

func BuildStatefulSet(tenant *v1alpha1.Tenant, warehouse *v1alpha1.Warehouse) (*appsv1.StatefulSet, error) {
	builder := statefulset.NewStatefulSetBuilder(tenant, warehouse)
	return builder.Build(), nil
}
