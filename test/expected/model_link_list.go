// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// LinkList is an object.
type LinkList struct {
	// Data:
	Data []LinkResponse `json:"data"`
	// Page: Contains the pagination metadata for a response
	Page PageInfo `json:"page"`
}

// GetData returns the Data property
func (m LinkList) GetData() []LinkResponse {
	return m.Data
}

// SetData sets the Data property
func (m LinkList) SetData(val []LinkResponse) {
	m.Data = val
}

// GetPage returns the Page property
func (m LinkList) GetPage() PageInfo {
	return m.Page
}

// SetPage sets the Page property
func (m LinkList) SetPage(val PageInfo) {
	m.Page = val
}
