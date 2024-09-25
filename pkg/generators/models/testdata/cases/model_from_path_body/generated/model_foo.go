// Code generated by openapi-generator-go DO NOT EDIT.
//
// Source:
//
//	Title: Test
//	Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Foo is an object.
type Foo map[string]interface{}

// NewFoo instantiates a new Foo with default values overriding them as follows:
// 1. Default values specified in the Foo schema
// 2. Default values specified per Foo property
func NewFoo() Foo {
	m := Foo{}

	return m
}

// Validate implements basic validation for this model
func (m Foo) Validate() error {
	errors := validation.Errors{}
	for k, v := range m {
		if err := validation.Validate(v); err != nil {
			errors[k] = err
		}
	}
	return errors.Filter()
}
