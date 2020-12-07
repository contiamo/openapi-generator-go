// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// QueryRequest is an object.
type QueryRequest struct {
	// CustomReference:
	CustomReference string `json:"customReference,omitempty"`
	// QueryId:
	QueryId string `json:"queryId,omitempty"`
	// Sql:
	Sql string `json:"sql"`
}

// GetCustomReference returns the CustomReference property
func (m QueryRequest) GetCustomReference() string {
	return m.CustomReference
}

// SetCustomReference sets the CustomReference property
func (m QueryRequest) SetCustomReference(val string) {
	m.CustomReference = val
}

// GetQueryId returns the QueryId property
func (m QueryRequest) GetQueryId() string {
	return m.QueryId
}

// SetQueryId sets the QueryId property
func (m QueryRequest) SetQueryId(val string) {
	m.QueryId = val
}

// GetSql returns the Sql property
func (m QueryRequest) GetSql() string {
	return m.Sql
}

// SetSql sets the Sql property
func (m QueryRequest) SetSql(val string) {
	m.Sql = val
}
