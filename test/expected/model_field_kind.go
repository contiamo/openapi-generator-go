// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// FieldKind is an enum. The source type of the field.
//
// - System fields are non-editable default fields.  These can not be modified for deleted.
//
// - Configurable fields are editable default fields. These can be modified but not deleted.
//
// - Custom are user defined fields. These are fully defined by the user.
type FieldKind string

var (
	FieldKindConfigurable FieldKind = "configurable"
	FieldKindCustom       FieldKind = "custom"
	FieldKindSystem       FieldKind = "system"

	// KnownFieldKind is the list of valid FieldKind
	KnownFieldKind = []FieldKind{
		FieldKindConfigurable,
		FieldKindCustom,
		FieldKindSystem,
	}
	// KnownFieldKindString is the list of valid FieldKind as string
	KnownFieldKindString = []string{
		string(FieldKindConfigurable),
		string(FieldKindCustom),
		string(FieldKindSystem),
	}

	// InKnownFieldKind is an ozzo-validator for FieldKind
	InKnownFieldKind = validation.In(
		FieldKindConfigurable,
		FieldKindCustom,
		FieldKindSystem,
	)
)
