// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: Test oneOf discriminator support
//	Version: 0.1.0
package generatortest

import (
	"encoding/json"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/mitchellh/mapstructure"
)

// ErrorDiscriminator is the discriminator value used during serialization and validation.
// Retrieve this value using  the ErrorDiscriminator() method
type ErrorDiscriminator string

const (
	// ErrorDiscriminatorAuth maps Error to GenericError
	ErrorDiscriminatorAuth ErrorDiscriminator = "auth"
	// ErrorDiscriminatorExternal maps Error to ExternalError
	ErrorDiscriminatorExternal ErrorDiscriminator = "external"
	// ErrorDiscriminatorField maps Error to FieldError
	ErrorDiscriminatorField ErrorDiscriminator = "field"
	// ErrorDiscriminatorGeneric maps Error to GenericError
	ErrorDiscriminatorGeneric ErrorDiscriminator = "generic"
)

type errorer interface {
	ErrorDiscriminator() ErrorDiscriminator
	Validate() error
}

// Error is a oneOf type.
type Error struct {
	data errorer
}

var EmptyErrorError = fmt.Errorf("empty data is not an Error")
var NotErrorError = fmt.Errorf("could not convert to type Error")

type ErrorNilableRule struct{}

func (nn ErrorNilableRule) Validate(v interface{}) error {
	if m, ok := v.(Error); !ok {
		return NotErrorError
	} else if m.data == nil {
		return nil
	} else {
		return m.Validate()
	}
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

	value := fmt.Sprintf("%v", discriminator.Value)
	switch ErrorDiscriminator(value) {
	case ErrorDiscriminatorAuth:
		data := GenericError{}
		err := json.Unmarshal(bs, &data)
		if err != nil {
			return err
		}
		m.data = data
	case ErrorDiscriminatorExternal:
		data := ExternalError{}
		err := json.Unmarshal(bs, &data)
		if err != nil {
			return err
		}
		m.data = data
	case ErrorDiscriminatorField:
		data := FieldError{}
		err := json.Unmarshal(bs, &data)
		if err != nil {
			return err
		}
		m.data = data
	case ErrorDiscriminatorGeneric:
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
	if m.data == nil {
		return EmptyErrorError
	}
	return mapstructure.Decode(m.data, target)
}

// AsGenericError converts Error to a GenericError
func (m Error) AsGenericError() (result GenericError, err error) {
	if m.data == nil {
		return result, EmptyErrorError
	}
	return result, mapstructure.Decode(m.data, &result)
}

// AsFieldError converts Error to a FieldError
func (m Error) AsFieldError() (result FieldError, err error) {
	if m.data == nil {
		return result, EmptyErrorError
	}
	return result, mapstructure.Decode(m.data, &result)
}

// AsExternalError converts Error to a ExternalError
func (m Error) AsExternalError() (result ExternalError, err error) {
	if m.data == nil {
		return result, EmptyErrorError
	}
	return result, mapstructure.Decode(m.data, &result)
}

// Discriminator returns the ErrorDiscriminator oneOf value
func (m Error) Discriminator() ErrorDiscriminator {
	return m.data.ErrorDiscriminator()
}

// Validate implements basic validation for this model
func (m Error) Validate() error {
	if m.data == nil {
		return EmptyErrorError
	}
	discriminator := m.data.ErrorDiscriminator()
	switch discriminator {
	case ErrorDiscriminatorAuth:
		return m.data.Validate()
	case ErrorDiscriminatorExternal:
		return m.data.Validate()
	case ErrorDiscriminatorField:
		return m.data.Validate()
	case ErrorDiscriminatorGeneric:
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
	case ErrorDiscriminatorAuth:
		return true
	case ErrorDiscriminatorExternal:
		return true
	case ErrorDiscriminatorField:
		return true
	case ErrorDiscriminatorGeneric:
		return true
	default:
		return false
	}
}

// ErrorDiscriminator implements errorer and returns the discriminator value.
func (m GenericError) ErrorDiscriminator() ErrorDiscriminator {
	value := fmt.Sprintf("%v", m.GetKind())
	return ErrorDiscriminator(value)
}

// ErrorDiscriminator implements errorer and returns the discriminator value.
func (m FieldError) ErrorDiscriminator() ErrorDiscriminator {
	value := fmt.Sprintf("%v", m.GetKind())
	return ErrorDiscriminator(value)
}

// ErrorDiscriminator implements errorer and returns the discriminator value.
func (m ExternalError) ErrorDiscriminator() ErrorDiscriminator {
	value := fmt.Sprintf("%v", m.GetKind())
	return ErrorDiscriminator(value)
}
