// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Food is an enum.
type Food string

// Validate implements basic validation for this model
func (m Food) Validate() error {
	return InKnownFood.Validate(m)
}

var (
	FoodApple  Food = "apple"
	FoodBanana Food = "banana"
	FoodPizza  Food = "pizza"

	// KnownFood is the list of valid Food
	KnownFood = []Food{
		FoodApple,
		FoodBanana,
		FoodPizza,
	}
	// KnownFoodString is the list of valid Food as string
	KnownFoodString = []string{
		string(FoodApple),
		string(FoodBanana),
		string(FoodPizza),
	}

	// InKnownFood is an ozzo-validator for Food
	InKnownFood = validation.In(
		FoodApple,
		FoodBanana,
		FoodPizza,
	)
)
