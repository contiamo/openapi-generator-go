// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// ListConnectionsResponse is an object.
type ListConnectionsResponse struct {
	// Data: A list of connections for a given data source.
	Data []Connection `json:"data"`
}

// GetData returns the Data property
func (m ListConnectionsResponse) GetData() []Connection {
	return m.Data
}

// SetData sets the Data property
func (m ListConnectionsResponse) SetData(val []Connection) {
	m.Data = val
}
