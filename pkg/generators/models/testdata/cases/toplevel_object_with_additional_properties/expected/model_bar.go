// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Bar is an object.
type Bar map[string]Foo

// Validate implements basic validation for this model
func (m Bar) Validate() error {
	return validation.Errors{}.Filter()
}
