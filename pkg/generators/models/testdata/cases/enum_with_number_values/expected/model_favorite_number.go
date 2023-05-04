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

// FavoriteNumber is an enum.
type FavoriteNumber float32

// Validate implements basic validation for this model
func (m FavoriteNumber) Validate() error {
	return InKnownFavoriteNumber.Validate(m)
}

var (
	FavoriteNumber1381  FavoriteNumber = 1.381
	FavoriteNumber2718  FavoriteNumber = 2.718
	FavoriteNumber3141  FavoriteNumber = 3.141
	FavoriteNumber60228 FavoriteNumber = 6.0228

	// KnownFavoriteNumber is the list of valid FavoriteNumber
	KnownFavoriteNumber = []FavoriteNumber{
		FavoriteNumber1381,
		FavoriteNumber2718,
		FavoriteNumber3141,
		FavoriteNumber60228,
	}
	// KnownFavoriteNumberFloat32 is the list of valid FavoriteNumber as float32
	KnownFavoriteNumberFloat32 = []float32{
		float32(FavoriteNumber1381),
		float32(FavoriteNumber2718),
		float32(FavoriteNumber3141),
		float32(FavoriteNumber60228),
	}

	// InKnownFavoriteNumber is an ozzo-validator for FavoriteNumber
	InKnownFavoriteNumber = validation.In(
		FavoriteNumber1381,
		FavoriteNumber2718,
		FavoriteNumber3141,
		FavoriteNumber60228,
	)
)
