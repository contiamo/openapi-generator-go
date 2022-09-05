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

// IntegrationType is an enum.
type IntegrationType string

// Validate implements basic validation for this model
func (m IntegrationType) Validate() error {
	return InKnownIntegrationType.Validate(m)
}

var (
	IntegrationTypeGeneric IntegrationType = "generic"
	IntegrationTypeTableau IntegrationType = "tableau"

	// KnownIntegrationType is the list of valid IntegrationType
	KnownIntegrationType = []IntegrationType{
		IntegrationTypeGeneric,
		IntegrationTypeTableau,
	}
	// KnownIntegrationTypeString is the list of valid IntegrationType as string
	KnownIntegrationTypeString = []string{
		string(IntegrationTypeGeneric),
		string(IntegrationTypeTableau),
	}

	// InKnownIntegrationType is an ozzo-validator for IntegrationType
	InKnownIntegrationType = validation.In(
		IntegrationTypeGeneric,
		IntegrationTypeTableau,
	)
)
