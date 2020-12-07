// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// PipelineType is an enum.
type PipelineType string

var (
	PipelineTypeGeneric PipelineType = "generic"
	PipelineTypeUnknown PipelineType = "unknown"

	// KnownPipelineType is the list of valid PipelineType
	KnownPipelineType = []PipelineType{
		PipelineTypeGeneric,
		PipelineTypeUnknown,
	}
	// KnownPipelineTypeString is the list of valid PipelineType as string
	KnownPipelineTypeString = []string{
		string(PipelineTypeGeneric),
		string(PipelineTypeUnknown),
	}

	// InKnownPipelineType is an ozzo-validator for PipelineType
	InKnownPipelineType = validation.In(
		PipelineTypeGeneric,
		PipelineTypeUnknown,
	)
)
