// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test oneOf discriminator support
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// FieldError is an object. Represents a validation error. The request must be corrected before retrying.
type FieldError struct {
	// Errors:
	Errors []ValidationError `json:"errors"`
	// Kind:
	Kind string `json:"kind"`
	// TraceId: the request tracing id, this can be submitted during bug reports to help with debugging the underlying cause.
	TraceId string `json:"traceId"`
}

// Validate implements basic validation for this model
func (m FieldError) Validate() error {
	return validation.Errors{
		"errors": validation.Validate(
			m.Errors, validation.NotNil,
		),
	}.Filter()
}

// GetErrors returns the Errors property
func (m FieldError) GetErrors() []ValidationError {
	return m.Errors
}

// SetErrors sets the Errors property
func (m *FieldError) SetErrors(val []ValidationError) {
	m.Errors = val
}

// GetKind returns the Kind property
func (m FieldError) GetKind() string {
	return m.Kind
}

// SetKind sets the Kind property
func (m *FieldError) SetKind(val string) {
	m.Kind = val
}

// GetTraceId returns the TraceId property
func (m FieldError) GetTraceId() string {
	return m.TraceId
}

// SetTraceId sets the TraceId property
func (m *FieldError) SetTraceId(val string) {
	m.TraceId = val
}
