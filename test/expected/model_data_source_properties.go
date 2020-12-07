// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// DataSourceProperties is an object.
type DataSourceProperties struct {
	// LastUpdateErrorMessage: contains an error message indicating why this data source schema could not be updated
	LastUpdateErrorMessage string `json:"lastUpdateErrorMessage,omitempty"`
	// Technology: Type of a data source
	Technology DataSourceTechnology `json:"technology,omitempty"`
	// Type: Type of a data source
	Type DataSourceType `json:"type"`
}

// GetLastUpdateErrorMessage returns the LastUpdateErrorMessage property
func (m DataSourceProperties) GetLastUpdateErrorMessage() string {
	return m.LastUpdateErrorMessage
}

// SetLastUpdateErrorMessage sets the LastUpdateErrorMessage property
func (m DataSourceProperties) SetLastUpdateErrorMessage(val string) {
	m.LastUpdateErrorMessage = val
}

// GetTechnology returns the Technology property
func (m DataSourceProperties) GetTechnology() DataSourceTechnology {
	return m.Technology
}

// SetTechnology sets the Technology property
func (m DataSourceProperties) SetTechnology(val DataSourceTechnology) {
	m.Technology = val
}

// GetType returns the Type property
func (m DataSourceProperties) GetType() DataSourceType {
	return m.Type
}

// SetType sets the Type property
func (m DataSourceProperties) SetType(val DataSourceType) {
	m.Type = val
}
