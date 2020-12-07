// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// DataSourceType is an enum. Type of a data source
type DataSourceType string

var (
	DataSourceTypeExternal DataSourceType = "external"
	DataSourceTypeManaged  DataSourceType = "managed"
	DataSourceTypeVirtual  DataSourceType = "virtual"

	// KnownDataSourceType is the list of valid DataSourceType
	KnownDataSourceType = []DataSourceType{
		DataSourceTypeExternal,
		DataSourceTypeManaged,
		DataSourceTypeVirtual,
	}
	// KnownDataSourceTypeString is the list of valid DataSourceType as string
	KnownDataSourceTypeString = []string{
		string(DataSourceTypeExternal),
		string(DataSourceTypeManaged),
		string(DataSourceTypeVirtual),
	}

	// InKnownDataSourceType is an ozzo-validator for DataSourceType
	InKnownDataSourceType = validation.In(
		DataSourceTypeExternal,
		DataSourceTypeManaged,
		DataSourceTypeVirtual,
	)
)
