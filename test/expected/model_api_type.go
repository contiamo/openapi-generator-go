// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// APIType is an enum.
type APIType string

var (
	APITypeGeneric APIType = "generic"
	APITypeGraphql APIType = "graphql"
	APITypeRest    APIType = "rest"
	APITypeRpc     APIType = "rpc"

	// KnownAPIType is the list of valid APIType
	KnownAPIType = []APIType{
		APITypeGeneric,
		APITypeGraphql,
		APITypeRest,
		APITypeRpc,
	}
	// KnownAPITypeString is the list of valid APIType as string
	KnownAPITypeString = []string{
		string(APITypeGeneric),
		string(APITypeGraphql),
		string(APITypeRest),
		string(APITypeRpc),
	}

	// InKnownAPIType is an ozzo-validator for APIType
	InKnownAPIType = validation.In(
		APITypeGeneric,
		APITypeGraphql,
		APITypeRest,
		APITypeRpc,
	)
)
