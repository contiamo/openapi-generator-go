// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// FullConnectionsResponse is an object.
type FullConnectionsResponse struct {
	// Connections:
	Connections []Connection `json:"connections"`
	// DataSource:
	DataSource ExternalDataSourceSummaryResponse `json:"dataSource"`
}

// GetConnections returns the Connections property
func (m FullConnectionsResponse) GetConnections() []Connection {
	return m.Connections
}

// SetConnections sets the Connections property
func (m FullConnectionsResponse) SetConnections(val []Connection) {
	m.Connections = val
}

// GetDataSource returns the DataSource property
func (m FullConnectionsResponse) GetDataSource() ExternalDataSourceSummaryResponse {
	return m.DataSource
}

// SetDataSource sets the DataSource property
func (m FullConnectionsResponse) SetDataSource(val ExternalDataSourceSummaryResponse) {
	m.DataSource = val
}
