// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// PagedCustomFieldResponse is an object.
type PagedCustomFieldResponse struct {
	// Data:
	Data []FieldSpec `json:"data"`
	// Page: Contains the pagination metadata for a response
	Page PageInfo `json:"page"`
}

// GetData returns the Data property
func (m PagedCustomFieldResponse) GetData() []FieldSpec {
	return m.Data
}

// SetData sets the Data property
func (m PagedCustomFieldResponse) SetData(val []FieldSpec) {
	m.Data = val
}

// GetPage returns the Page property
func (m PagedCustomFieldResponse) GetPage() PageInfo {
	return m.Page
}

// SetPage sets the Page property
func (m PagedCustomFieldResponse) SetPage(val PageInfo) {
	m.Page = val
}
