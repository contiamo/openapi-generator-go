// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test oneOf discriminator support
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// ExternalError is an object. Represents an error from an external service. The system is stable, but can not complete the request. Try again later.
type ExternalError struct {
	// Kind:
	Kind string `json:"kind"`
	// Service: name of the service that is returning errors
	Service string `json:"service,omitempty"`
	// TraceId: the request tracing id, this can be submitted during bug reports to help with debugging the underlying cause.
	TraceId string `json:"traceId"`
}

// Validate implements basic validation for this model
func (m ExternalError) Validate() error {
	return validation.Errors{}.Filter()
}

// GetKind returns the Kind property
func (m ExternalError) GetKind() string {
	return m.Kind
}

// SetKind sets the Kind property
func (m *ExternalError) SetKind(val string) {
	m.Kind = val
}

// GetService returns the Service property
func (m ExternalError) GetService() string {
	return m.Service
}

// SetService sets the Service property
func (m *ExternalError) SetService(val string) {
	m.Service = val
}

// GetTraceId returns the TraceId property
func (m ExternalError) GetTraceId() string {
	return m.TraceId
}

// SetTraceId sets the TraceId property
func (m *ExternalError) SetTraceId(val string) {
	m.TraceId = val
}
