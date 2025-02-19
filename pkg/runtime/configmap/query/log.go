package query

import (
	v1alpha1 "github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1"
)

const (
	OTLPLogEndpoint string = "http://localhost:4318"
)

type QueryLogConfig struct {
	File    QueryLogConfigFile    `toml:"file" json:"file"`
	Stderr  QueryLogConfigStderr  `toml:"stderr" json:"stderr"`
	Query   QueryLogConfigQuery   `toml:"query,omitempty" json:"query,omitempty"`
	Profile QueryLogConfigProfile `toml:"profile,omitempty" json:"profile,omitempty"`
}

type QueryLogConfigFile struct {
	On     bool   `toml:"on" json:"on"`
	Level  string `toml:"level,omitempty" json:"level,omitempty"`
	Dir    string `toml:"dir,omitempty" json:"dir,omitempty"`
	Format string `toml:"format,omitempty" json:"format,omitempty"`
}

type QueryLogConfigStderr struct {
	On     bool   `toml:"on" json:"on"`
	Level  string `toml:"level,omitempty" json:"level,omitempty"`
	Format string `toml:"format,omitempty" json:"format,omitempty"`
}

type QueryLogConfigQuery struct {
	On           bool              `toml:"on" json:"on"`
	Dir          string            `toml:"dir,omitempty" json:"dir,omitempty"`
	OTLPEndpoint string            `toml:"otlp_endpoint,omitempty" json:"otlp_endpoint,omitempty"`
	OTLPProtocol string            `toml:"otlp_protocol,omitempty" json:"otlp_protocol,omitempty"`
	OTLPLabels   map[string]string `toml:"otlp_labels,omitempty" json:"otlp_labels,omitempty"`
}

type QueryLogConfigProfile struct {
	On           bool              `toml:"on" json:"on"`
	Dir          string            `toml:"dir,omitempty" json:"dir,omitempty"`
	OTLPEndpoint string            `toml:"otlp_endpoint,omitempty" json:"otlp_endpoint,omitempty"`
	OTLPProtocol string            `toml:"otlp_protocol,omitempty" json:"otlp_protocol,omitempty"`
	OTLPLabels   map[string]string `toml:"otlp_labels,omitempty" json:"otlp_labels,omitempty"`
}

// NewQueryLogConfig create new instance of Something
func NewQueryLogConfig(wh *v1alpha1.Warehouse) *QueryLogConfig {
	cfg := QueryLogConfig{
		File: QueryLogConfigFile{
			On:     wh.Spec.Log.File.Enabled,
			Level:  "INFO",
			Format: "json",
			Dir:    "/var/log/databend-query",
		},
		Stderr: QueryLogConfigStderr{
			On:     true,
			Level:  "WARN,databend=info,opendal=info,openraft=info",
			Format: "json",
		},
		Query: QueryLogConfigQuery{
			On:           wh.Spec.Log.Query.Enabled,
			Dir:          "",
			OTLPEndpoint: OTLPLogEndpoint,
			OTLPProtocol: "http",
			OTLPLabels: map[string]string{
				"tenant":    wh.Spec.Tenant.Name,
				"warehouse": wh.Name,
			},
		},
		Profile: QueryLogConfigProfile{
			On:           wh.Spec.Log.Profile.Enabled,
			Dir:          "",
			OTLPEndpoint: OTLPLogEndpoint,
			OTLPProtocol: "http",
			OTLPLabels: map[string]string{
				"tenant":    wh.Spec.Tenant.Name,
				"warehouse": wh.Name,
			},
		},
	}

	if wh.Spec.Log.Stderr.Level != "" {
		cfg.Stderr.Level = wh.Spec.Log.Stderr.Level
	}
	if wh.Spec.Log.Query.Endpoint != "" {
		cfg.Query.OTLPEndpoint = wh.Spec.Log.Query.Endpoint
	}
	if wh.Spec.Log.Query.Protocol != "" {
		cfg.Query.OTLPProtocol = wh.Spec.Log.Query.Protocol
	}
	if wh.Spec.Log.Profile.Endpoint != "" {
		cfg.Profile.OTLPEndpoint = wh.Spec.Log.Profile.Endpoint
	}
	if wh.Spec.Log.Profile.Protocol != "" {
		cfg.Profile.OTLPProtocol = wh.Spec.Log.Profile.Protocol
	}

	return &cfg
}
