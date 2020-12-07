// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// PagedSearchResponse is an object.
type PagedSearchResponse struct {
	// Data:
	Data []DataSourceSearchResponse `json:"data"`
	// DataSourceFieldspec:
	DataSourceFieldspec []FieldSpec `json:"dataSourceFieldspec"`
	// Page: Contains the pagination metadata for a response
	Page PageInfo `json:"page"`
	// TableFieldspec:
	TableFieldspec []FieldSpec `json:"tableFieldspec"`
}

// GetData returns the Data property
func (m PagedSearchResponse) GetData() []DataSourceSearchResponse {
	return m.Data
}

// SetData sets the Data property
func (m PagedSearchResponse) SetData(val []DataSourceSearchResponse) {
	m.Data = val
}

// GetDataSourceFieldspec returns the DataSourceFieldspec property
func (m PagedSearchResponse) GetDataSourceFieldspec() []FieldSpec {
	return m.DataSourceFieldspec
}

// SetDataSourceFieldspec sets the DataSourceFieldspec property
func (m PagedSearchResponse) SetDataSourceFieldspec(val []FieldSpec) {
	m.DataSourceFieldspec = val
}

// GetPage returns the Page property
func (m PagedSearchResponse) GetPage() PageInfo {
	return m.Page
}

// SetPage sets the Page property
func (m PagedSearchResponse) SetPage(val PageInfo) {
	m.Page = val
}

// GetTableFieldspec returns the TableFieldspec property
func (m PagedSearchResponse) GetTableFieldspec() []FieldSpec {
	return m.TableFieldspec
}

// SetTableFieldspec sets the TableFieldspec property
func (m PagedSearchResponse) SetTableFieldspec(val []FieldSpec) {
	m.TableFieldspec = val
}
