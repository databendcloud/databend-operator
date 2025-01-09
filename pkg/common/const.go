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

const (
	KeyWarehouse      = "databend.io/warehouse"
	KeyWarehouseSize  = "databend.io/warehouse-size"
	KeyApp            = "databend.io/app"
	ValueAppWarehouse = "warehouse"

	KeyTenant = "databend.io/tenant"
)
