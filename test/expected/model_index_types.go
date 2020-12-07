// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// IndexTypes is an enum. Supported index types
type IndexTypes string

var (
	IndexTypesBtree IndexTypes = "btree"
	IndexTypesHash  IndexTypes = "hash"

	// KnownIndexTypes is the list of valid IndexTypes
	KnownIndexTypes = []IndexTypes{
		IndexTypesBtree,
		IndexTypesHash,
	}
	// KnownIndexTypesString is the list of valid IndexTypes as string
	KnownIndexTypesString = []string{
		string(IndexTypesBtree),
		string(IndexTypesHash),
	}

	// InKnownIndexTypes is an ozzo-validator for IndexTypes
	InKnownIndexTypes = validation.In(
		IndexTypesBtree,
		IndexTypesHash,
	)
)
