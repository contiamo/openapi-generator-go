// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// FullConnectionsListResponse is an object.
type FullConnectionsListResponse struct {
	// Data:
	Data []FullConnectionsResponse `json:"data"`
}

// GetData returns the Data property
func (m FullConnectionsListResponse) GetData() []FullConnectionsResponse {
	return m.Data
}

// SetData sets the Data property
func (m FullConnectionsListResponse) SetData(val []FullConnectionsResponse) {
	m.Data = val
}
