// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// CreateConnectionRequest is an object.
type CreateConnectionRequest struct {
	// Name:
	Name string `json:"name"`
	// Properties:
	Properties ConnectionProperties `json:"properties"`
	// Type:
	Type ConnectionType `json:"type"`
}

// GetName returns the Name property
func (m CreateConnectionRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m CreateConnectionRequest) SetName(val string) {
	m.Name = val
}

// GetProperties returns the Properties property
func (m CreateConnectionRequest) GetProperties() ConnectionProperties {
	return m.Properties
}

// SetProperties sets the Properties property
func (m CreateConnectionRequest) SetProperties(val ConnectionProperties) {
	m.Properties = val
}

// GetType returns the Type property
func (m CreateConnectionRequest) GetType() ConnectionType {
	return m.Type
}

// SetType sets the Type property
func (m CreateConnectionRequest) SetType(val ConnectionType) {
	m.Type = val
}
