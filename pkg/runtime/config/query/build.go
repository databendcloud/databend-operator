package query

import (
	corev1 "k8s.io/api/core/v1"

	databendv1alpha1 "github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1"
	"github.com/databendcloud/databend-operator/pkg/runtime/config"
)

type QueryTomlBuilder struct {
	warehouse *databendv1alpha1.Warehouse
	tenant    *databendv1alpha1.Tenant
}

func NewQueryTomlBuilder(tenant *databendv1alpha1.Tenant, warehouse *databendv1alpha1.Warehouse) config.TomlConfig {
	return &QueryTomlBuilder{
		warehouse: warehouse,
		tenant:    tenant,
	}
}

func (b *QueryTomlBuilder) BuildConfigMap() (*corev1.ConfigMap, error) {
	return nil, nil
}
