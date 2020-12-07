// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// StreamType is an enum.
type StreamType string

var (
	StreamTypeGeneric StreamType = "generic"
	StreamTypeUnknown StreamType = "unknown"

	// KnownStreamType is the list of valid StreamType
	KnownStreamType = []StreamType{
		StreamTypeGeneric,
		StreamTypeUnknown,
	}
	// KnownStreamTypeString is the list of valid StreamType as string
	KnownStreamTypeString = []string{
		string(StreamTypeGeneric),
		string(StreamTypeUnknown),
	}

	// InKnownStreamType is an ozzo-validator for StreamType
	InKnownStreamType = validation.In(
		StreamTypeGeneric,
		StreamTypeUnknown,
	)
)
