// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// UpdateViewRequest is an object. Request for updating a View definition in a VirtualDB
type UpdateViewRequest struct {
	// Materialization:
	Materialization MaterializationSpec `json:"materialization"`
	// Sql:
	Sql string `json:"sql"`
}

// GetMaterialization returns the Materialization property
func (m UpdateViewRequest) GetMaterialization() MaterializationSpec {
	return m.Materialization
}

// SetMaterialization sets the Materialization property
func (m UpdateViewRequest) SetMaterialization(val MaterializationSpec) {
	m.Materialization = val
}

// GetSql returns the Sql property
func (m UpdateViewRequest) GetSql() string {
	return m.Sql
}

// SetSql sets the Sql property
func (m UpdateViewRequest) SetSql(val string) {
	m.Sql = val
}
