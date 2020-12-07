// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test Service
//     Version: 0.1.0
package api

// ObjectWithNullableReference is an object.
type ObjectWithNullableReference struct {
	// Foo:
	Foo *SubType `json:"foo,omitempty"`
}

// GetFoo returns the Foo property
func (m ObjectWithNullableReference) GetFoo() *SubType {
	return m.Foo
}

// SetFoo sets the Foo property
func (m ObjectWithNullableReference) SetFoo(val *SubType) {
	m.Foo = val
}
