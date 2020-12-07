// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// OperationType is an enum.
type OperationType string

var (
	OperationTypeEq      OperationType = "eq"
	OperationTypeGt      OperationType = "gt"
	OperationTypeGte     OperationType = "gte"
	OperationTypeIn      OperationType = "in"
	OperationTypeLt      OperationType = "lt"
	OperationTypeLte     OperationType = "lte"
	OperationTypeOverlap OperationType = "overlap"
	OperationTypeSub     OperationType = "sub"
	OperationTypeSup     OperationType = "sup"

	// KnownOperationType is the list of valid OperationType
	KnownOperationType = []OperationType{
		OperationTypeEq,
		OperationTypeGt,
		OperationTypeGte,
		OperationTypeIn,
		OperationTypeLt,
		OperationTypeLte,
		OperationTypeOverlap,
		OperationTypeSub,
		OperationTypeSup,
	}
	// KnownOperationTypeString is the list of valid OperationType as string
	KnownOperationTypeString = []string{
		string(OperationTypeEq),
		string(OperationTypeGt),
		string(OperationTypeGte),
		string(OperationTypeIn),
		string(OperationTypeLt),
		string(OperationTypeLte),
		string(OperationTypeOverlap),
		string(OperationTypeSub),
		string(OperationTypeSup),
	}

	// InKnownOperationType is an ozzo-validator for OperationType
	InKnownOperationType = validation.In(
		OperationTypeEq,
		OperationTypeGt,
		OperationTypeGte,
		OperationTypeIn,
		OperationTypeLt,
		OperationTypeLte,
		OperationTypeOverlap,
		OperationTypeSub,
		OperationTypeSup,
	)
)
