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

// Bar is a oneOf type.
type Bar struct {
	data interface{}
}

// MarshalJSON implements the json.Marshaller interface
func (m Bar) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.data)
}

// UnmarshalJSON implements the json.Unmarshaller interface
func (m *Bar) UnmarshalJSON(bs []byte) error {
	return json.Unmarshal(bs, &m.data)
}

// FromString sets the Bar data.
func (m *Bar) FromString(data string) {
	m.data = data
}

// FromInt32 sets the Bar data.
func (m *Bar) FromInt32(data int32) {
	m.data = data
}

// As converts Bar to a user defined structure.
func (m Bar) As(target interface{}) error {
	return mapstructure.Decode(m.data, target)
}

// AsString converts Bar to a string
func (m Bar) AsString() (result string, err error) {
	return result, mapstructure.Decode(m.data, &result)
}

// AsInt32 converts Bar to a int32
func (m Bar) AsInt32() (result int32, err error) {
	return result, mapstructure.Decode(m.data, &result)
}
