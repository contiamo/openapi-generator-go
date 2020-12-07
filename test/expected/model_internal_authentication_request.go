// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// InternalAuthenticationRequest is an object.
type InternalAuthenticationRequest struct {
	// ConnectionName: Name of the data source connection to auth against. Optional, authenticated against user's defualt if not set.
	ConnectionName string `json:"connectionName,omitempty"`
	// Identifier: User email if a personal access token is provided
	Identifier string `json:"identifier,omitempty"`
	// Token: Personal access token or service account token
	Token string `json:"token"`
}

// GetConnectionName returns the ConnectionName property
func (m InternalAuthenticationRequest) GetConnectionName() string {
	return m.ConnectionName
}

// SetConnectionName sets the ConnectionName property
func (m InternalAuthenticationRequest) SetConnectionName(val string) {
	m.ConnectionName = val
}

// GetIdentifier returns the Identifier property
func (m InternalAuthenticationRequest) GetIdentifier() string {
	return m.Identifier
}

// SetIdentifier sets the Identifier property
func (m InternalAuthenticationRequest) SetIdentifier(val string) {
	m.Identifier = val
}

// GetToken returns the Token property
func (m InternalAuthenticationRequest) GetToken() string {
	return m.Token
}

// SetToken sets the Token property
func (m InternalAuthenticationRequest) SetToken(val string) {
	m.Token = val
}
