// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package testpkg

// SubType is an object.
type SubType struct {
	// Foo:
	Foo string `json:"foo,omitempty"`
}

// GetFoo returns the Foo property
func (m SubType) GetFoo() string {
	return m.Foo
}

// SetFoo sets the Foo property
func (m SubType) SetFoo(val string) {
	m.Foo = val
}
