// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// RelationList is an object. A list of relations (dependants and dependencies) for a given resource.
type RelationList struct {
	// Dependants: List of resources that depend on the current one
	Dependants []RelationResponse `json:"dependants"`
	// Dependencies: List of resources that the current resource depends on
	Dependencies []RelationResponse `json:"dependencies"`
}

// GetDependants returns the Dependants property
func (m RelationList) GetDependants() []RelationResponse {
	return m.Dependants
}

// SetDependants sets the Dependants property
func (m RelationList) SetDependants(val []RelationResponse) {
	m.Dependants = val
}

// GetDependencies returns the Dependencies property
func (m RelationList) GetDependencies() []RelationResponse {
	return m.Dependencies
}

// SetDependencies sets the Dependencies property
func (m RelationList) SetDependencies(val []RelationResponse) {
	m.Dependencies = val
}
