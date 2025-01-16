package common

const (
	// TenantCreationSucceededMessage is status condition message for the
	// {"type": "Created", "status": "True", "reason": "TenantCreationSucceeded"} condition.
	TenantCreationSucceededMessage = "Succeeded to create Tenant"

	// TenantStorageErrorMessage is status condition message for the
	// {"type": "Error", "status": "False", "reason": "TenantStorageError"} condition.
	TenantStorageErrorMessage = "Invalid storage configuration"

	// TenantStorageErrorMessage is status condition message for the
	// {"type": "Error", "status": "False", "reason": "TenantMetaError"} condition.
	TenantMetaErrorMessage = "Invalid meta configuration"

	// TenantStorageErrorMessage is status condition message for the
	// {"type": "Error", "status": "False", "reason": "TenantUserError"} condition.
	TenantUserErrorMessage = "Invalid built-in user configurations"
)

const (
	// WarehouseRunningMessage is the status condition message for the
	// {"type": "Running", "status": "True", "reason": "WarehouseRunning"} condition.
	WarehouseRunningMessage = "Warehouse is running"

	// WarehouseCreatedMessage is the status condition message for the
	// {"type": "Created", "status": "True", "reason": "WarehouseCreated"} condition.
	WarehouseCreatedMessage = "Succeeded to create Warehouse"

	// WarehouseBuildFailedMessage is the status condition message for the
	// {"type": "Failed", "status": "False", "reason": "WarehouseBuildFailed"} condition.
	WarehouseBuildFailedMessage = "Failed to build Warehouse"

	// WarehouseUpdateFailedMessage is the status condition message for the
	// {"type": "Failed", "status": "False", "reason": "WarehouseBuildFailed"} condition.
	WarehouseUpdateFailedMessage = "Failed to update Warehouse"

	// WarehouseRunFailedMessage is the status condition message for the
	// {"type": "Failed", "status": "False", "reason": "WarehouseRunFailed"} condition.
	WarehouseRunFailedMessage = "Failed to run Warehouse"

	// WarehouseSuspendedMessage is the status condition message for the
	// {"type": "Suspended", "status": "True", "reason": "Suspended"} condition.
	WarehouseSuspendedMessage = "Warehouse is suspended"

	// WarehouseResumedMessage is the status condition message for the
	// {"type": "Suspended", "status": "True", "reason": "Resumed"} condition.
	WarehouseResumedMessage = "Warehouse is resumed"
)

type ServiceProtocol string

const (
	ServiceProtocolFlight    ServiceProtocol = "flight"
	ServiceProtocolAdmin     ServiceProtocol = "admin"
	ServiceProtocolMetrics   ServiceProtocol = "metrics"
	ServiceProtocolMySQL     ServiceProtocol = "mysql"
	ServiceProtocolCKHttp    ServiceProtocol = "ckhttp"
	ServiceProtocolQuery     ServiceProtocol = "query"
	ServiceProtocolFlightSQL ServiceProtocol = "flightsql"
)

const (
	ServicePortFlight    int = 9090
	ServicePortAdmin     int = 8080
	ServicePortMetrics   int = 7070
	ServicePortMySQL     int = 3307
	ServicePortCKHttp    int = 8124
	ServicePortQuery     int = 8000
	ServicePortFlightSQL int = 8900
)

const (
	KeyWarehouse      = "databend.io/warehouse"
	KeyWarehouseSize  = "databend.io/warehouse-size"
	KeyApp            = "databend.io/app"
	ValueAppWarehouse = "warehouse"

	KeyTenant = "databend.io/tenant"
)

const (
	DatabendQueryContainerName string = "databend-query"
	DefaultQueryImage          string = "datafuselabs/databend-query:latest"
)
