// Code generated by openapi-generator-go DO NOT EDIT.
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

// NewEitherIdOrNamespaceName creates a new EitherIdOrNamespaceName instance with no internal value.
func NewEitherIdOrNamespaceName() *EitherIdOrNamespaceName {
	return &EitherIdOrNamespaceName{}
}

// MarshalJSON implements the json.Marshaller interface
func (m EitherIdOrNamespaceName) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.data)
}

// UnmarshalJSON implements the json.Unmarshaller interface
func (m *EitherIdOrNamespaceName) UnmarshalJSON(bs []byte) error {
	return json.Unmarshal(bs, &m.data)
}

// FromEitherIdOrNamespaceNameOneOfIdx0 sets the EitherIdOrNamespaceName from a EitherIdOrNamespaceNameOneOfIdx0
func (m *EitherIdOrNamespaceName) FromEitherIdOrNamespaceNameOneOfIdx0(data EitherIdOrNamespaceNameOneOfIdx0) {
	m.data = data
}

// FromEitherIdOrNamespaceNameOneOfIdx1 sets the EitherIdOrNamespaceName from a EitherIdOrNamespaceNameOneOfIdx1
func (m *EitherIdOrNamespaceName) FromEitherIdOrNamespaceNameOneOfIdx1(data EitherIdOrNamespaceNameOneOfIdx1) {
	m.data = data
}

// As converts EitherIdOrNamespaceName to a user defined structure.
func (m EitherIdOrNamespaceName) As(target interface{}) error {
	return mapstructure.Decode(m.data, target)
}

// AsEitherIdOrNamespaceNameOneOfIdx0 converts EitherIdOrNamespaceName to a EitherIdOrNamespaceNameOneOfIdx0
func (m EitherIdOrNamespaceName) AsEitherIdOrNamespaceNameOneOfIdx0() (result EitherIdOrNamespaceNameOneOfIdx0, err error) {
	result = *NewEitherIdOrNamespaceNameOneOfIdx0()

	return result, mapstructure.Decode(m.data, &result)
}

// AsEitherIdOrNamespaceNameOneOfIdx1 converts EitherIdOrNamespaceName to a EitherIdOrNamespaceNameOneOfIdx1
func (m EitherIdOrNamespaceName) AsEitherIdOrNamespaceNameOneOfIdx1() (result EitherIdOrNamespaceNameOneOfIdx1, err error) {
	result = *NewEitherIdOrNamespaceNameOneOfIdx1()

	return result, mapstructure.Decode(m.data, &result)
}
