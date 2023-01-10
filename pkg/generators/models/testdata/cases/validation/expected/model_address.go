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

// Address is an object.
type Address struct {
	// Name:
	Name *string `json:"name" mapstructure:"name"`
	// Number:
	Number int32 `json:"number" mapstructure:"number"`
	// Street:
	Street string `json:"street" mapstructure:"street"`
}

// Validate implements basic validation for this model
func (m Address) Validate() error {
	return validation.Errors{
		"name": validation.Validate(
			m.Name, validation.Length(2, 0),
		),
		"street": validation.Validate(
			m.Street, validation.Required, validation.Length(2, 0),
		),
	}.Filter()
}

// GetName returns the Name property
func (m Address) GetName() *string {
	return m.Name
}

// SetName sets the Name property
func (m *Address) SetName(val *string) {
	m.Name = val
}

// GetNumber returns the Number property
func (m Address) GetNumber() int32 {
	return m.Number
}

// SetNumber sets the Number property
func (m *Address) SetNumber(val int32) {
	m.Number = val
}

// GetStreet returns the Street property
func (m Address) GetStreet() string {
	return m.Street
}

// SetStreet sets the Street property
func (m *Address) SetStreet(val string) {
	m.Street = val
}
