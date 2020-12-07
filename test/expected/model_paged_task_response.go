// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// PagedTaskResponse is an object.
type PagedTaskResponse struct {
	// Data:
	Data []TaskResponse `json:"data"`
	// Page: Contains the pagination metadata for a response
	Page PageInfo `json:"page"`
}

// GetData returns the Data property
func (m PagedTaskResponse) GetData() []TaskResponse {
	return m.Data
}

// SetData sets the Data property
func (m PagedTaskResponse) SetData(val []TaskResponse) {
	m.Data = val
}

// GetPage returns the Page property
func (m PagedTaskResponse) GetPage() PageInfo {
	return m.Page
}

// SetPage sets the Page property
func (m PagedTaskResponse) SetPage(val PageInfo) {
	m.Page = val
}
