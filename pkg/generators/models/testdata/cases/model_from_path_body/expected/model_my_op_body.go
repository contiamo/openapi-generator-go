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

// MyOpBody is an object.
type MyOpBody struct {
	// Foo:
	Foo string `json:"foo,omitempty"`
}

// Validate implements basic validation for this model
func (m MyOpBody) Validate() error {
	return validation.Errors{}.Filter()
}

// GetFoo returns the Foo property
func (m MyOpBody) GetFoo() string {
	return m.Foo
}

// SetFoo sets the Foo property
func (m *MyOpBody) SetFoo(val string) {
	m.Foo = val
}
