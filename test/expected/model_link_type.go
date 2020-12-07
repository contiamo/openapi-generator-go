// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// LinkType is an enum.
type LinkType string

var (
	LinkTypeDependency  LinkType = "dependency"
	LinkTypeGlossary    LinkType = "glossary"
	LinkTypeReference   LinkType = "reference"
	LinkTypeReplacement LinkType = "replacement"

	// KnownLinkType is the list of valid LinkType
	KnownLinkType = []LinkType{
		LinkTypeDependency,
		LinkTypeGlossary,
		LinkTypeReference,
		LinkTypeReplacement,
	}
	// KnownLinkTypeString is the list of valid LinkType as string
	KnownLinkTypeString = []string{
		string(LinkTypeDependency),
		string(LinkTypeGlossary),
		string(LinkTypeReference),
		string(LinkTypeReplacement),
	}

	// InKnownLinkType is an ozzo-validator for LinkType
	InKnownLinkType = validation.In(
		LinkTypeDependency,
		LinkTypeGlossary,
		LinkTypeReference,
		LinkTypeReplacement,
	)
)
