// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// PagedQueryHistoryResponse is an object.
type PagedQueryHistoryResponse struct {
	// Data:
	Data []QueryHistoryResponse `json:"data"`
	// Page: Contains the pagination metadata for a response
	Page PageInfo `json:"page"`
}

// GetData returns the Data property
func (m PagedQueryHistoryResponse) GetData() []QueryHistoryResponse {
	return m.Data
}

// SetData sets the Data property
func (m PagedQueryHistoryResponse) SetData(val []QueryHistoryResponse) {
	m.Data = val
}

// GetPage returns the Page property
func (m PagedQueryHistoryResponse) GetPage() PageInfo {
	return m.Page
}

// SetPage sets the Page property
func (m PagedQueryHistoryResponse) SetPage(val PageInfo) {
	m.Page = val
}
