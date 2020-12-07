// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// UpdatePermissionRequest is an object.
type UpdatePermissionRequest struct {
	// ConnectionId: ID of the connection the beneficiary would use to access the resource if the resource supports connecting to it. This parameter is optional and if not speicified the primary connection will be assigned to the beneficiary.
	ConnectionId string `json:"connectionId,omitempty"`
	// Role: A valid resource role identifier
	Role Role `json:"role"`
}

// GetConnectionId returns the ConnectionId property
func (m UpdatePermissionRequest) GetConnectionId() string {
	return m.ConnectionId
}

// SetConnectionId sets the ConnectionId property
func (m UpdatePermissionRequest) SetConnectionId(val string) {
	m.ConnectionId = val
}

// GetRole returns the Role property
func (m UpdatePermissionRequest) GetRole() Role {
	return m.Role
}

// SetRole sets the Role property
func (m UpdatePermissionRequest) SetRole(val Role) {
	m.Role = val
}
