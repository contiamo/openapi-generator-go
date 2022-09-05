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

// ConnectionTechnology is an enum. The connection technology is either the technology value of the related data source or the integration type
type ConnectionTechnology string

// Validate implements basic validation for this model
func (m ConnectionTechnology) Validate() error {
	return InKnownConnectionTechnology.Validate(m)
}

var (
	ConnectionTechnologyGeneric    ConnectionTechnology = "generic"
	ConnectionTechnologyMssql      ConnectionTechnology = "mssql"
	ConnectionTechnologyMysql      ConnectionTechnology = "mysql"
	ConnectionTechnologyOracle     ConnectionTechnology = "oracle"
	ConnectionTechnologyOther      ConnectionTechnology = "other"
	ConnectionTechnologyPostgresql ConnectionTechnology = "postgresql"
	ConnectionTechnologyTableau    ConnectionTechnology = "tableau"

	// KnownConnectionTechnology is the list of valid ConnectionTechnology
	KnownConnectionTechnology = []ConnectionTechnology{
		ConnectionTechnologyGeneric,
		ConnectionTechnologyMssql,
		ConnectionTechnologyMysql,
		ConnectionTechnologyOracle,
		ConnectionTechnologyOther,
		ConnectionTechnologyPostgresql,
		ConnectionTechnologyTableau,
	}
	// KnownConnectionTechnologyString is the list of valid ConnectionTechnology as string
	KnownConnectionTechnologyString = []string{
		string(ConnectionTechnologyGeneric),
		string(ConnectionTechnologyMssql),
		string(ConnectionTechnologyMysql),
		string(ConnectionTechnologyOracle),
		string(ConnectionTechnologyOther),
		string(ConnectionTechnologyPostgresql),
		string(ConnectionTechnologyTableau),
	}

	// InKnownConnectionTechnology is an ozzo-validator for ConnectionTechnology
	InKnownConnectionTechnology = validation.In(
		ConnectionTechnologyGeneric,
		ConnectionTechnologyMssql,
		ConnectionTechnologyMysql,
		ConnectionTechnologyOracle,
		ConnectionTechnologyOther,
		ConnectionTechnologyPostgresql,
		ConnectionTechnologyTableau,
	)
)
