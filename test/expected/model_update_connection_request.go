// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// UpdateConnectionRequest is an object.
type UpdateConnectionRequest struct {
	// Description:
	Description *string `json:"description,omitempty"`
	// Name:
	Name *string `json:"name,omitempty"`
	// Properties:
	Properties *interface{} `json:"properties,omitempty"`
	// Type:
	Type *interface{} `json:"type,omitempty"`
}

// GetDescription returns the Description property
func (m UpdateConnectionRequest) GetDescription() *string {
	return m.Description
}

// SetDescription sets the Description property
func (m UpdateConnectionRequest) SetDescription(val *string) {
	m.Description = val
}

// GetName returns the Name property
func (m UpdateConnectionRequest) GetName() *string {
	return m.Name
}

// SetName sets the Name property
func (m UpdateConnectionRequest) SetName(val *string) {
	m.Name = val
}

// GetProperties returns the Properties property
func (m UpdateConnectionRequest) GetProperties() *interface{} {
	return m.Properties
}

// SetProperties sets the Properties property
func (m UpdateConnectionRequest) SetProperties(val *interface{}) {
	m.Properties = val
}

// GetType returns the Type property
func (m UpdateConnectionRequest) GetType() *interface{} {
	return m.Type
}

// SetType sets the Type property
func (m UpdateConnectionRequest) SetType(val *interface{}) {
	m.Type = val
}
