package query

import (
	databendv1alpha1 "github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1"
)

type StorageConfig struct {
	StorageType string `toml:"type" json:"type"`
	NumCpus     uint64 `toml:"num_cpus" json:"num_cpus"`

	S3 *S3StorageConfig `toml:"s3" json:"s3"`
}

type S3StorageConfig struct {
	EndpointURL     string `toml:"endpoint_url" json:"endpoint_url"`
	AccessKeyId     string `toml:"access_key_id" json:"access_key_id"`
	SecretAccessKey string `toml:"secret_access_key" json:"secret_access_key"`
	Bucket          string `toml:"bucket" json:"bucket"`
	Root            string `toml:"root" json:"root"`
	Region          string `toml:"region" json:"region"`
}

func NewStorageConfig(tn *databendv1alpha1.Tenant, wh *databendv1alpha1.Warehouse) *StorageConfig {
	return &StorageConfig{
		StorageType: "s3",
		NumCpus:     uint64(wh.Spec.PodResource.Limits.Cpu().Value()),
		S3: &S3StorageConfig{
			AccessKeyId:     tn.Spec.S3.AccessKey,
			SecretAccessKey: tn.Spec.S3.SecretKey,
			Bucket:          tn.Spec.S3.BucketName,
			Root:            tn.Spec.S3.RootPath,
			Region:          tn.Spec.S3.Region,
		},
	}
}
