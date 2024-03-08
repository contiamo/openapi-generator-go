// Code generated by openapi-generator-go DO NOT EDIT.
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

// Person is an object.
type Person struct {
	// Address:
	Address Address `json:"address" mapstructure:"address"`
	// Age:
	Age int32 `json:"age" mapstructure:"age"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Pets:
	Pets []Pet `json:"pets,omitempty" mapstructure:"pets,omitempty"`
}

// NewPerson instantiates a new Person with default values overriding them as follows:
// 1. Default values specified in the Person schema
// 2. Default values specified per Person property
func NewPerson() *Person {
	m := &Person{
		Address: *NewAddress(),
		Age:     13,
	}

	return m
}

// UnmarshalJSON implements the json.Unmarshaler interface for Person. It set the default values for the Person type
func (m *Person) UnmarshalJSON(data []byte) error {
	// Set default values
	*m = *NewPerson()

	// Unmarshal using an alias to avoid an infinite loop
	type alias Person
	err := json.Unmarshal(data, (*alias)(m))
	if err != nil {
		return err
	}
	return nil
}

// Validate implements basic validation for this model
func (m Person) Validate() error {
	errors := validation.Errors{
		"address": validation.Validate(
			m.Address,
		),
		"pets": validation.Validate(
			m.Pets,
		),
	}
	return errors.Filter()
}

// GetAddress returns the Address property
func (m Person) GetAddress() Address {
	return m.Address
}

// SetAddress sets the Address property
func (m *Person) SetAddress(val Address) {
	m.Address = val
}

// GetAge returns the Age property
func (m Person) GetAge() int32 {
	return m.Age
}

// SetAge sets the Age property
func (m *Person) SetAge(val int32) {
	m.Age = val
}

// GetName returns the Name property
func (m Person) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *Person) SetName(val string) {
	m.Name = val
}

// GetPets returns the Pets property
func (m Person) GetPets() []Pet {
	return m.Pets
}

// SetPets sets the Pets property
func (m *Person) SetPets(val []Pet) {
	m.Pets = val
}