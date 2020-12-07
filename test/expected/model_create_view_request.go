// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// CreateViewRequest is an object. Request for defining a View in a VirtualDB
type CreateViewRequest struct {
	// Description:
	Description string `json:"description,omitempty"`
	// Materialization:
	Materialization MaterializationSpec `json:"materialization"`
	// Name:
	Name string `json:"name"`
	// Sql:
	Sql string `json:"sql"`
}

// GetDescription returns the Description property
func (m CreateViewRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m CreateViewRequest) SetDescription(val string) {
	m.Description = val
}

// GetMaterialization returns the Materialization property
func (m CreateViewRequest) GetMaterialization() MaterializationSpec {
	return m.Materialization
}

// SetMaterialization sets the Materialization property
func (m CreateViewRequest) SetMaterialization(val MaterializationSpec) {
	m.Materialization = val
}

// GetName returns the Name property
func (m CreateViewRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m CreateViewRequest) SetName(val string) {
	m.Name = val
}

// GetSql returns the Sql property
func (m CreateViewRequest) GetSql() string {
	return m.Sql
}

// SetSql sets the Sql property
func (m CreateViewRequest) SetSql(val string) {
	m.Sql = val
}
