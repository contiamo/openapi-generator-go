// Code generated by openapi-generator-go DO NOT EDIT.
//
// Source:
//
//	Title: Test
//	Version: 0.1.0
package generatortest

import (
	"encoding/json"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// EitherIdOrNamespaceNameOneOfIdx1 is an object.
type EitherIdOrNamespaceNameOneOfIdx1 struct {
	// Name:
	Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// Namespace:
	Namespace string `json:"namespace,omitempty" mapstructure:"namespace,omitempty"`
}

// NewEitherIdOrNamespaceNameOneOfIdx1 instantiates a new EitherIdOrNamespaceNameOneOfIdx1 with default values overriding them as follows:
// 1. Default values specified in the EitherIdOrNamespaceNameOneOfIdx1 schema
// 2. Default values specified per EitherIdOrNamespaceNameOneOfIdx1 property
func NewEitherIdOrNamespaceNameOneOfIdx1() *EitherIdOrNamespaceNameOneOfIdx1 {
	m := &EitherIdOrNamespaceNameOneOfIdx1{}

	return m
}

// UnmarshalJSON implements the json.Unmarshaler interface for EitherIdOrNamespaceNameOneOfIdx1. It set the default values for the EitherIdOrNamespaceNameOneOfIdx1 type
func (m *EitherIdOrNamespaceNameOneOfIdx1) UnmarshalJSON(data []byte) error {
	// Set default values
	*m = *NewEitherIdOrNamespaceNameOneOfIdx1()

	// Unmarshal using an alias to avoid an infinite loop
	type alias EitherIdOrNamespaceNameOneOfIdx1
	err := json.Unmarshal(data, (*alias)(m))
	if err != nil {
		return err
	}
	return nil
}

// Validate implements basic validation for this model
func (m EitherIdOrNamespaceNameOneOfIdx1) Validate() error {
	errors := validation.Errors{}
	return errors.Filter()
}

// GetName returns the Name property
func (m EitherIdOrNamespaceNameOneOfIdx1) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *EitherIdOrNamespaceNameOneOfIdx1) SetName(val string) {
	m.Name = val
}

// GetNamespace returns the Namespace property
func (m EitherIdOrNamespaceNameOneOfIdx1) GetNamespace() string {
	return m.Namespace
}

// SetNamespace sets the Namespace property
func (m *EitherIdOrNamespaceNameOneOfIdx1) SetNamespace(val string) {
	m.Namespace = val
}