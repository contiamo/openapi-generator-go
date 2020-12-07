// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// IngestInsertModeType is an enum. Defines how to apply rows to a table.
// upsert: insert or update rows from the upload to the table
// append: insert rows from the upload to the table.
type IngestInsertModeType string

var (
	IngestInsertModeTypeAppend IngestInsertModeType = "append"
	IngestInsertModeTypeUpsert IngestInsertModeType = "upsert"

	// KnownIngestInsertModeType is the list of valid IngestInsertModeType
	KnownIngestInsertModeType = []IngestInsertModeType{
		IngestInsertModeTypeAppend,
		IngestInsertModeTypeUpsert,
	}
	// KnownIngestInsertModeTypeString is the list of valid IngestInsertModeType as string
	KnownIngestInsertModeTypeString = []string{
		string(IngestInsertModeTypeAppend),
		string(IngestInsertModeTypeUpsert),
	}

	// InKnownIngestInsertModeType is an ozzo-validator for IngestInsertModeType
	InKnownIngestInsertModeType = validation.In(
		IngestInsertModeTypeAppend,
		IngestInsertModeTypeUpsert,
	)
)
