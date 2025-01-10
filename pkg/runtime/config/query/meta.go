package query

import (
	databendv1alpha1 "github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1"
)

type MetaConfig struct {
	Endpoints []string `toml:"endpoints" json:"endpoints"`
	Username  string   `toml:"username" json:"username"`
	Password  string   `toml:"password" json:"password"`

	ClientTimeoutInSecond int `toml:"client_timeout_in_second" json:"client_timeout_in_second"`
	AutoSyncInterval      int `toml:"auto_sync_interval" json:"auto_sync_interval"`
}

func NewMetaConfig(tn *databendv1alpha1.Tenant) *MetaConfig {
	meta := &MetaConfig{
		Endpoints:             tn.Spec.Meta.Endpoints,
		Username:              "root",
		Password:              "root",
		ClientTimeoutInSecond: tn.Spec.Meta.TimeoutInSeconds,
		AutoSyncInterval:      tn.Spec.Meta.AutoSyncInterval,
	}
	if tn.Spec.Meta.User != "" && tn.Spec.Meta.Password != "" {
		meta.Username = tn.Spec.Meta.User
		meta.Password = tn.Spec.Meta.Password
	}
	return meta
}
