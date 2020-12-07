// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// DataSourceTechnology is an enum. Type of a data source
type DataSourceTechnology string

var (
	DataSourceTechnologyBigquery       DataSourceTechnology = "bigquery"
	DataSourceTechnologyClickhouse     DataSourceTechnology = "clickhouse"
	DataSourceTechnologyDb2            DataSourceTechnology = "db2"
	DataSourceTechnologyHana           DataSourceTechnology = "hana"
	DataSourceTechnologyHive           DataSourceTechnology = "hive"
	DataSourceTechnologyHsqldbFoodmart DataSourceTechnology = "hsqldb-foodmart"
	DataSourceTechnologyMsdynamics     DataSourceTechnology = "msdynamics"
	DataSourceTechnologyMssql          DataSourceTechnology = "mssql"
	DataSourceTechnologyMssqlCds       DataSourceTechnology = "mssql-cds"
	DataSourceTechnologyMysql          DataSourceTechnology = "mysql"
	DataSourceTechnologyOracle         DataSourceTechnology = "oracle"
	DataSourceTechnologyPostgresql     DataSourceTechnology = "postgresql"
	DataSourceTechnologyRedshift       DataSourceTechnology = "redshift"
	DataSourceTechnologyS3             DataSourceTechnology = "s3"
	DataSourceTechnologySaperp         DataSourceTechnology = "saperp"
	DataSourceTechnologyTeradata       DataSourceTechnology = "teradata"

	// KnownDataSourceTechnology is the list of valid DataSourceTechnology
	KnownDataSourceTechnology = []DataSourceTechnology{
		DataSourceTechnologyBigquery,
		DataSourceTechnologyClickhouse,
		DataSourceTechnologyDb2,
		DataSourceTechnologyHana,
		DataSourceTechnologyHive,
		DataSourceTechnologyHsqldbFoodmart,
		DataSourceTechnologyMsdynamics,
		DataSourceTechnologyMssql,
		DataSourceTechnologyMssqlCds,
		DataSourceTechnologyMysql,
		DataSourceTechnologyOracle,
		DataSourceTechnologyPostgresql,
		DataSourceTechnologyRedshift,
		DataSourceTechnologyS3,
		DataSourceTechnologySaperp,
		DataSourceTechnologyTeradata,
	}
	// KnownDataSourceTechnologyString is the list of valid DataSourceTechnology as string
	KnownDataSourceTechnologyString = []string{
		string(DataSourceTechnologyBigquery),
		string(DataSourceTechnologyClickhouse),
		string(DataSourceTechnologyDb2),
		string(DataSourceTechnologyHana),
		string(DataSourceTechnologyHive),
		string(DataSourceTechnologyHsqldbFoodmart),
		string(DataSourceTechnologyMsdynamics),
		string(DataSourceTechnologyMssql),
		string(DataSourceTechnologyMssqlCds),
		string(DataSourceTechnologyMysql),
		string(DataSourceTechnologyOracle),
		string(DataSourceTechnologyPostgresql),
		string(DataSourceTechnologyRedshift),
		string(DataSourceTechnologyS3),
		string(DataSourceTechnologySaperp),
		string(DataSourceTechnologyTeradata),
	}

	// InKnownDataSourceTechnology is an ozzo-validator for DataSourceTechnology
	InKnownDataSourceTechnology = validation.In(
		DataSourceTechnologyBigquery,
		DataSourceTechnologyClickhouse,
		DataSourceTechnologyDb2,
		DataSourceTechnologyHana,
		DataSourceTechnologyHive,
		DataSourceTechnologyHsqldbFoodmart,
		DataSourceTechnologyMsdynamics,
		DataSourceTechnologyMssql,
		DataSourceTechnologyMssqlCds,
		DataSourceTechnologyMysql,
		DataSourceTechnologyOracle,
		DataSourceTechnologyPostgresql,
		DataSourceTechnologyRedshift,
		DataSourceTechnologyS3,
		DataSourceTechnologySaperp,
		DataSourceTechnologyTeradata,
	)
)
