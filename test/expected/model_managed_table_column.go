// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// ManagedTableColumn is an object.
type ManagedTableColumn struct {
	// Name: The name of the column, length restrictions are inherited from the underlying database restriction on identifiers
	Name string `json:"name"`
	// Required:
	Required bool `json:"required"`
	// Type: The data type of the column in a managed table
	Type ManagedTableColumnType `json:"type"`
}

// GetName returns the Name property
func (m ManagedTableColumn) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m ManagedTableColumn) SetName(val string) {
	m.Name = val
}

// GetRequired returns the Required property
func (m ManagedTableColumn) GetRequired() bool {
	return m.Required
}

// SetRequired sets the Required property
func (m ManagedTableColumn) SetRequired(val bool) {
	m.Required = val
}

// GetType returns the Type property
func (m ManagedTableColumn) GetType() ManagedTableColumnType {
	return m.Type
}

// SetType sets the Type property
func (m ManagedTableColumn) SetType(val ManagedTableColumnType) {
	m.Type = val
}
