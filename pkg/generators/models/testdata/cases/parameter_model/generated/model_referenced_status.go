// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: Test
//	Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// ReferencedStatus is an enum. A referenced status
type ReferencedStatus string

// Validate implements basic validation for this model
func (m ReferencedStatus) Validate() error {
	return InKnownReferencedStatus.Validate(m)
}

var (
	ReferencedStatusClosed    ReferencedStatus = "closed"
	ReferencedStatusDismissed ReferencedStatus = "dismissed"
	ReferencedStatusFixed     ReferencedStatus = "fixed"
	ReferencedStatusOpen      ReferencedStatus = "open"

	// KnownReferencedStatus is the list of valid ReferencedStatus
	KnownReferencedStatus = []ReferencedStatus{
		ReferencedStatusClosed,
		ReferencedStatusDismissed,
		ReferencedStatusFixed,
		ReferencedStatusOpen,
	}
	// KnownReferencedStatusString is the list of valid ReferencedStatus as string
	KnownReferencedStatusString = []string{
		string(ReferencedStatusClosed),
		string(ReferencedStatusDismissed),
		string(ReferencedStatusFixed),
		string(ReferencedStatusOpen),
	}

	// InKnownReferencedStatus is an ozzo-validator for ReferencedStatus
	InKnownReferencedStatus = validation.In(
		ReferencedStatusClosed,
		ReferencedStatusDismissed,
		ReferencedStatusFixed,
		ReferencedStatusOpen,
	)
)