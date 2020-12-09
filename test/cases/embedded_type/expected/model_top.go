// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test
//     Version: 0.1.0
package generatortest

// Top is an object.
type Top struct {
	// Bar:
	Bar Sub `json:"bar,omitempty"`
}

// GetBar returns the Bar property
func (m Top) GetBar() Sub {
	return m.Bar
}

// SetBar sets the Bar property
func (m Top) SetBar(val Sub) {
	m.Bar = val
}
