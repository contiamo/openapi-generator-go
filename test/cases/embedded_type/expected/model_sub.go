// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test
//     Version: 0.1.0
package generatortest

// Sub is an object.
type Sub struct {
	// Foo:
	Foo string `json:"foo,omitempty"`
}

// GetFoo returns the Foo property
func (m Sub) GetFoo() string {
	return m.Foo
}

// SetFoo sets the Foo property
func (m Sub) SetFoo(val string) {
	m.Foo = val
}
