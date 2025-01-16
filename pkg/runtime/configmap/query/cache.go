package query

import (
	databendv1alpha1 "github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1"
	"github.com/databendcloud/databend-operator/pkg/runtime/resource"
)

type CacheConfig struct {
	DataCacheStorage               string           `toml:"data_cache_storage,omitempty"`
	Disk                           *DiskCacheConfig `toml:"disk,omitempty"`
	TableDataDeserializedDataBytes uint64           `toml:"table_data_deserialized_data_bytes,omitempty"`
	TableBloomIndexFilterSize      uint64           `toml:"table_bloom_index_filter_size,omitempty"`
	TableBloomIndexMetaCount       uint64           `toml:"table_bloom_index_meta_count,omitempty"`
}

type DiskCacheConfig struct {
	Path     string `toml:"path"`
	MaxBytes uint64 `toml:"max_bytes"`
}

func NewCacheConfig(tn *databendv1alpha1.Tenant, wh *databendv1alpha1.Warehouse) *CacheConfig {
	cfg := &CacheConfig{}
	cfg.TableBloomIndexFilterSize = 5368709120
	cfg.TableBloomIndexMetaCount = 200000

	// TODO(ws): ask about this configuration
	// cfg.TableDataDeserializedDataBytes = resource.MaxTableMemoryCache

	if wh == nil || tn == nil {
		return cfg
	}
	settings := resource.GetCacheSettings(tn, wh)
	if settings == nil {
		return cfg
	}
	cfg.DataCacheStorage = string(settings.DataCacheStorage)
	cfg.Disk = &DiskCacheConfig{
		Path:     settings.Path,
		MaxBytes: settings.MaxBytes,
	}
	return cfg
}
