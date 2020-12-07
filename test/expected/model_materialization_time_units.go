// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// MaterializationTimeUnits is an enum. The time unit associated with expiresIn; can only be null if expiresIn is 0.
type MaterializationTimeUnits string

var (
	MaterializationTimeUnitsDays  MaterializationTimeUnits = "days"
	MaterializationTimeUnitsHours MaterializationTimeUnits = "hours"
	MaterializationTimeUnitsWeeks MaterializationTimeUnits = "weeks"

	// KnownMaterializationTimeUnits is the list of valid MaterializationTimeUnits
	KnownMaterializationTimeUnits = []MaterializationTimeUnits{
		MaterializationTimeUnitsDays,
		MaterializationTimeUnitsHours,
		MaterializationTimeUnitsWeeks,
	}
	// KnownMaterializationTimeUnitsString is the list of valid MaterializationTimeUnits as string
	KnownMaterializationTimeUnitsString = []string{
		string(MaterializationTimeUnitsDays),
		string(MaterializationTimeUnitsHours),
		string(MaterializationTimeUnitsWeeks),
	}

	// InKnownMaterializationTimeUnits is an ozzo-validator for MaterializationTimeUnits
	InKnownMaterializationTimeUnits = validation.In(
		MaterializationTimeUnitsDays,
		MaterializationTimeUnitsHours,
		MaterializationTimeUnitsWeeks,
	)
)
