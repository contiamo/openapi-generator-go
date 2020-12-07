// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// ColumnUpdateRequest is an object. PATCH request body for Column metadata
type ColumnUpdateRequest struct {
	// Description:
	Description *string `json:"description,omitempty"`
	// Documentation:
	Documentation *string `json:"documentation,omitempty"`
	// Entities: A list of detected entities in the column
	Entities []string `json:"entities,omitempty"`
	// Properties:
	Properties map[string]interface{} `json:"properties,omitempty"`
}

// GetDescription returns the Description property
func (m ColumnUpdateRequest) GetDescription() *string {
	return m.Description
}

// SetDescription sets the Description property
func (m ColumnUpdateRequest) SetDescription(val *string) {
	m.Description = val
}

// GetDocumentation returns the Documentation property
func (m ColumnUpdateRequest) GetDocumentation() *string {
	return m.Documentation
}

// SetDocumentation sets the Documentation property
func (m ColumnUpdateRequest) SetDocumentation(val *string) {
	m.Documentation = val
}

// GetEntities returns the Entities property
func (m ColumnUpdateRequest) GetEntities() []string {
	return m.Entities
}

// SetEntities sets the Entities property
func (m ColumnUpdateRequest) SetEntities(val []string) {
	m.Entities = val
}

// GetProperties returns the Properties property
func (m ColumnUpdateRequest) GetProperties() map[string]interface{} {
	return m.Properties
}

// SetProperties sets the Properties property
func (m ColumnUpdateRequest) SetProperties(val map[string]interface{}) {
	m.Properties = val
}
