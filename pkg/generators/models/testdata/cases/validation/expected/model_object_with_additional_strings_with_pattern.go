// Code generated by openapi-generator-go DO NOT EDIT.
//
// Source:
//
//	Title: Test
//	Version: 0.1.0
package generatortest

import (
	"encoding/json"
	"fmt"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"regexp"
)

// objectWithAdditionalStringsWithPatternAdditionalPropertiesPattern is the validation pattern for ObjectWithAdditionalStringsWithPattern.AdditionalProperties
var objectWithAdditionalStringsWithPatternAdditionalPropertiesPattern = regexp.MustCompile(`^[a-z]+$`)

// ObjectWithAdditionalStringsWithPattern is an object with additional properties.
type ObjectWithAdditionalStringsWithPattern struct {
	// Field1:
	Field1 string `json:"field1,omitempty" mapstructure:"field1,omitempty"`
	// AdditionalProperties:
	AdditionalProperties map[string]string `json:"-"`
}

// NewObjectWithAdditionalStringsWithPattern instantiates a new ObjectWithAdditionalStringsWithPattern with default values overriding them as follows:
// 1. Default values specified in the ObjectWithAdditionalStringsWithPattern schema
// 2. Default values specified per ObjectWithAdditionalStringsWithPattern property
func NewObjectWithAdditionalStringsWithPattern() *ObjectWithAdditionalStringsWithPattern {
	m := &ObjectWithAdditionalStringsWithPattern{}

	return m
}

// UnmarshalJSON all named properties into ObjectWithAdditionalStringsWithPattern and the rest into the AdditionalProperties map
func (m *ObjectWithAdditionalStringsWithPattern) UnmarshalJSON(data []byte) error {
	// Set default values
	*m = *NewObjectWithAdditionalStringsWithPattern()

	var generic map[string]json.RawMessage
	if err := json.Unmarshal(data, &generic); err != nil {
		return err
	}

	var additionalProperties = make(map[string]string)
	for k, v := range generic {
		switch k {
		case "field1":
			if err := json.Unmarshal(v, &(m.Field1)); err != nil {
				return err
			}
		default:
			var prop string
			if err := json.Unmarshal(v, &prop); err != nil {
				return err
			}
			additionalProperties[k] = prop
		}
	}

	if len(additionalProperties) > 0 {
		m.AdditionalProperties = additionalProperties
	}
	return nil
}

const ObjectWithAdditionalStringsWithPatternNamedProperties = "|field1"

// MarshalJSON ObjectWithAdditionalStringsWithPattern by combining the AdditionalProperties with the named properties in a single JSON object.
// An error will be returned if there are duplicate keys.
func (m ObjectWithAdditionalStringsWithPattern) MarshalJSON() ([]byte, error) {
	type alias ObjectWithAdditionalStringsWithPattern
	data, err := json.Marshal(alias(m))
	if err != nil {
		return nil, err
	}

	if len(m.AdditionalProperties) > 0 {
		for k, _ := range m.AdditionalProperties {
			if strings.Contains(ObjectWithAdditionalStringsWithPatternNamedProperties, k) {
				return nil, fmt.Errorf("named key: %s was found in additionalProperties field", k)
			}
		}
		additionalData, err := json.Marshal(m.AdditionalProperties)
		if err != nil {
			return nil, err
		}

		// merge the two JSON objects, we do it at byte level to avoid re-encoding the JSON
		data = data[:len(data)-1]
		additionalData = additionalData[1:]
		if len(data) > 1 {
			data = append(data, ',')
		}
		data = append(data, additionalData...)
	}

	return data, err
}

// Validate implements basic validation for this model
func (m ObjectWithAdditionalStringsWithPattern) Validate() error {
	errors := validation.Errors{
		"field1": validation.Validate(
			m.Field1, validation.Length(2, 0),
		),
	}
	for k, v := range m.AdditionalProperties {
		if err := validation.Validate(v, validation.Required, validation.Length(2, 0), validation.Match(objectWithAdditionalStringsWithPatternAdditionalPropertiesPattern)); err != nil {
			errors[k] = err
		}
	}
	return errors.Filter()
}

// GetField1 returns the Field1 property
func (m ObjectWithAdditionalStringsWithPattern) GetField1() string {
	return m.Field1
}

// SetField1 sets the Field1 property
func (m *ObjectWithAdditionalStringsWithPattern) SetField1(val string) {
	m.Field1 = val
}

func (m *ObjectWithAdditionalStringsWithPattern) GetAdditionalProperties() map[string]string {
	return m.AdditionalProperties
}

func (m *ObjectWithAdditionalStringsWithPattern) SetAdditionalProperties(val map[string]string) {
	m.AdditionalProperties = val
}
