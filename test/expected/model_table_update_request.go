// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// TableUpdateRequest is an object. PATCH request body for Table metadata
type TableUpdateRequest struct {
	// Description:
	Description *string `json:"description,omitempty"`
	// Documentation:
	Documentation *string `json:"documentation,omitempty"`
	// Properties:
	Properties map[string]interface{} `json:"properties,omitempty"`
}

// GetDescription returns the Description property
func (m TableUpdateRequest) GetDescription() *string {
	return m.Description
}

// SetDescription sets the Description property
func (m TableUpdateRequest) SetDescription(val *string) {
	m.Description = val
}

// GetDocumentation returns the Documentation property
func (m TableUpdateRequest) GetDocumentation() *string {
	return m.Documentation
}

// SetDocumentation sets the Documentation property
func (m TableUpdateRequest) SetDocumentation(val *string) {
	m.Documentation = val
}

// GetProperties returns the Properties property
func (m TableUpdateRequest) GetProperties() map[string]interface{} {
	return m.Properties
}

// SetProperties sets the Properties property
func (m TableUpdateRequest) SetProperties(val map[string]interface{}) {
	m.Properties = val
}
