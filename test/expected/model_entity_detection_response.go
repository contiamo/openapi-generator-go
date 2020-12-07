// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// EntityDetectionResponse is an object.
type EntityDetectionResponse struct {
	// Columns:
	Columns interface{} `json:"columns"`
	// Metadata:
	Metadata Metadata `json:"metadata"`
}

// GetColumns returns the Columns property
func (m EntityDetectionResponse) GetColumns() interface{} {
	return m.Columns
}

// SetColumns sets the Columns property
func (m EntityDetectionResponse) SetColumns(val interface{}) {
	m.Columns = val
}

// GetMetadata returns the Metadata property
func (m EntityDetectionResponse) GetMetadata() Metadata {
	return m.Metadata
}

// SetMetadata sets the Metadata property
func (m EntityDetectionResponse) SetMetadata(val Metadata) {
	m.Metadata = val
}
