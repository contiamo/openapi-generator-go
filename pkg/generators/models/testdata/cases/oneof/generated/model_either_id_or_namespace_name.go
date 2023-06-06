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

// EitherIdOrNamespaceName is a oneOf type. Either an object with id or a namespace with a name
type EitherIdOrNamespaceName struct {
	data interface{}
}

// MarshalJSON implements the json.Marshaller interface
func (m EitherIdOrNamespaceName) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.data)
}

// UnmarshalJSON implements the json.Unmarshaller interface
func (m *EitherIdOrNamespaceName) UnmarshalJSON(bs []byte) error {
	return json.Unmarshal(bs, &m.data)
}

// FromStructId sets the EitherIdOrNamespaceName from a struct {
// Id string `json:"id,omitempty" mapstructure:"id,omitempty"`
// }
func (m *EitherIdOrNamespaceName) FromStructId(data struct {
	Id string `json:"id,omitempty" mapstructure:"id,omitempty"`
}) {
	m.data = data
}

// FromStructNameNamespace sets the EitherIdOrNamespaceName from a struct {
// Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
// Namespace string `json:"namespace,omitempty" mapstructure:"namespace,omitempty"`
// }
func (m *EitherIdOrNamespaceName) FromStructNameNamespace(data struct {
	Name      string `json:"name,omitempty" mapstructure:"name,omitempty"`
	Namespace string `json:"namespace,omitempty" mapstructure:"namespace,omitempty"`
}) {
	m.data = data
}

// As converts EitherIdOrNamespaceName to a user defined structure.
func (m EitherIdOrNamespaceName) As(target interface{}) error {
	return mapstructure.Decode(m.data, target)
}

// AsStructId converts EitherIdOrNamespaceName to a struct {
// Id string `json:"id,omitempty" mapstructure:"id,omitempty"`
// }
func (m EitherIdOrNamespaceName) AsStructId() (result struct {
	Id string `json:"id,omitempty" mapstructure:"id,omitempty"`
}, err error) {
	return result, mapstructure.Decode(m.data, &result)
}

// AsStructNameNamespace converts EitherIdOrNamespaceName to a struct {
// Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
// Namespace string `json:"namespace,omitempty" mapstructure:"namespace,omitempty"`
// }
func (m EitherIdOrNamespaceName) AsStructNameNamespace() (result struct {
	Name      string `json:"name,omitempty" mapstructure:"name,omitempty"`
	Namespace string `json:"namespace,omitempty" mapstructure:"namespace,omitempty"`
}, err error) {
	return result, mapstructure.Decode(m.data, &result)
}
