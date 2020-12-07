// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// InternalAuthenticationResponse is an object.
type InternalAuthenticationResponse struct {
	// Connection:
	Connection Connection `json:"connection"`
	// Identity: Short description of the identity accessing the system.
	Identity Identity `json:"identity"`
}

// GetConnection returns the Connection property
func (m InternalAuthenticationResponse) GetConnection() Connection {
	return m.Connection
}

// SetConnection sets the Connection property
func (m InternalAuthenticationResponse) SetConnection(val Connection) {
	m.Connection = val
}

// GetIdentity returns the Identity property
func (m InternalAuthenticationResponse) GetIdentity() Identity {
	return m.Identity
}

// SetIdentity sets the Identity property
func (m InternalAuthenticationResponse) SetIdentity(val Identity) {
	m.Identity = val
}
