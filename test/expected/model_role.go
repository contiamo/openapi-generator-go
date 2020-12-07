// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// Role is an enum. A valid resource role identifier
type Role string

var (
	RoleOwner   Role = "owner"
	RoleQuery   Role = "query"
	RoleSteward Role = "steward"

	// KnownRole is the list of valid Role
	KnownRole = []Role{
		RoleOwner,
		RoleQuery,
		RoleSteward,
	}
	// KnownRoleString is the list of valid Role as string
	KnownRoleString = []string{
		string(RoleOwner),
		string(RoleQuery),
		string(RoleSteward),
	}

	// InKnownRole is an ozzo-validator for Role
	InKnownRole = validation.In(
		RoleOwner,
		RoleQuery,
		RoleSteward,
	)
)
