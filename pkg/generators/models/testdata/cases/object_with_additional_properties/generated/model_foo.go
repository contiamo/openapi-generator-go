// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: Test
//	Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Foo is an object.
type Foo struct {
	// Foo:
	Foo map[string]string `json:"foo,omitempty"`
}

// Validate implements basic validation for this model
func (m Foo) Validate() error {
	return validation.Errors{
		"foo": validation.Validate(
			m.Foo,
		),
	}.Filter()
}

// GetFoo returns the Foo property
func (m Foo) GetFoo() map[string]string {
	return m.Foo
}

// SetFoo sets the Foo property
func (m *Foo) SetFoo(val map[string]string) {
	m.Foo = val
}
