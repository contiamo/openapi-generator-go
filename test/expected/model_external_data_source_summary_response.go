// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	"time"
)

// ExternalDataSourceSummaryResponse is an object.
type ExternalDataSourceSummaryResponse struct {
	// CreatedAt:
	CreatedAt time.Time `json:"createdAt"`
	// Id:
	Id string `json:"id"`
	// LastUpdateErrorMessage: contains an error message indicating why this data source schema could not be updated
	LastUpdateErrorMessage string `json:"lastUpdateErrorMessage,omitempty"`
	// Name:
	Name string `json:"name"`
	// ProjectId:
	ProjectId string `json:"projectId"`
	// Technology: Type of a data source
	Technology DataSourceTechnology `json:"technology"`
	// UpdatedAt:
	UpdatedAt time.Time `json:"updatedAt"`
}

// GetCreatedAt returns the CreatedAt property
func (m ExternalDataSourceSummaryResponse) GetCreatedAt() time.Time {
	return m.CreatedAt
}

// SetCreatedAt sets the CreatedAt property
func (m ExternalDataSourceSummaryResponse) SetCreatedAt(val time.Time) {
	m.CreatedAt = val
}

// GetId returns the Id property
func (m ExternalDataSourceSummaryResponse) GetId() string {
	return m.Id
}

// SetId sets the Id property
func (m ExternalDataSourceSummaryResponse) SetId(val string) {
	m.Id = val
}

// GetLastUpdateErrorMessage returns the LastUpdateErrorMessage property
func (m ExternalDataSourceSummaryResponse) GetLastUpdateErrorMessage() string {
	return m.LastUpdateErrorMessage
}

// SetLastUpdateErrorMessage sets the LastUpdateErrorMessage property
func (m ExternalDataSourceSummaryResponse) SetLastUpdateErrorMessage(val string) {
	m.LastUpdateErrorMessage = val
}

// GetName returns the Name property
func (m ExternalDataSourceSummaryResponse) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m ExternalDataSourceSummaryResponse) SetName(val string) {
	m.Name = val
}

// GetProjectId returns the ProjectId property
func (m ExternalDataSourceSummaryResponse) GetProjectId() string {
	return m.ProjectId
}

// SetProjectId sets the ProjectId property
func (m ExternalDataSourceSummaryResponse) SetProjectId(val string) {
	m.ProjectId = val
}

// GetTechnology returns the Technology property
func (m ExternalDataSourceSummaryResponse) GetTechnology() DataSourceTechnology {
	return m.Technology
}

// SetTechnology sets the Technology property
func (m ExternalDataSourceSummaryResponse) SetTechnology(val DataSourceTechnology) {
	m.Technology = val
}

// GetUpdatedAt returns the UpdatedAt property
func (m ExternalDataSourceSummaryResponse) GetUpdatedAt() time.Time {
	return m.UpdatedAt
}

// SetUpdatedAt sets the UpdatedAt property
func (m ExternalDataSourceSummaryResponse) SetUpdatedAt(val time.Time) {
	m.UpdatedAt = val
}
