// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: Test
//	Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Color is an enum.
type Color string

// Validate implements basic validation for this model
func (m Color) Validate() error {
	return InKnownColor.Validate(m)
}

var (
	ColorBlue   Color = "blue"
	ColorGreen  Color = "green"
	ColorRed    Color = "red"
	ColorYellow Color = "yellow"

	// KnownColor is the list of valid Color
	KnownColor = []Color{
		ColorBlue,
		ColorGreen,
		ColorRed,
		ColorYellow,
	}
	// KnownColorString is the list of valid Color as string
	KnownColorString = []string{
		string(ColorBlue),
		string(ColorGreen),
		string(ColorRed),
		string(ColorYellow),
	}

	// InKnownColor is an ozzo-validator for Color
	InKnownColor = validation.In(
		ColorBlue,
		ColorGreen,
		ColorRed,
		ColorYellow,
	)
)
