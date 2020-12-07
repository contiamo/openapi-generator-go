// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// ConnectionType is an enum.
type ConnectionType string

var (
	ConnectionTypeOther   ConnectionType = "other"
	ConnectionTypePrimary ConnectionType = "primary"

	// KnownConnectionType is the list of valid ConnectionType
	KnownConnectionType = []ConnectionType{
		ConnectionTypeOther,
		ConnectionTypePrimary,
	}
	// KnownConnectionTypeString is the list of valid ConnectionType as string
	KnownConnectionTypeString = []string{
		string(ConnectionTypeOther),
		string(ConnectionTypePrimary),
	}

	// InKnownConnectionType is an ozzo-validator for ConnectionType
	InKnownConnectionType = validation.In(
		ConnectionTypeOther,
		ConnectionTypePrimary,
	)
)
