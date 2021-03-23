// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Foo is an object.
type Foo struct {
	// A: property
	A string `json:"a,omitempty"`
}

// Validate implements basic validation for this model
func (m Foo) Validate() error {
	return validation.Errors{}.Filter()
}

// GetA returns the A property
func (m Foo) GetA() string {
	return m.A
}

// SetA sets the A property
func (m Foo) SetA(val string) {
	m.A = val
}
