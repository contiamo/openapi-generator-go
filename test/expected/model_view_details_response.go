// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// ViewDetailsResponse is an object.
type ViewDetailsResponse struct {
	// LastUpdateErrorMessage: contains an error message indicating why this table schema could not be updated
	LastUpdateErrorMessage string `json:"lastUpdateErrorMessage,omitempty"`
	// Materialization:
	Materialization MaterializationSpec `json:"materialization"`
	// Name:
	Name string `json:"name"`
	// Sql:
	Sql string `json:"sql"`
}

// GetLastUpdateErrorMessage returns the LastUpdateErrorMessage property
func (m ViewDetailsResponse) GetLastUpdateErrorMessage() string {
	return m.LastUpdateErrorMessage
}

// SetLastUpdateErrorMessage sets the LastUpdateErrorMessage property
func (m ViewDetailsResponse) SetLastUpdateErrorMessage(val string) {
	m.LastUpdateErrorMessage = val
}

// GetMaterialization returns the Materialization property
func (m ViewDetailsResponse) GetMaterialization() MaterializationSpec {
	return m.Materialization
}

// SetMaterialization sets the Materialization property
func (m ViewDetailsResponse) SetMaterialization(val MaterializationSpec) {
	m.Materialization = val
}

// GetName returns the Name property
func (m ViewDetailsResponse) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m ViewDetailsResponse) SetName(val string) {
	m.Name = val
}

// GetSql returns the Sql property
func (m ViewDetailsResponse) GetSql() string {
	return m.Sql
}

// SetSql sets the Sql property
func (m ViewDetailsResponse) SetSql(val string) {
	m.Sql = val
}
