// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test
//     Version: 0.1.0
package generatortest

// Foo is an object.
type Foo struct {
	// Bar:
	Bar struct {
		Bar string `json:"bar,omitempty"`
		Foo string `json:"foo,omitempty"`
	} `json:"bar,omitempty"`
}

// GetBar returns the Bar property
func (m Foo) GetBar() struct {
	Bar string `json:"bar,omitempty"`
	Foo string `json:"foo,omitempty"`
} {
	return m.Bar
}

// SetBar sets the Bar property
func (m Foo) SetBar(val struct {
	Bar string `json:"bar,omitempty"`
	Foo string `json:"foo,omitempty"`
}) {
	m.Bar = val
}
