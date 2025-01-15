package query

import (
	"bytes"

	"github.com/BurntSushi/toml"
	corev1 "k8s.io/api/core/v1"

	databendv1alpha1 "github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1"
	"github.com/databendcloud/databend-operator/pkg/common"
	"github.com/databendcloud/databend-operator/pkg/runtime/config"
	"github.com/databendcloud/databend-operator/pkg/runtime/objectmeta"
)

const (
	// ConfigMapKeyName is the name of the configmap
	ConfigMapKeyName = "config.toml"
)

type QueryTomlBuilder struct {
	warehouse *databendv1alpha1.Warehouse
	tenant    *databendv1alpha1.Tenant
}

type Config struct {
	Log     *QueryLogConfig `toml:"log" json:"log,omitempty"`
	Query   *QueryConfig    `toml:"query" json:"query"`
	Meta    *MetaConfig     `toml:"meta" json:"meta"`
	Storage *StorageConfig  `toml:"storage" json:"storage"`
	Cache   *CacheConfig    `toml:"cache,omitempty" json:"cache,omitempty"`
}

func NewQueryTomlBuilder(tenant *databendv1alpha1.Tenant, warehouse *databendv1alpha1.Warehouse) config.TomlConfig {
	return &QueryTomlBuilder{
		warehouse: warehouse,
		tenant:    tenant,
	}
}

func (b *QueryTomlBuilder) QueryConfig() *Config {
	return &Config{
		Log:     NewQueryLogConfig(b.warehouse),
		Query:   NewQueryConfig(b.tenant, b.warehouse),
		Meta:    NewMetaConfig(b.tenant),
		Storage: NewStorageConfig(b.tenant, b.warehouse),
		Cache:   NewCacheConfig(b.tenant, b.warehouse),
	}
}

func (b *QueryTomlBuilder) BuildConfigMap() (*corev1.ConfigMap, error) {
	// Retrieve toml config from QueryConfig
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(b.QueryConfig()); err != nil {
		return nil, err
	}

	// Build ConfigMap
	configMapName := common.GetQueryConfigMapName(b.warehouse.Name)
	objectMeta := objectmeta.BuildObjectMetaUnderWarehouse(b.warehouse, configMapName)
	configMap := corev1.ConfigMap{
		ObjectMeta: *objectMeta,
	}
	if configMap.Data == nil {
		configMap.Data = make(map[string]string)
	}
	configMap.Data[ConfigMapKeyName] = buf.String()

	return &configMap, nil
}
