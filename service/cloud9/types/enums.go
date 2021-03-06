// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ConnectionType string

// Enum values for ConnectionType
const (
	ConnectionTypeConnectSsh ConnectionType = "CONNECT_SSH"
	ConnectionTypeConnectSsm ConnectionType = "CONNECT_SSM"
)

// Values returns all known values for ConnectionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ConnectionType) Values() []ConnectionType {
	return []ConnectionType{
		"CONNECT_SSH",
		"CONNECT_SSM",
	}
}

type EnvironmentLifecycleStatus string

// Enum values for EnvironmentLifecycleStatus
const (
	EnvironmentLifecycleStatusCreating     EnvironmentLifecycleStatus = "CREATING"
	EnvironmentLifecycleStatusCreated      EnvironmentLifecycleStatus = "CREATED"
	EnvironmentLifecycleStatusCreateFailed EnvironmentLifecycleStatus = "CREATE_FAILED"
	EnvironmentLifecycleStatusDeleting     EnvironmentLifecycleStatus = "DELETING"
	EnvironmentLifecycleStatusDeleteFailed EnvironmentLifecycleStatus = "DELETE_FAILED"
)

// Values returns all known values for EnvironmentLifecycleStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (EnvironmentLifecycleStatus) Values() []EnvironmentLifecycleStatus {
	return []EnvironmentLifecycleStatus{
		"CREATING",
		"CREATED",
		"CREATE_FAILED",
		"DELETING",
		"DELETE_FAILED",
	}
}

type EnvironmentStatus string

// Enum values for EnvironmentStatus
const (
	EnvironmentStatusError      EnvironmentStatus = "error"
	EnvironmentStatusCreating   EnvironmentStatus = "creating"
	EnvironmentStatusConnecting EnvironmentStatus = "connecting"
	EnvironmentStatusReady      EnvironmentStatus = "ready"
	EnvironmentStatusStopping   EnvironmentStatus = "stopping"
	EnvironmentStatusStopped    EnvironmentStatus = "stopped"
	EnvironmentStatusDeleting   EnvironmentStatus = "deleting"
)

// Values returns all known values for EnvironmentStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EnvironmentStatus) Values() []EnvironmentStatus {
	return []EnvironmentStatus{
		"error",
		"creating",
		"connecting",
		"ready",
		"stopping",
		"stopped",
		"deleting",
	}
}

type EnvironmentType string

// Enum values for EnvironmentType
const (
	EnvironmentTypeSsh EnvironmentType = "ssh"
	EnvironmentTypeEc2 EnvironmentType = "ec2"
)

// Values returns all known values for EnvironmentType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EnvironmentType) Values() []EnvironmentType {
	return []EnvironmentType{
		"ssh",
		"ec2",
	}
}

type MemberPermissions string

// Enum values for MemberPermissions
const (
	MemberPermissionsReadWrite MemberPermissions = "read-write"
	MemberPermissionsReadOnly  MemberPermissions = "read-only"
)

// Values returns all known values for MemberPermissions. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (MemberPermissions) Values() []MemberPermissions {
	return []MemberPermissions{
		"read-write",
		"read-only",
	}
}

type Permissions string

// Enum values for Permissions
const (
	PermissionsOwner     Permissions = "owner"
	PermissionsReadWrite Permissions = "read-write"
	PermissionsReadOnly  Permissions = "read-only"
)

// Values returns all known values for Permissions. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (Permissions) Values() []Permissions {
	return []Permissions{
		"owner",
		"read-write",
		"read-only",
	}
}
