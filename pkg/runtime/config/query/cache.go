package query

import (
	"fmt"

	databendv1alpha1 "github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1"
)

type CacheType string

const (
	DiskCache CacheType = "disk"

	cachePath       = "/var/lib/databend/cache"
	cacheVolumeName = "ephemeral-block-cache"

	Gi = 1024 * 1024 * 1024
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

type CacheSettings struct {
	DataCacheStorage CacheType     `toml:"cache_type"`
	Disk             *DiskSettings `toml:"disk"`
}

type DiskSettings struct {
	Path               string `toml:"path"`
	MaxBytes           uint64 `toml:"max_bytes"`
	K8sResourceRequest string `toml:"k8s_resource_request"`
	K8sResourceLimit   string `toml:"k8s_resource_limit"`
	VolumeName         string `toml:"volume_name"`
}

func NewDiskCacheSetting(size uint64) *CacheSettings {
	if size <= 0 {
		return nil
	}
	return &CacheSettings{
		DataCacheStorage: DiskCache,
		Disk: &DiskSettings{
			Path:               cachePath,
			MaxBytes:           size * Gi,
			K8sResourceRequest: fmt.Sprintf("%dGi", size),
			K8sResourceLimit:   fmt.Sprintf("%dGi", size),
			VolumeName:         cacheVolumeName,
		},
	}
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
	settings := GetCacheSettings(tn, wh)
	if settings == nil {
		return cfg
	}
	cfg.DataCacheStorage = string(settings.DataCacheStorage)
	cfg.Disk = BuildDiskCacheConfig(settings)
	return cfg
}

func BuildDiskCacheConfig(settings *CacheSettings) *DiskCacheConfig {
	return &DiskCacheConfig{
		Path:     settings.Disk.Path,
		MaxBytes: settings.Disk.MaxBytes,
	}
}

func GetCacheSettings(tenant *databendv1alpha1.Tenant, wh *databendv1alpha1.Warehouse) *CacheSettings {
	if wh == nil || !wh.Spec.Cache.Enabled || wh.Spec.Cache.MaxBytes <= 0 {
		return nil
	}

	settings := NewDiskCacheSetting(uint64(wh.Spec.Cache.MaxBytes / Gi))
	if wh.Spec.Cache.MaxBytes < 20*Gi {
		settings = NewDiskCacheSetting(20)
	} else if int64(wh.Spec.Cache.MaxBytes) > wh.Spec.PodResource.Limits.Cpu().Value()*30*Gi {
		settings = NewDiskCacheSetting(20)
	}

	if wh.Spec.Cache.Path != "" {
		settings.Disk.Path = wh.Spec.Cache.Path
	}
	return settings
}
