// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// HsqldbFoodmartDataSourceConnectionInfo is an object.
type HsqldbFoodmartDataSourceConnectionInfo struct {
	// MaxDirectConnections: Max numbers of connections that can be used for direct queries. `20` by default.
	MaxDirectConnections int32 `json:"maxDirectConnections,omitempty"`
	// MaxEtlConnections: Max number of connections that can be used for ETL queries. `20` by default.
	MaxEtlConnections int32 `json:"maxEtlConnections,omitempty"`
}

// GetMaxDirectConnections returns the MaxDirectConnections property
func (m HsqldbFoodmartDataSourceConnectionInfo) GetMaxDirectConnections() int32 {
	return m.MaxDirectConnections
}

// SetMaxDirectConnections sets the MaxDirectConnections property
func (m HsqldbFoodmartDataSourceConnectionInfo) SetMaxDirectConnections(val int32) {
	m.MaxDirectConnections = val
}

// GetMaxEtlConnections returns the MaxEtlConnections property
func (m HsqldbFoodmartDataSourceConnectionInfo) GetMaxEtlConnections() int32 {
	return m.MaxEtlConnections
}

// SetMaxEtlConnections sets the MaxEtlConnections property
func (m HsqldbFoodmartDataSourceConnectionInfo) SetMaxEtlConnections(val int32) {
	m.MaxEtlConnections = val
}
