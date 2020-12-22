// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Top is an object.
type Top struct {
	// Bar:
	Bar Sub `json:"bar,omitempty"`
}

// Validate implements basic validation for this model
func (m Top) Validate() error {
	return validation.Errors{
		"bar": validation.Validate(
			m.Bar, validation.Required,
		),
	}.Filter()
}

// GetBar returns the Bar property
func (m Top) GetBar() Sub {
	return m.Bar
}

// SetBar sets the Bar property
func (m Top) SetBar(val Sub) {
	m.Bar = val
}
