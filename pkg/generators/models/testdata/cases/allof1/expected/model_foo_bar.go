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

// FooBar is an object.
type FooBar struct {
	// Bar:
	Bar string `json:"bar,omitempty" mapstructure:"bar,omitempty"`
	// Foo:
	Foo string `json:"foo,omitempty" mapstructure:"foo,omitempty"`
}

// NewFooBar instantiates a new FooBar with default values overriding them as follows:
// 1. Default values specified in the FooBar schema
// 2. Default values specified per FooBar property
func NewFooBar() *FooBar {
	m := &FooBar{}

	return m
}

// UnmarshalJSON implements the json.Unmarshaler interface for FooBar. It set the default values for the FooBar type
func (m *FooBar) UnmarshalJSON(data []byte) error {
	// Set default values
	*m = *NewFooBar()

	// Unmarshal using an alias to avoid an infinite loop
	type alias FooBar
	err := json.Unmarshal(data, (*alias)(m))
	if err != nil {
		return err
	}
	return nil
}

// Validate implements basic validation for this model
func (m FooBar) Validate() error {
	errors := validation.Errors{}
	return errors.Filter()
}

// GetBar returns the Bar property
func (m FooBar) GetBar() string {
	return m.Bar
}

// SetBar sets the Bar property
func (m *FooBar) SetBar(val string) {
	m.Bar = val
}

// GetFoo returns the Foo property
func (m FooBar) GetFoo() string {
	return m.Foo
}

// SetFoo sets the Foo property
func (m *FooBar) SetFoo(val string) {
	m.Foo = val
}
