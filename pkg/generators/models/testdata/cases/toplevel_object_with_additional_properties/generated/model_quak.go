// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Quak is an object.
type Quak map[string]interface{}

// Validate implements basic validation for this model
func (m Quak) Validate() error {
	return validation.Errors{}.Filter()
}
