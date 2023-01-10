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

// ValidationError is an object. a validation error for a form field
type ValidationError struct {
	// Message: the user friendly validation error message
	Message string `json:"message,omitempty" mapstructure:"message,omitempty"`
	// Name: the field name
	Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
}

// Validate implements basic validation for this model
func (m ValidationError) Validate() error {
	return validation.Errors{}.Filter()
}

// GetMessage returns the Message property
func (m ValidationError) GetMessage() string {
	return m.Message
}

// SetMessage sets the Message property
func (m *ValidationError) SetMessage(val string) {
	m.Message = val
}

// GetName returns the Name property
func (m ValidationError) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *ValidationError) SetName(val string) {
	m.Name = val
}
