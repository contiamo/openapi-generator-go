// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test
//     Version: 0.1.0
package generatortest

import (
	"encoding/json"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/mitchellh/mapstructure"
)

type geometryer interface {
	GeometryDiscriminator() string
	Validate() error
}

// Geometry is a oneOf type.
type Geometry struct {
	data geometryer
}

// MarshalJSON implements the json.Marshaller interface
func (m Geometry) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.data)
}

// UnmarshalJSON implements the json.Unmarshaller interface
func (m *Geometry) UnmarshalJSON(bs []byte) error {
	discriminator := struct {
		Value string `json:"type"`
	}{}
	err := json.Unmarshal(bs, &discriminator)
	if err != nil {
		return err
	}

	if discriminator.Value == "" {
		return validation.Errors{
			"type": fmt.Errorf("can not be empty"),
		}.Filter()
	}

	switch discriminator.Value {
	case "Line":
		data := Line{}
		err := json.Unmarshal(bs, &data)
		if err != nil {
			return err
		}
		m.data = data
	case "Shape":
		data := Shape{}
		err := json.Unmarshal(bs, &data)
		if err != nil {
			return err
		}
		m.data = data
	default:
		return validation.Errors{
			"type": fmt.Errorf("invalid value"),
		}.Filter()
	}
	return nil
}

// FromLine sets the Geometry data.
func (m *Geometry) FromLine(data Line) {
	m.data = data
}

// FromShape sets the Geometry data.
func (m *Geometry) FromShape(data Shape) {
	m.data = data
}

// As converts Geometry to a user defined structure.
func (m Geometry) As(target interface{}) error {
	return mapstructure.Decode(m.data, target)
}

// AsLine converts Geometry to a Line
func (m Geometry) AsLine() (result Line, err error) {
	return result, mapstructure.Decode(m.data, &result)
}

// AsShape converts Geometry to a Shape
func (m Geometry) AsShape() (result Shape, err error) {
	return result, mapstructure.Decode(m.data, &result)
}

// Validate implements basic validation for this model
func (m Geometry) Validate() error {
	discriminator := m.data.GeometryDiscriminator()
	switch discriminator {
	case "Line":
		return m.Validate()
	case "Shape":
		return m.Validate()
	default:
		return validation.Errors{
			"type": fmt.Errorf("unknown type value"),
		}.Filter()
	}
}

// IsGeometry tests if data is one of the discriminated sub-types of Geometry.
func IsGeometry(data interface{}) bool {
	t, ok := data.(geometryer)
	if !ok {
		return false
	}

	discriminator := t.GeometryDiscriminator()
	switch discriminator {
	case "Line":
		return true
	case "Shape":
		return true
	default:
		return false
	}
}

// GeometryDiscriminator implements geometryer and returns the discriminator value as a string.
func (m Line) GeometryDiscriminator() string {
	return string(m.GetType())
}

// GeometryDiscriminator implements geometryer and returns the discriminator value as a string.
func (m Shape) GeometryDiscriminator() string {
	return string(m.GetType())
}
