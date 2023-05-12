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
type FavoriteNumber int32

// Validate implements basic validation for this model
func (m FavoriteNumber) Validate() error {
	return InKnownFavoriteNumber.Validate(m)
}

var (
	FavoriteNumber1 FavoriteNumber = 1
	FavoriteNumber3 FavoriteNumber = 3
	FavoriteNumber5 FavoriteNumber = 5
	FavoriteNumber8 FavoriteNumber = 8

	// KnownFavoriteNumber is the list of valid FavoriteNumber
	KnownFavoriteNumber = []FavoriteNumber{
		FavoriteNumber1,
		FavoriteNumber3,
		FavoriteNumber5,
		FavoriteNumber8,
	}
	// KnownFavoriteNumberInt32 is the list of valid FavoriteNumber as int32
	KnownFavoriteNumberInt32 = []int32{
		int32(FavoriteNumber1),
		int32(FavoriteNumber3),
		int32(FavoriteNumber5),
		int32(FavoriteNumber8),
	}

	// InKnownFavoriteNumber is an ozzo-validator for FavoriteNumber
	InKnownFavoriteNumber = validation.In(
		FavoriteNumber1,
		FavoriteNumber3,
		FavoriteNumber5,
		FavoriteNumber8,
	)
)