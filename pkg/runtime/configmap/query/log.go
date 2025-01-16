package query

import (
	"fmt"

	v1alpha1 "github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1"
)

const (
	OTLPTraceEndpoint string = "http://localhost:4317"
	OTLPLogEndpoint   string = "http://localhost:4318"
)

type QueryLogConfig struct {
	File    QueryLogConfigFile    `toml:"file" json:"file"`
	Stderr  QueryLogConfigStderr  `toml:"stderr" json:"stderr"`
	Query   QueryLogConfigQuery   `toml:"query,omitempty" json:"query,omitempty"`
	Profile QueryLogConfigProfile `toml:"profile,omitempty" json:"profile,omitempty"`
	Tracing QueryLogConfigTracing `toml:"tracing,omitempty" json:"tracing,omitempty"`
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

type QueryLogConfigTracing struct {
	On              bool   `toml:"on" json:"on"`
	CaptureLogLevel string `toml:"capture_log_level,omitempty" json:"capture_log_level,omitempty"`
	OTLPEndpoint    string `toml:"otlp_endpoint,omitempty" json:"otlp_endpoint,omitempty"`
}

// NewQueryLogConfig create new instance of Something
func NewQueryLogConfig(wh *v1alpha1.Warehouse) *QueryLogConfig {
	cfg := QueryLogConfig{
		File: QueryLogConfigFile{
			On: false,
		},
		Stderr: QueryLogConfigStderr{
			On:     true,
			Level:  "INFO",
			Format: "json",
		},
		Query: QueryLogConfigQuery{
			On:           true,
			Dir:          "",
			OTLPEndpoint: OTLPLogEndpoint,
			OTLPProtocol: "http",
		},
		Profile: QueryLogConfigProfile{
			On:           true,
			Dir:          "",
			OTLPEndpoint: OTLPLogEndpoint,
			OTLPProtocol: "http",
		},
		Tracing: QueryLogConfigTracing{
			On:              false,
			CaptureLogLevel: "DEBUG",
			OTLPEndpoint:    OTLPTraceEndpoint,
		},
	}
	if wh == nil {
		return &cfg
	}
	cfg.Query.OTLPLabels = map[string]string{
		"tenant":         wh.Spec.Tenant.Name,
		"warehouse":      wh.Name,
		"warehouse_size": fmt.Sprint(wh.Spec.Replicas),
	}
	cfg.Profile.OTLPLabels = map[string]string{
		"tenant":         wh.Spec.Tenant.Name,
		"warehouse":      wh.Name,
		"warehouse_size": fmt.Sprint(wh.Spec.Replicas),
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
	if !wh.Spec.Log.Query.Enabled {
		cfg.Query.On = false
		cfg.Profile.On = false
	}
	return &cfg
}
