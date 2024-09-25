// Code generated by openapi-generator-go DO NOT EDIT.
//
// Source:
//
//	Title: Test oneOf discriminator support
//	Version: 0.1.0
package generatortest

import (
	"encoding/json"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// ExternalError is an object. Represents an error from an external service. The system is stable, but can not complete the request. Try again later.
type ExternalError struct {
	// Kind:
	Kind ExternalErrorKind `json:"kind" mapstructure:"kind"`
	// Service: name of the service that is returning errors
	Service string `json:"service,omitempty" mapstructure:"service,omitempty"`
	// TraceId: the request tracing id, this can be submitted during bug reports to help with debugging the underlying cause.
	TraceId string `json:"traceId" mapstructure:"traceId"`
}

// NewExternalError instantiates a new ExternalError with default values overriding them as follows:
// 1. Default values specified in the ExternalError schema
// 2. Default values specified per ExternalError property
func NewExternalError() *ExternalError {
	m := &ExternalError{}

	return m
}

// UnmarshalJSON implements the json.Unmarshaler interface for ExternalError. It set the default values for the ExternalError type
func (m *ExternalError) UnmarshalJSON(data []byte) error {
	// Set default values
	*m = *NewExternalError()

	// Unmarshal using an alias to avoid an infinite loop
	type alias ExternalError
	err := json.Unmarshal(data, (*alias)(m))
	if err != nil {
		return err
	}
	return nil
}

// Validate implements basic validation for this model
func (m ExternalError) Validate() error {
	errors := validation.Errors{
		"kind": validation.Validate(
			m.Kind, validation.Required,
		),
	}
	return errors.Filter()
}

// GetKind returns the Kind property
func (m ExternalError) GetKind() ExternalErrorKind {
	return m.Kind
}

// SetKind sets the Kind property
func (m *ExternalError) SetKind(val ExternalErrorKind) {
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
