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

// DataSourceTechnology is an enum. Technology of a data source. For external data sources this also specifies the type of connection (driver) being used to connect to the data source. Generic data sources may use any technology value (but no connection will be established). Generic data sources may also use the value "other" for technology in case no technology should be shown in the catalog or the technology is not known to the catalog.
type DataSourceTechnology string

// Validate implements basic validation for this model
func (m DataSourceTechnology) Validate() error {
	return InKnownDataSourceTechnology.Validate(m)
}

var (
	DataSourceTechnologyMssql      DataSourceTechnology = "mssql"
	DataSourceTechnologyMysql      DataSourceTechnology = "mysql"
	DataSourceTechnologyOracle     DataSourceTechnology = "oracle"
	DataSourceTechnologyOther      DataSourceTechnology = "other"
	DataSourceTechnologyPostgresql DataSourceTechnology = "postgresql"

	// KnownDataSourceTechnology is the list of valid DataSourceTechnology
	KnownDataSourceTechnology = []DataSourceTechnology{
		DataSourceTechnologyMssql,
		DataSourceTechnologyMysql,
		DataSourceTechnologyOracle,
		DataSourceTechnologyOther,
		DataSourceTechnologyPostgresql,
	}
	// KnownDataSourceTechnologyString is the list of valid DataSourceTechnology as string
	KnownDataSourceTechnologyString = []string{
		string(DataSourceTechnologyMssql),
		string(DataSourceTechnologyMysql),
		string(DataSourceTechnologyOracle),
		string(DataSourceTechnologyOther),
		string(DataSourceTechnologyPostgresql),
	}

	// InKnownDataSourceTechnology is an ozzo-validator for DataSourceTechnology
	InKnownDataSourceTechnology = validation.In(
		DataSourceTechnologyMssql,
		DataSourceTechnologyMysql,
		DataSourceTechnologyOracle,
		DataSourceTechnologyOther,
		DataSourceTechnologyPostgresql,
	)
)
