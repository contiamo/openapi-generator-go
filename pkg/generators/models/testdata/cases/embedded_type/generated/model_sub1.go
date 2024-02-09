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

// Sub1 is an object.
type Sub1 struct {
	// Foo:
	Foo string `json:"foo,omitempty" mapstructure:"foo,omitempty"`
}

// NewSub1 instantiates a new Sub1 with default values overriding them as follows:
// 1. Default values specified in the Sub1 schema
// 2. Default values specified per Sub1 property
func NewSub1() *Sub1 {
	m := &Sub1{}

	return m
}

// UnmarshalJSON implements the json.Unmarshaler interface for Sub1. It set the default values for the Sub1 type
func (m *Sub1) UnmarshalJSON(data []byte) error {
	// Set default values
	*m = *NewSub1()

	// Unmarshal using an alias to avoid an infinite loop
	type alias Sub1
	err := json.Unmarshal(data, (*alias)(m))
	if err != nil {
		return err
	}
	return nil
}

// Validate implements basic validation for this model
func (m Sub1) Validate() error {
	errors := validation.Errors{}
	return errors.Filter()
}

// GetFoo returns the Foo property
func (m Sub1) GetFoo() string {
	return m.Foo
}

// SetFoo sets the Foo property
func (m *Sub1) SetFoo(val string) {
	m.Foo = val
}
