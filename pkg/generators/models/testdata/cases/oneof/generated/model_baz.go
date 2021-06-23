// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test
//     Version: 0.1.0
package generatortest

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
)

// Baz is a oneOf type.
type Baz struct {
	data interface{}
}

// MarshalJSON implementes the json.Marshaller interface
func (m *Baz) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.data)
}

// UnmarshalJSON implementes the json.Unmarshaller interface
func (m *Baz) UnmarshalJSON(bs []byte) error {
	return json.Unmarshal(bs, &m.data)
}

// BazAsFoo converts Baz to a Foo
func (m Baz) AsFoo() (result Foo, err error) {
	return result, mapstructure.Decode(m.data, &result)
}

// BazAsBar converts Baz to a Bar
func (m Baz) AsBar() (result Bar, err error) {
	return result, mapstructure.Decode(m.data, &result)
}

// BazAsPerson converts Baz to a Person
func (m Baz) AsPerson() (result Person, err error) {
	return result, mapstructure.Decode(m.data, &result)
}
