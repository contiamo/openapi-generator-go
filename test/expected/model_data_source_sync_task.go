// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	"time"
)

// DataSourceSyncTask is an object.
type DataSourceSyncTask struct {
	// CreatedAt:
	CreatedAt time.Time `json:"createdAt"`
	// FinishedAt:
	FinishedAt *time.Time `json:"finishedAt"`
	// Message:
	Message *string `json:"message"`
	// StartedAt:
	StartedAt *time.Time `json:"startedAt"`
	// Status:
	Status string `json:"status"`
	// TaskId:
	TaskId string `json:"taskId"`
	// UpdatedAt:
	UpdatedAt *time.Time `json:"updatedAt"`
}

// GetCreatedAt returns the CreatedAt property
func (m DataSourceSyncTask) GetCreatedAt() time.Time {
	return m.CreatedAt
}

// SetCreatedAt sets the CreatedAt property
func (m DataSourceSyncTask) SetCreatedAt(val time.Time) {
	m.CreatedAt = val
}

// GetFinishedAt returns the FinishedAt property
func (m DataSourceSyncTask) GetFinishedAt() *time.Time {
	return m.FinishedAt
}

// SetFinishedAt sets the FinishedAt property
func (m DataSourceSyncTask) SetFinishedAt(val *time.Time) {
	m.FinishedAt = val
}

// GetMessage returns the Message property
func (m DataSourceSyncTask) GetMessage() *string {
	return m.Message
}

// SetMessage sets the Message property
func (m DataSourceSyncTask) SetMessage(val *string) {
	m.Message = val
}

// GetStartedAt returns the StartedAt property
func (m DataSourceSyncTask) GetStartedAt() *time.Time {
	return m.StartedAt
}

// SetStartedAt sets the StartedAt property
func (m DataSourceSyncTask) SetStartedAt(val *time.Time) {
	m.StartedAt = val
}

// GetStatus returns the Status property
func (m DataSourceSyncTask) GetStatus() string {
	return m.Status
}

// SetStatus sets the Status property
func (m DataSourceSyncTask) SetStatus(val string) {
	m.Status = val
}

// GetTaskId returns the TaskId property
func (m DataSourceSyncTask) GetTaskId() string {
	return m.TaskId
}

// SetTaskId sets the TaskId property
func (m DataSourceSyncTask) SetTaskId(val string) {
	m.TaskId = val
}

// GetUpdatedAt returns the UpdatedAt property
func (m DataSourceSyncTask) GetUpdatedAt() *time.Time {
	return m.UpdatedAt
}

// SetUpdatedAt sets the UpdatedAt property
func (m DataSourceSyncTask) SetUpdatedAt(val *time.Time) {
	m.UpdatedAt = val
}
