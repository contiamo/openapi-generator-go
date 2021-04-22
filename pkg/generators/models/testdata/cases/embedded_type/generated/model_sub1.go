// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Sub1 is an object.
type Sub1 struct {
	// Foo:
	Foo string `json:"foo,omitempty"`
}

// Validate implements basic validation for this model
func (m Sub1) Validate() error {
	return validation.Errors{}.Filter()
}

// GetFoo returns the Foo property
func (m Sub1) GetFoo() string {
	return m.Foo
}

// SetFoo sets the Foo property
func (m Sub1) SetFoo(val string) {
	m.Foo = val
}
