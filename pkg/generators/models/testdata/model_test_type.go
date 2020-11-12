// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package testpkg

// TestType is an object.
type TestType struct {
	// Bar
	Bar string `json:"bar,omitempty"`
	// Baz
	Baz []SubType `json:"baz,omitempty"`
	// Foo
	Foo int `json:"foo,omitempty"`
}

// GetBar returns the Bar property
func (m TestType) GetBar() string {
	return m.Bar
}

// SetBar sets the Bar property
func (m TestType) SetBar(val string) {
	m.Bar = val
}

// GetBaz returns the Baz property
func (m TestType) GetBaz() []SubType {
	return m.Baz
}

// SetBaz sets the Baz property
func (m TestType) SetBaz(val []SubType) {
	m.Baz = val
}

// GetFoo returns the Foo property
func (m TestType) GetFoo() int {
	return m.Foo
}

// SetFoo sets the Foo property
func (m TestType) SetFoo(val int) {
	m.Foo = val
}
