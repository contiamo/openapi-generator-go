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

// FooBar is an object.
type FooBar struct {
	// Baz:
	Baz string `json:"baz" mapstructure:"baz"`
}

// Validate implements basic validation for this model
func (m FooBar) Validate() error {
	return validation.Errors{}.Filter()
}

// GetBaz returns the Baz property
func (m FooBar) GetBaz() string {
	return m.Baz
}

// SetBaz sets the Baz property
func (m *FooBar) SetBaz(val string) {
	m.Baz = val
}
