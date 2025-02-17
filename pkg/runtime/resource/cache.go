package resource

import (
	"fmt"

	v1alpha1 "github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1"
)

type CacheType string

const (
	DiskCache CacheType = "disk"
	NvmeCache CacheType = "nvme"

	cachePath       = "/var/lib/databend/cache"
	cacheVolumeName = "ephemeral-block-cache"

	Gi = 1024 * 1024 * 1024
)

type CacheSettings struct {
	DataCacheStorage   CacheType `toml:"cache_type"`
	Path               string    `toml:"path"`
	MaxBytes           uint64    `toml:"max_bytes"`
	K8sResourceRequest string    `toml:"k8s_resource_request"`
	K8sResourceLimit   string    `toml:"k8s_resource_limit"`
	VolumeName         string    `toml:"volume_name"`
}

func GetCacheSettings(tenant *v1alpha1.Tenant, wh *v1alpha1.Warehouse) *CacheSettings {
	if wh == nil || !wh.Spec.Cache.Enabled || wh.Spec.Cache.MaxSize.Value() <= 0 {
		return nil
	}

	settings := NewDiskCacheSetting(uint64(wh.Spec.Cache.MaxSize.Value() / Gi))
	if wh.Spec.Cache.MaxSize.Value() < 20*Gi {
		settings = NewDiskCacheSetting(20)
	} else if wh.Spec.Cache.MaxSize.Value() > wh.Spec.PodResource.Limits.Cpu().Value()*30*Gi {
		settings = NewDiskCacheSetting(20)
	}

	if wh.Spec.Cache.Path != "" {
		settings.Path = wh.Spec.Cache.Path
	}

	if wh.Spec.Cache.StorageClass != "" {
		settings.DataCacheStorage = DiskCache
	}
	return settings
}

func NewDiskCacheSetting(size uint64) *CacheSettings {
	if size <= 0 {
		return nil
	}
	return &CacheSettings{
		DataCacheStorage:   NvmeCache,
		Path:               cachePath,
		MaxBytes:           size * Gi,
		K8sResourceRequest: fmt.Sprintf("%dGi", size),
		K8sResourceLimit:   fmt.Sprintf("%dGi", size),
		VolumeName:         cacheVolumeName,
	}
}
