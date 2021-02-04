// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Baz is an object.
type Baz map[string][]Bar

// Validate implements basic validation for this model
func (m Baz) Validate() error {
	return validation.Errors{}.Filter()
}
