// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// TableProperties is an object.
type TableProperties struct {
	// CurrentColumnSetId:
	CurrentColumnSetId string `json:"currentColumnSetId"`
	// LastUpdateErrorMessage: contains an error message indicating why this table schema could not be updated
	LastUpdateErrorMessage string `json:"lastUpdateErrorMessage,omitempty"`
	// Sql:
	Sql string `json:"sql"`
	// Type:
	Type TableType `json:"type"`
}

// GetCurrentColumnSetId returns the CurrentColumnSetId property
func (m TableProperties) GetCurrentColumnSetId() string {
	return m.CurrentColumnSetId
}

// SetCurrentColumnSetId sets the CurrentColumnSetId property
func (m TableProperties) SetCurrentColumnSetId(val string) {
	m.CurrentColumnSetId = val
}

// GetLastUpdateErrorMessage returns the LastUpdateErrorMessage property
func (m TableProperties) GetLastUpdateErrorMessage() string {
	return m.LastUpdateErrorMessage
}

// SetLastUpdateErrorMessage sets the LastUpdateErrorMessage property
func (m TableProperties) SetLastUpdateErrorMessage(val string) {
	m.LastUpdateErrorMessage = val
}

// GetSql returns the Sql property
func (m TableProperties) GetSql() string {
	return m.Sql
}

// SetSql sets the Sql property
func (m TableProperties) SetSql(val string) {
	m.Sql = val
}

// GetType returns the Type property
func (m TableProperties) GetType() TableType {
	return m.Type
}

// SetType sets the Type property
func (m TableProperties) SetType(val TableType) {
	m.Type = val
}
