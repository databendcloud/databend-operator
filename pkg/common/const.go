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
