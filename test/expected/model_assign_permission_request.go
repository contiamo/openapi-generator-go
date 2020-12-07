// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// AssignPermissionRequest is an object.
type AssignPermissionRequest struct {
	// ConnectionId: ID of the connection the beneficiary would use to access the resource if the resource supports connecting to it. This parameter is optional and if not speicified the primary connection will be assigned to the beneficiary.
	ConnectionId string `json:"connectionId,omitempty"`
	// PrincipalId: ID of the beneficiary of the permission
	PrincipalId string `json:"principalId"`
	// Role: A valid resource role identifier
	Role Role `json:"role"`
}

// GetConnectionId returns the ConnectionId property
func (m AssignPermissionRequest) GetConnectionId() string {
	return m.ConnectionId
}

// SetConnectionId sets the ConnectionId property
func (m AssignPermissionRequest) SetConnectionId(val string) {
	m.ConnectionId = val
}

// GetPrincipalId returns the PrincipalId property
func (m AssignPermissionRequest) GetPrincipalId() string {
	return m.PrincipalId
}

// SetPrincipalId sets the PrincipalId property
func (m AssignPermissionRequest) SetPrincipalId(val string) {
	m.PrincipalId = val
}

// GetRole returns the Role property
func (m AssignPermissionRequest) GetRole() Role {
	return m.Role
}

// SetRole sets the Role property
func (m AssignPermissionRequest) SetRole(val Role) {
	m.Role = val
}
