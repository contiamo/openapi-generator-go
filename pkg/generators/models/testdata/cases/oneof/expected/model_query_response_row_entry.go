// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: Test
//	Version: 0.1.0
package generatortest

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
)

// QueryResponseRowEntry is a oneOf type. One field from query results
type QueryResponseRowEntry struct {
	data interface{}
}

// MarshalJSON implements the json.Marshaller interface
func (m QueryResponseRowEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.data)
}

// UnmarshalJSON implements the json.Unmarshaller interface
func (m *QueryResponseRowEntry) UnmarshalJSON(bs []byte) error {
	return json.Unmarshal(bs, &m.data)
}

// From*string sets the QueryResponseRowEntry data.
func (m *QueryResponseRowEntry) FromNullableString(data *string) {
	m.data = data
}

// FromBool sets the QueryResponseRowEntry data.
func (m *QueryResponseRowEntry) FromBool(data bool) {
	m.data = data
}

// FromFloat32 sets the QueryResponseRowEntry data.
func (m *QueryResponseRowEntry) FromFloat32(data float32) {
	m.data = data
}

// As converts QueryResponseRowEntry to a user defined structure.
func (m QueryResponseRowEntry) As(target interface{}) error {
	return mapstructure.Decode(m.data, target)
}

// As*string converts QueryResponseRowEntry to a *string
func (m QueryResponseRowEntry) AsNullableString() (result *string, err error) {
	return result, mapstructure.Decode(m.data, &result)
}

// AsBool converts QueryResponseRowEntry to a bool
func (m QueryResponseRowEntry) AsBool() (result bool, err error) {
	return result, mapstructure.Decode(m.data, &result)
}

// AsFloat32 converts QueryResponseRowEntry to a float32
func (m QueryResponseRowEntry) AsFloat32() (result float32, err error) {
	return result, mapstructure.Decode(m.data, &result)
}
