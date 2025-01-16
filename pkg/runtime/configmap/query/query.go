package query

import (
	"fmt"

	databendv1alpha1 "github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1"
	"github.com/databendcloud/databend-operator/pkg/common"
)

const (
	ContainerHost   string = "0.0.0.0"
	DefaultUser     string = "admin"
	DefaultPassword string = "admin"
)

type QueryConfig struct {
	TenantId              string `toml:"tenant_id" json:"tenant_id"`
	ClusterId             string `toml:"cluster_id" json:"cluster_id"`
	NumCpus               uint64 `toml:"num_cpus" json:"num_cpus"`
	MaxServerMemoryUsage  uint64 `toml:"max_server_memory_usage" json:"max_server_memory_usage"`
	MaxMemoryLimitEnabled bool   `toml:"max_memory_limit_enabled" json:"max_memory_limit_enabled"`
	MaxRunningQueries     uint64 `toml:"max_running_queries" json:"max_running_queries"`

	FlightApiAddress string `toml:"flight_api_address" json:"flight_api_address"`
	AdminApiAddress  string `toml:"admin_api_address" json:"admin_api_address"`
	MetricApiAddress string `toml:"metric_api_address" json:"metric_api_address"`

	MysqlHandlerHost             string `toml:"mysql_handler_host" json:"mysql_handler_host"`
	MysqlHandlerPort             int    `toml:"mysql_handler_port" json:"mysql_handler_port"`
	ClickhouseHttpHandlerHost    string `toml:"clickhouse_http_handler_host" json:"clickhouse_http_handler_host"`
	ClickhouseHttpHandlerPort    int    `toml:"clickhouse_http_handler_port" json:"clickhouse_http_handler_port"`
	HttpHandlerHost              string `toml:"http_handler_host" json:"http_handler_host"`
	HttpHandlerPort              int    `toml:"http_handler_port" json:"http_handler_port"`
	HttpHandlerResultTimeoutSecs int    `toml:"http_handler_result_timeout_secs" json:"http_handler_result_timeout_secs"`
	FlightSQLHandlerHost         string `toml:"flight_sql_handler_host" json:"flight_sql_handler_host"`
	FlightSQLHandlerPort         int    `toml:"flight_sql_handler_port" json:"flight_sql_handler_port"`
	RpcClientTimeoutSecs         int    `toml:"rpc_client_timeout_secs" json:"rpc_client_timeout_secs"`

	Users []BuiltInUser `toml:"users" json:"users"`

	Settings map[string]string `toml:"settings,omitempty" json:"settings,omitempty"`
}

type BuiltInUser struct {
	Name       string `toml:"name" json:"name"`
	AuthType   string `toml:"auth_type" json:"auth_type"`
	AuthString string `toml:"auth_string" json:"auth_string"`
}

func NewQueryConfig(tn *databendv1alpha1.Tenant, wh *databendv1alpha1.Warehouse) *QueryConfig {
	q := &QueryConfig{
		TenantId:              tn.Name,
		ClusterId:             wh.Name,
		NumCpus:               1,
		MaxServerMemoryUsage:  0,
		MaxMemoryLimitEnabled: false,
		MaxRunningQueries:     8,

		FlightApiAddress: fmt.Sprintf("%s:%d", ContainerHost, common.ServicePortFlight),
		AdminApiAddress:  fmt.Sprintf("%s:%d", ContainerHost, common.ServicePortAdmin),
		MetricApiAddress: fmt.Sprintf("%s:%d", ContainerHost, common.ServicePortMetrics),

		MysqlHandlerHost:             ContainerHost,
		MysqlHandlerPort:             common.ServicePortMySQL,
		ClickhouseHttpHandlerHost:    ContainerHost,
		ClickhouseHttpHandlerPort:    common.ServicePortCKHttp,
		HttpHandlerHost:              ContainerHost,
		HttpHandlerPort:              common.ServicePortQuery,
		HttpHandlerResultTimeoutSecs: 240,
		FlightSQLHandlerHost:         ContainerHost,
		FlightSQLHandlerPort:         common.ServicePortFlightSQL,
		RpcClientTimeoutSecs:         30,

		Settings: wh.Spec.Settings,
	}
	return q.withResourceLimit(wh).withUsers(tn)
}

func (q *QueryConfig) withResourceLimit(wh *databendv1alpha1.Warehouse) *QueryConfig {
	if wh.Spec.PodResource.Limits.Cpu().Value() > 0 {
		q.NumCpus = uint64(wh.Spec.PodResource.Limits.Cpu().Value())
	}
	if wh.Spec.PodResource.Limits.Memory().Value() > 0 {
		q.MaxMemoryLimitEnabled = true
		q.MaxServerMemoryUsage = uint64(wh.Spec.PodResource.Limits.Memory().Value())
	}
	return q
}

func (q *QueryConfig) withUsers(tn *databendv1alpha1.Tenant) *QueryConfig {
	if len(tn.Spec.Users) > 0 {
		for _, u := range tn.Spec.Users {
			q.Users = append(q.Users, BuiltInUser{
				Name:       u.Name,
				AuthType:   string(u.AuthType),
				AuthString: u.AuthString,
			})
		}
	} else {
		// Apend default user
		q.Users = append(q.Users, BuiltInUser{
			Name:       DefaultUser,
			AuthType:   string(databendv1alpha1.SHA256),
			AuthString: common.SHA256String(DefaultPassword),
		})
	}
	return q
}
