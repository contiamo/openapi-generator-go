// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: Test
//	Version: 0.1.0
package generatortest

import (
	"encoding/json"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Foo is an object.
type Foo struct {
	FooProperties
	AdditionalProperties map[string]interface{}
}

type FooProperties struct {
	// Bar:
	Bar int32 `json:"bar,omitempty"`
}

// Unmarshal all named properties into FooProperties and
// the rest into the AdditionalProperties map
func (obj *Foo) UnmarshalJSON(data []byte) error {
	var generic map[string]json.RawMessage
	if err := json.Unmarshal(data, &generic); err != nil {
		return err
	}

	obj.FooProperties = FooProperties{}

	var additionalProperties = make(map[string]interface{})
	for k, v := range generic {
		if k == "bar" {
			if err := json.Unmarshal(v, &(obj.FooProperties.Bar)); err != nil {
				return err
			}
			continue
		}
		var prop interface{}
		if err := json.Unmarshal(v, &prop); err != nil {
			return err
		}
		additionalProperties[k] = prop
	}

	obj.AdditionalProperties = additionalProperties
	return nil
}

// Marshal Foo by combining the AdditionalProperties with the
// named properties in a single JSON object. Named properties take
// precedence.
func (obj Foo) MarshalJSON() ([]byte, error) {
	props := make(map[string]json.RawMessage)

	// start with additional properties so regular properties overwrite them
	for k, v := range obj.AdditionalProperties {
		if propData, err := json.Marshal(v); err == nil {
			props[k] = propData
		} else {
			return nil, err
		}
	}
	if propData, err := json.Marshal(obj.FooProperties.Bar); err == nil {
		props["bar"] = propData
	} else {
		return nil, err
	}

	data, err := json.Marshal(props)
	return data, err
}

// Validate implements basic validation for this model
func (m Foo) Validate() error {
	return validation.Errors{}.Filter()
}

// GetBar returns the Bar property
func (m Foo) GetBar() int32 {
	return m.Bar
}

// SetBar sets the Bar property
func (m *Foo) SetBar(val int32) {
	m.Bar = val
}

func (m *Foo) GetAdditionalProperties() map[string]interface{} {
	return m.AdditionalProperties
}

func (m *Foo) SetAdditionalProperties(val map[string]interface{}) {
	m.AdditionalProperties = val
}
