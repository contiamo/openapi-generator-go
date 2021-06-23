// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test
//     Version: 0.1.0
package generatortest

import "github.com/mitchellh/mapstructure"

// Baz is a value type.
type Baz interface{}

// BazAsFoo converts Baz to a Foo
func BazAsFoo(m Baz) (result Foo, err error) {
	return result, mapstructure.Decode(m, &result)
}

// BazAsBar converts Baz to a Bar
func BazAsBar(m Baz) (result Bar, err error) {
	return result, mapstructure.Decode(m, &result)
}
