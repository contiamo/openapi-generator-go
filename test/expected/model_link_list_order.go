// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// LinkListOrder is an enum.
type LinkListOrder string

var (
	LinkListOrderSourceId LinkListOrder = "sourceId"
	LinkListOrderTargetId LinkListOrder = "targetId"

	// KnownLinkListOrder is the list of valid LinkListOrder
	KnownLinkListOrder = []LinkListOrder{
		LinkListOrderSourceId,
		LinkListOrderTargetId,
	}
	// KnownLinkListOrderString is the list of valid LinkListOrder as string
	KnownLinkListOrderString = []string{
		string(LinkListOrderSourceId),
		string(LinkListOrderTargetId),
	}

	// InKnownLinkListOrder is an ozzo-validator for LinkListOrder
	InKnownLinkListOrder = validation.In(
		LinkListOrderSourceId,
		LinkListOrderTargetId,
	)
)
