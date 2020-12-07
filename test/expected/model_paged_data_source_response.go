// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// PagedDataSourceResponse is an object.
type PagedDataSourceResponse struct {
	// Data:
	Data []DataSourceResponse `json:"data"`
	// Fieldspec:
	Fieldspec []FieldSpec `json:"fieldspec"`
	// Page: Contains the pagination metadata for a response
	Page PageInfo `json:"page"`
}

// GetData returns the Data property
func (m PagedDataSourceResponse) GetData() []DataSourceResponse {
	return m.Data
}

// SetData sets the Data property
func (m PagedDataSourceResponse) SetData(val []DataSourceResponse) {
	m.Data = val
}

// GetFieldspec returns the Fieldspec property
func (m PagedDataSourceResponse) GetFieldspec() []FieldSpec {
	return m.Fieldspec
}

// SetFieldspec sets the Fieldspec property
func (m PagedDataSourceResponse) SetFieldspec(val []FieldSpec) {
	m.Fieldspec = val
}

// GetPage returns the Page property
func (m PagedDataSourceResponse) GetPage() PageInfo {
	return m.Page
}

// SetPage sets the Page property
func (m PagedDataSourceResponse) SetPage(val PageInfo) {
	m.Page = val
}
