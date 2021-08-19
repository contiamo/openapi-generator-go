// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test oneOf discriminator support
//     Version: 0.1.0
package generatortest

import (
	"encoding/json"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/mitchellh/mapstructure"
)

type errorer interface {
	ErrorDiscriminator() string
	Validate() error
}

// Error is a oneOf type.
type Error struct {
	data errorer
}

// MarshalJSON implements the json.Marshaller interface
func (m Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.data)
}

// UnmarshalJSON implements the json.Unmarshaller interface
func (m *Error) UnmarshalJSON(bs []byte) error {
	discriminator := struct {
		Value string `json:"kind"`
	}{}
	err := json.Unmarshal(bs, &discriminator)
	if err != nil {
		return err
	}

	if discriminator.Value == "" {
		return validation.Errors{
			"kind": fmt.Errorf("can not be empty"),
		}.Filter()
	}

	switch discriminator.Value {
	case "auth":
		data := GenericError{}
		err := json.Unmarshal(bs, &data)
		if err != nil {
			return err
		}
		m.data = data
	case "external":
		data := ExternalError{}
		err := json.Unmarshal(bs, &data)
		if err != nil {
			return err
		}
		m.data = data
	case "field":
		data := FieldError{}
		err := json.Unmarshal(bs, &data)
		if err != nil {
			return err
		}
		m.data = data
	case "generic":
		data := GenericError{}
		err := json.Unmarshal(bs, &data)
		if err != nil {
			return err
		}
		m.data = data
	default:
		return validation.Errors{
			"kind": fmt.Errorf("invalid value"),
		}.Filter()
	}
	return nil
}

// FromGenericError sets the Error data.
func (m *Error) FromGenericError(data GenericError) {
	m.data = data
}

// FromFieldError sets the Error data.
func (m *Error) FromFieldError(data FieldError) {
	m.data = data
}

// FromExternalError sets the Error data.
func (m *Error) FromExternalError(data ExternalError) {
	m.data = data
}

// As converts Error to a user defined structure.
func (m Error) As(target interface{}) error {
	return mapstructure.Decode(m.data, target)
}

// AsGenericError converts Error to a GenericError
func (m Error) AsGenericError() (result GenericError, err error) {
	return result, mapstructure.Decode(m.data, &result)
}

// AsFieldError converts Error to a FieldError
func (m Error) AsFieldError() (result FieldError, err error) {
	return result, mapstructure.Decode(m.data, &result)
}

// AsExternalError converts Error to a ExternalError
func (m Error) AsExternalError() (result ExternalError, err error) {
	return result, mapstructure.Decode(m.data, &result)
}

// Validate implements basic validation for this model
func (m Error) Validate() error {
	discriminator := m.data.ErrorDiscriminator()
	switch discriminator {
	case "auth":
		return m.data.Validate()
	case "external":
		return m.data.Validate()
	case "field":
		return m.data.Validate()
	case "generic":
		return m.data.Validate()
	default:
		return validation.Errors{
			"kind": fmt.Errorf("unknown kind value"),
		}.Filter()
	}
}

// IsError tests if data is one of the discriminated sub-types of Error.
func IsError(data interface{}) bool {
	t, ok := data.(errorer)
	if !ok {
		return false
	}

	discriminator := t.ErrorDiscriminator()
	switch discriminator {
	case "auth":
		return true
	case "external":
		return true
	case "field":
		return true
	case "generic":
		return true
	default:
		return false
	}
}

// ErrorDiscriminator implements errorer and returns the discriminator value as a string.
func (m GenericError) ErrorDiscriminator() string {
	return string(m.GetKind())
}

// ErrorDiscriminator implements errorer and returns the discriminator value as a string.
func (m FieldError) ErrorDiscriminator() string {
	return string(m.GetKind())
}

// ErrorDiscriminator implements errorer and returns the discriminator value as a string.
func (m ExternalError) ErrorDiscriminator() string {
	return string(m.GetKind())
}
