// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test
//     Version: 0.1.0
package generatortest

// Foo is an object.
type Foo struct {
	// Foo:
	Foo map[string]string `json:"foo,omitempty"`
}

// GetFoo returns the Foo property
func (m Foo) GetFoo() map[string]string {
	return m.Foo
}

// SetFoo sets the Foo property
func (m Foo) SetFoo(val map[string]string) {
	m.Foo = val
}
