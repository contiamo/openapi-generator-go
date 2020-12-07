// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// Permission is an object. A single assignment of a role to a user or group for a specific
// resource.
// Note that this is a sub-resource of a data source so the affected data source is
// implicit and especially its ID is not part of this object.
//
// The permission matrix is as following:
//   | role    | actions                                         |
//   |---------|-------------------------------------------------|
//   | owner   | read, edit, edit-metadata, delete, grant, query |
//   | steward | read, edit-metadata, query                      |
//   | query   | read, query                                     |
type Permission struct {
	// ConnectionId: ID of the connection the beneficiary would use to access the resource if the resource supports connecting to it. This parameter is available only for resources that support connections.
	ConnectionId string `json:"connectionId,omitempty"`
	// Id: ID of the permission
	Id string `json:"id"`
	// PrincipalId: ID of the beneficiary of the permission
	PrincipalId string `json:"principalId"`
	// Role: A valid resource role identifier
	Role Role `json:"role"`
}

// GetConnectionId returns the ConnectionId property
func (m Permission) GetConnectionId() string {
	return m.ConnectionId
}

// SetConnectionId sets the ConnectionId property
func (m Permission) SetConnectionId(val string) {
	m.ConnectionId = val
}

// GetId returns the Id property
func (m Permission) GetId() string {
	return m.Id
}

// SetId sets the Id property
func (m Permission) SetId(val string) {
	m.Id = val
}

// GetPrincipalId returns the PrincipalId property
func (m Permission) GetPrincipalId() string {
	return m.PrincipalId
}

// SetPrincipalId sets the PrincipalId property
func (m Permission) SetPrincipalId(val string) {
	m.PrincipalId = val
}

// GetRole returns the Role property
func (m Permission) GetRole() Role {
	return m.Role
}

// SetRole sets the Role property
func (m Permission) SetRole(val Role) {
	m.Role = val
}
