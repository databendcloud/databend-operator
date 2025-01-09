package query

import (
	"bytes"
	"fmt"

	"github.com/BurntSushi/toml"
	corev1 "k8s.io/api/core/v1"

	databendv1alpha1 "github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1"
	"github.com/databendcloud/databend-operator/pkg/runtime/config"
	"github.com/databendcloud/databend-operator/pkg/runtime/utils"
)

const (
	// ConfigMapKeyName is the name of the configmap
	ConfigMapKeyName = "config.toml"
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
	// Retrieve toml config from QueryConfig
	config, err := b.QueryConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to build query config: %v", err)
	}
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(config); err != nil {
		return nil, err
	}

	// Build ConfigMap
	configMapName := utils.GetQueryConfigMapName(b.warehouse.Name)
	objectMeta := utils.BuildObjectMetaUnderWarehouse(b.warehouse, configMapName)
	configMap := corev1.ConfigMap{
		ObjectMeta: *objectMeta,
	}
	if configMap.Data == nil {
		configMap.Data = make(map[string]string)
	}
	configMap.Data[ConfigMapKeyName] = buf.String()

	return &configMap, nil
}
