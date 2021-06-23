// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test
//     Version: 0.1.0
package generatortest

import "github.com/mitchellh/mapstructure"

// Bar is a value type.
type Bar interface{}

// BarAsString converts Bar to a string
func BarAsString(m Bar) (result string, err error) {
	return result, mapstructure.Decode(m, &result)
}

// BarAsInt32 converts Bar to a int32
func BarAsInt32(m Bar) (result int32, err error) {
	return result, mapstructure.Decode(m, &result)
}
