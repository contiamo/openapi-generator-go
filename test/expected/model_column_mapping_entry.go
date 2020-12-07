// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// ColumnMappingEntry is an object.
type ColumnMappingEntry struct {
	// Source: Name of the source column
	Source string `json:"source"`
	// Target: Name of the target column
	Target string `json:"target"`
}

// GetSource returns the Source property
func (m ColumnMappingEntry) GetSource() string {
	return m.Source
}

// SetSource sets the Source property
func (m ColumnMappingEntry) SetSource(val string) {
	m.Source = val
}

// GetTarget returns the Target property
func (m ColumnMappingEntry) GetTarget() string {
	return m.Target
}

// SetTarget sets the Target property
func (m ColumnMappingEntry) SetTarget(val string) {
	m.Target = val
}
