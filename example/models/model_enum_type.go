
// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test Service
//     Version: 0.1.0
package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// EnumType is an enum.
type EnumType string

var (
	EnumTypeA EnumType = "a"
	EnumTypeB EnumType = "b"
	EnumTypeC EnumType = "c"

	// KnownEnumType is the list of valid EnumType
	KnownEnumType = []EnumType{
		EnumTypeA,
		EnumTypeB,
		EnumTypeC,
	}
	// KnownEnumTypeString is the list of valid EnumType as string
	KnownEnumTypeString = []string{
		string(EnumTypeA),
		string(EnumTypeB),
		string(EnumTypeC),
	}

	// InKnownEnumType is an ozzo-validator for EnumType
	InKnownEnumType = validation.In(
		EnumTypeA,
		EnumTypeB,
		EnumTypeC,
	)
)
