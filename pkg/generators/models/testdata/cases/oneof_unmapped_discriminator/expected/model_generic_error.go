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

// GenericError is an object. Represents and unknown error. The system may be unstable, it is unknown if retrying the request will succeed.
type GenericError struct {
	// Code: machine friendly error code.
	Code string `json:"code,omitempty" mapstructure:"code,omitempty"`
	// Kind:
	Kind string `json:"kind" mapstructure:"kind"`
	// Message: the user friendly error message.
	Message string `json:"message" mapstructure:"message"`
	// TraceId: the request tracing id, this can be submitted during bug reports to help with debugging the underlying cause.
	TraceId string `json:"traceId,omitempty" mapstructure:"traceId,omitempty"`
}

// Validate implements basic validation for this model
func (m GenericError) Validate() error {
	return validation.Errors{}.Filter()
}

// GetCode returns the Code property
func (m GenericError) GetCode() string {
	return m.Code
}

// SetCode sets the Code property
func (m *GenericError) SetCode(val string) {
	m.Code = val
}

// GetKind returns the Kind property
func (m GenericError) GetKind() string {
	return m.Kind
}

// SetKind sets the Kind property
func (m *GenericError) SetKind(val string) {
	m.Kind = val
}

// GetMessage returns the Message property
func (m GenericError) GetMessage() string {
	return m.Message
}

// SetMessage sets the Message property
func (m *GenericError) SetMessage(val string) {
	m.Message = val
}

// GetTraceId returns the TraceId property
func (m GenericError) GetTraceId() string {
	return m.TraceId
}

// SetTraceId sets the TraceId property
func (m *GenericError) SetTraceId(val string) {
	m.TraceId = val
}
