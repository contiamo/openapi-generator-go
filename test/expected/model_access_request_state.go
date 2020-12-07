// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// AccessRequestState is an enum. state of the request, either 'pending', 'granted' or 'rejected'
type AccessRequestState string

var (
	AccessRequestStateGranted  AccessRequestState = "granted"
	AccessRequestStatePending  AccessRequestState = "pending"
	AccessRequestStateRejected AccessRequestState = "rejected"

	// KnownAccessRequestState is the list of valid AccessRequestState
	KnownAccessRequestState = []AccessRequestState{
		AccessRequestStateGranted,
		AccessRequestStatePending,
		AccessRequestStateRejected,
	}
	// KnownAccessRequestStateString is the list of valid AccessRequestState as string
	KnownAccessRequestStateString = []string{
		string(AccessRequestStateGranted),
		string(AccessRequestStatePending),
		string(AccessRequestStateRejected),
	}

	// InKnownAccessRequestState is an ozzo-validator for AccessRequestState
	InKnownAccessRequestState = validation.In(
		AccessRequestStateGranted,
		AccessRequestStatePending,
		AccessRequestStateRejected,
	)
)
