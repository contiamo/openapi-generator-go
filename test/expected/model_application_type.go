// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// ApplicationType is an enum.
type ApplicationType string

var (
	ApplicationTypeGeneric ApplicationType = "generic"
	ApplicationTypeUnknown ApplicationType = "unknown"

	// KnownApplicationType is the list of valid ApplicationType
	KnownApplicationType = []ApplicationType{
		ApplicationTypeGeneric,
		ApplicationTypeUnknown,
	}
	// KnownApplicationTypeString is the list of valid ApplicationType as string
	KnownApplicationTypeString = []string{
		string(ApplicationTypeGeneric),
		string(ApplicationTypeUnknown),
	}

	// InKnownApplicationType is an ozzo-validator for ApplicationType
	InKnownApplicationType = validation.In(
		ApplicationTypeGeneric,
		ApplicationTypeUnknown,
	)
)
