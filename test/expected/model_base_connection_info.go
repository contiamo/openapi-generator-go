// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// BaseConnectionInfo is an object. Contains parameters for connecting to the data source.
// If a parameter is not required and not set it will be replaced by the documented default value if any.
type BaseConnectionInfo struct {
	// MaxDirectConnections: Max numbers of connections that can be used for direct queries. `20` by default.
	MaxDirectConnections int32 `json:"maxDirectConnections,omitempty"`
	// MaxEtlConnections: Max number of connections that can be used for ETL queries. `20` by default.
	MaxEtlConnections int32 `json:"maxEtlConnections,omitempty"`
}

// GetMaxDirectConnections returns the MaxDirectConnections property
func (m BaseConnectionInfo) GetMaxDirectConnections() int32 {
	return m.MaxDirectConnections
}

// SetMaxDirectConnections sets the MaxDirectConnections property
func (m BaseConnectionInfo) SetMaxDirectConnections(val int32) {
	m.MaxDirectConnections = val
}

// GetMaxEtlConnections returns the MaxEtlConnections property
func (m BaseConnectionInfo) GetMaxEtlConnections() int32 {
	return m.MaxEtlConnections
}

// SetMaxEtlConnections sets the MaxEtlConnections property
func (m BaseConnectionInfo) SetMaxEtlConnections(val int32) {
	m.MaxEtlConnections = val
}
