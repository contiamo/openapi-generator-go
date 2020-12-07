// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	"time"
)

// QueryHistoryResponse is an object.
type QueryHistoryResponse struct {
	// BackendLogicalPlan:
	BackendLogicalPlan string `json:"backendLogicalPlan,omitempty"`
	// BackendPhysicalPlan:
	BackendPhysicalPlan string `json:"backendPhysicalPlan,omitempty"`
	// CompletedAt:
	CompletedAt time.Time `json:"completedAt,omitempty"`
	// CompletionStatus:
	CompletionStatus struct {
		Msg  string `json:"msg,omitempty"`
		Type string `json:"type,omitempty"`
	} `json:"completionStatus,omitempty"`
	// CustomReference:
	CustomReference string `json:"customReference,omitempty"`
	// DataSourceId:
	DataSourceId string `json:"dataSourceId,omitempty"`
	// Id:
	Id string `json:"id"`
	// Origin:
	Origin string `json:"origin"`
	// Plan:
	Plan string `json:"plan,omitempty"`
	// Query:
	Query string `json:"query"`
	// StartedAt:
	StartedAt time.Time `json:"startedAt"`
	// Type:
	Type QueryHistoryRecordType `json:"type"`
	// UserId:
	UserId string `json:"userId,omitempty"`
}

// GetBackendLogicalPlan returns the BackendLogicalPlan property
func (m QueryHistoryResponse) GetBackendLogicalPlan() string {
	return m.BackendLogicalPlan
}

// SetBackendLogicalPlan sets the BackendLogicalPlan property
func (m QueryHistoryResponse) SetBackendLogicalPlan(val string) {
	m.BackendLogicalPlan = val
}

// GetBackendPhysicalPlan returns the BackendPhysicalPlan property
func (m QueryHistoryResponse) GetBackendPhysicalPlan() string {
	return m.BackendPhysicalPlan
}

// SetBackendPhysicalPlan sets the BackendPhysicalPlan property
func (m QueryHistoryResponse) SetBackendPhysicalPlan(val string) {
	m.BackendPhysicalPlan = val
}

// GetCompletedAt returns the CompletedAt property
func (m QueryHistoryResponse) GetCompletedAt() time.Time {
	return m.CompletedAt
}

// SetCompletedAt sets the CompletedAt property
func (m QueryHistoryResponse) SetCompletedAt(val time.Time) {
	m.CompletedAt = val
}

// GetCompletionStatus returns the CompletionStatus property
func (m QueryHistoryResponse) GetCompletionStatus() struct {
	Msg  string `json:"msg,omitempty"`
	Type string `json:"type,omitempty"`
} {
	return m.CompletionStatus
}

// SetCompletionStatus sets the CompletionStatus property
func (m QueryHistoryResponse) SetCompletionStatus(val struct {
	Msg  string `json:"msg,omitempty"`
	Type string `json:"type,omitempty"`
}) {
	m.CompletionStatus = val
}

// GetCustomReference returns the CustomReference property
func (m QueryHistoryResponse) GetCustomReference() string {
	return m.CustomReference
}

// SetCustomReference sets the CustomReference property
func (m QueryHistoryResponse) SetCustomReference(val string) {
	m.CustomReference = val
}

// GetDataSourceId returns the DataSourceId property
func (m QueryHistoryResponse) GetDataSourceId() string {
	return m.DataSourceId
}

// SetDataSourceId sets the DataSourceId property
func (m QueryHistoryResponse) SetDataSourceId(val string) {
	m.DataSourceId = val
}

// GetId returns the Id property
func (m QueryHistoryResponse) GetId() string {
	return m.Id
}

// SetId sets the Id property
func (m QueryHistoryResponse) SetId(val string) {
	m.Id = val
}

// GetOrigin returns the Origin property
func (m QueryHistoryResponse) GetOrigin() string {
	return m.Origin
}

// SetOrigin sets the Origin property
func (m QueryHistoryResponse) SetOrigin(val string) {
	m.Origin = val
}

// GetPlan returns the Plan property
func (m QueryHistoryResponse) GetPlan() string {
	return m.Plan
}

// SetPlan sets the Plan property
func (m QueryHistoryResponse) SetPlan(val string) {
	m.Plan = val
}

// GetQuery returns the Query property
func (m QueryHistoryResponse) GetQuery() string {
	return m.Query
}

// SetQuery sets the Query property
func (m QueryHistoryResponse) SetQuery(val string) {
	m.Query = val
}

// GetStartedAt returns the StartedAt property
func (m QueryHistoryResponse) GetStartedAt() time.Time {
	return m.StartedAt
}

// SetStartedAt sets the StartedAt property
func (m QueryHistoryResponse) SetStartedAt(val time.Time) {
	m.StartedAt = val
}

// GetType returns the Type property
func (m QueryHistoryResponse) GetType() QueryHistoryRecordType {
	return m.Type
}

// SetType sets the Type property
func (m QueryHistoryResponse) SetType(val QueryHistoryRecordType) {
	m.Type = val
}

// GetUserId returns the UserId property
func (m QueryHistoryResponse) GetUserId() string {
	return m.UserId
}

// SetUserId sets the UserId property
func (m QueryHistoryResponse) SetUserId(val string) {
	m.UserId = val
}
