// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: Test oneOf discriminator support
//	Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Container is an object.
type Container struct {
	// Error:
	Error Error `json:"error,omitempty"`
}

// Validate implements basic validation for this model
func (m Container) Validate() error {
	return validation.Errors{
		"error": validation.Validate(
			m.Error,
		),
	}.Filter()
}

// GetError returns the Error property
func (m Container) GetError() Error {
	return m.Error
}

// SetError sets the Error property
func (m *Container) SetError(val Error) {
	m.Error = val
}
