// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// ModelType is an enum.
type ModelType string

var (
	ModelTypeGeneric ModelType = "generic"
	ModelTypeUnknown ModelType = "unknown"

	// KnownModelType is the list of valid ModelType
	KnownModelType = []ModelType{
		ModelTypeGeneric,
		ModelTypeUnknown,
	}
	// KnownModelTypeString is the list of valid ModelType as string
	KnownModelTypeString = []string{
		string(ModelTypeGeneric),
		string(ModelTypeUnknown),
	}

	// InKnownModelType is an ozzo-validator for ModelType
	InKnownModelType = validation.In(
		ModelTypeGeneric,
		ModelTypeUnknown,
	)
)
