// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// SchemaSyncStatus is an enum. Describes the sync status for the data source schema. `partial` means the synchronization process failed but some of the tables were fetched successfully.
type SchemaSyncStatus string

var (
	SchemaSyncStatusFail    SchemaSyncStatus = "fail"
	SchemaSyncStatusPartial SchemaSyncStatus = "partial"
	SchemaSyncStatusRunning SchemaSyncStatus = "running"
	SchemaSyncStatusSuccess SchemaSyncStatus = "success"

	// KnownSchemaSyncStatus is the list of valid SchemaSyncStatus
	KnownSchemaSyncStatus = []SchemaSyncStatus{
		SchemaSyncStatusFail,
		SchemaSyncStatusPartial,
		SchemaSyncStatusRunning,
		SchemaSyncStatusSuccess,
	}
	// KnownSchemaSyncStatusString is the list of valid SchemaSyncStatus as string
	KnownSchemaSyncStatusString = []string{
		string(SchemaSyncStatusFail),
		string(SchemaSyncStatusPartial),
		string(SchemaSyncStatusRunning),
		string(SchemaSyncStatusSuccess),
	}

	// InKnownSchemaSyncStatus is an ozzo-validator for SchemaSyncStatus
	InKnownSchemaSyncStatus = validation.In(
		SchemaSyncStatusFail,
		SchemaSyncStatusPartial,
		SchemaSyncStatusRunning,
		SchemaSyncStatusSuccess,
	)
)
