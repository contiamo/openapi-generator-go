// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test
//     Version: 0.1.0
package generatortest

// Foo is an object.
type Foo struct {
	// Bar:
	Bar string `json:"bar,omitempty"`
	// Foo:
	Foo string `json:"foo,omitempty"`
}

// GetBar returns the Bar property
func (m Foo) GetBar() string {
	return m.Bar
}

// SetBar sets the Bar property
func (m Foo) SetBar(val string) {
	m.Bar = val
}

// GetFoo returns the Foo property
func (m Foo) GetFoo() string {
	return m.Foo
}

// SetFoo sets the Foo property
func (m Foo) SetFoo(val string) {
	m.Foo = val
}