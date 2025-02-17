package query

import (
	v1alpha1 "github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1"
	"github.com/databendcloud/databend-operator/pkg/runtime/resource"
)

type CacheConfig struct {
	DataCacheStorage string           `toml:"data_cache_storage,omitempty"`
	Disk             *DiskCacheConfig `toml:"disk,omitempty"`
}

type DiskCacheConfig struct {
	Path     string `toml:"path"`
	MaxBytes uint64 `toml:"max_bytes"`
}

func NewCacheConfig(tn *v1alpha1.Tenant, wh *v1alpha1.Warehouse) *CacheConfig {
	cfg := &CacheConfig{}
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
