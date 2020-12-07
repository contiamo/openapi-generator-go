// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// PredefinedEntity is an enum. The available predefined entity types that the Profiler can detect, as opposed to any potential user-defined entity types and rules.
type PredefinedEntity string

var (
	PredefinedEntityCREDITCARD      PredefinedEntity = "CREDIT_CARD"
	PredefinedEntityCRYPTO          PredefinedEntity = "CRYPTO"
	PredefinedEntityDATETIME        PredefinedEntity = "DATE_TIME"
	PredefinedEntityDOMAINNAME      PredefinedEntity = "DOMAIN_NAME"
	PredefinedEntityEMAILADDRESS    PredefinedEntity = "EMAIL_ADDRESS"
	PredefinedEntityGERMANIDCARD    PredefinedEntity = "GERMAN_ID_CARD"
	PredefinedEntityIBANCODE        PredefinedEntity = "IBAN_CODE"
	PredefinedEntityIPADDRESS       PredefinedEntity = "IP_ADDRESS"
	PredefinedEntityLOCATION        PredefinedEntity = "LOCATION"
	PredefinedEntityNRP             PredefinedEntity = "NRP"
	PredefinedEntityPERSON          PredefinedEntity = "PERSON"
	PredefinedEntityPHONENUMBER     PredefinedEntity = "PHONE_NUMBER"
	PredefinedEntitySGNRICFIN       PredefinedEntity = "SG_NRIC_FIN"
	PredefinedEntityUKNHS           PredefinedEntity = "UK_NHS"
	PredefinedEntityUSBANKNUMBER    PredefinedEntity = "US_BANK_NUMBER"
	PredefinedEntityUSDRIVERLICENSE PredefinedEntity = "US_DRIVER_LICENSE"
	PredefinedEntityUSITIN          PredefinedEntity = "US_ITIN"
	PredefinedEntityUSPASSPORT      PredefinedEntity = "US_PASSPORT"
	PredefinedEntityUSSSN           PredefinedEntity = "US_SSN"

	// KnownPredefinedEntity is the list of valid PredefinedEntity
	KnownPredefinedEntity = []PredefinedEntity{
		PredefinedEntityCREDITCARD,
		PredefinedEntityCRYPTO,
		PredefinedEntityDATETIME,
		PredefinedEntityDOMAINNAME,
		PredefinedEntityEMAILADDRESS,
		PredefinedEntityGERMANIDCARD,
		PredefinedEntityIBANCODE,
		PredefinedEntityIPADDRESS,
		PredefinedEntityLOCATION,
		PredefinedEntityNRP,
		PredefinedEntityPERSON,
		PredefinedEntityPHONENUMBER,
		PredefinedEntitySGNRICFIN,
		PredefinedEntityUKNHS,
		PredefinedEntityUSBANKNUMBER,
		PredefinedEntityUSDRIVERLICENSE,
		PredefinedEntityUSITIN,
		PredefinedEntityUSPASSPORT,
		PredefinedEntityUSSSN,
	}
	// KnownPredefinedEntityString is the list of valid PredefinedEntity as string
	KnownPredefinedEntityString = []string{
		string(PredefinedEntityCREDITCARD),
		string(PredefinedEntityCRYPTO),
		string(PredefinedEntityDATETIME),
		string(PredefinedEntityDOMAINNAME),
		string(PredefinedEntityEMAILADDRESS),
		string(PredefinedEntityGERMANIDCARD),
		string(PredefinedEntityIBANCODE),
		string(PredefinedEntityIPADDRESS),
		string(PredefinedEntityLOCATION),
		string(PredefinedEntityNRP),
		string(PredefinedEntityPERSON),
		string(PredefinedEntityPHONENUMBER),
		string(PredefinedEntitySGNRICFIN),
		string(PredefinedEntityUKNHS),
		string(PredefinedEntityUSBANKNUMBER),
		string(PredefinedEntityUSDRIVERLICENSE),
		string(PredefinedEntityUSITIN),
		string(PredefinedEntityUSPASSPORT),
		string(PredefinedEntityUSSSN),
	}

	// InKnownPredefinedEntity is an ozzo-validator for PredefinedEntity
	InKnownPredefinedEntity = validation.In(
		PredefinedEntityCREDITCARD,
		PredefinedEntityCRYPTO,
		PredefinedEntityDATETIME,
		PredefinedEntityDOMAINNAME,
		PredefinedEntityEMAILADDRESS,
		PredefinedEntityGERMANIDCARD,
		PredefinedEntityIBANCODE,
		PredefinedEntityIPADDRESS,
		PredefinedEntityLOCATION,
		PredefinedEntityNRP,
		PredefinedEntityPERSON,
		PredefinedEntityPHONENUMBER,
		PredefinedEntitySGNRICFIN,
		PredefinedEntityUKNHS,
		PredefinedEntityUSBANKNUMBER,
		PredefinedEntityUSDRIVERLICENSE,
		PredefinedEntityUSITIN,
		PredefinedEntityUSPASSPORT,
		PredefinedEntityUSSSN,
	)
)
