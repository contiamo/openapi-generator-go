// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// QueryResponse is an object.
type QueryResponse struct {
	// Columns:
	Columns []QueryResponseMetadata `json:"columns"`
	// Rows: 2D array of values
	Rows [][]*interface{} `json:"rows"`
}

// GetColumns returns the Columns property
func (m QueryResponse) GetColumns() []QueryResponseMetadata {
	return m.Columns
}

// SetColumns sets the Columns property
func (m QueryResponse) SetColumns(val []QueryResponseMetadata) {
	m.Columns = val
}

// GetRows returns the Rows property
func (m QueryResponse) GetRows() [][]*interface{} {
	return m.Rows
}

// SetRows sets the Rows property
func (m QueryResponse) SetRows(val [][]*interface{}) {
	m.Rows = val
}
