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

// Foo is an object.
type Foo struct {
	// Foo:
	Foo FooFoo `json:"foo,omitempty" mapstructure:"foo,omitempty"`
}

// NewFoo instantiates a new Foo with default values overriding them as follows:
// 1. Default values specified in the Foo schema
// 2. Default values specified per Foo property
func NewFoo() *Foo {
	m := &Foo{}

	return m
}

// UnmarshalJSON implements the json.Unmarshaler interface for Foo. It set the default values for the Foo type
func (m *Foo) UnmarshalJSON(data []byte) error {
	// Set default values
	*m = *NewFoo()

	// Unmarshal using an alias to avoid an infinite loop
	type alias Foo
	err := json.Unmarshal(data, (*alias)(m))
	if err != nil {
		return err
	}
	return nil
}

// Validate implements basic validation for this model
func (m Foo) Validate() error {
	errors := validation.Errors{
		"foo": validation.Validate(
			m.Foo,
		),
	}
	return errors.Filter()
}

// GetFoo returns the Foo property
func (m Foo) GetFoo() FooFoo {
	return m.Foo
}

// SetFoo sets the Foo property
func (m *Foo) SetFoo(val FooFoo) {
	m.Foo = val
}
