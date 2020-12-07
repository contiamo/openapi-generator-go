// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	"time"
)

// MaterializationTaskResponse is an object.
type MaterializationTaskResponse struct {
	// CreatedAt:
	CreatedAt time.Time `json:"createdAt"`
	// FinishedAt:
	FinishedAt *time.Time `json:"finishedAt"`
	// StartedAt:
	StartedAt *time.Time `json:"startedAt"`
	// Status: Current state of the task
	Status TaskStatus `json:"status"`
	// StatusData:
	StatusData TaskStatusData `json:"statusData"`
	// TaskId:
	TaskId string `json:"taskId"`
}

// GetCreatedAt returns the CreatedAt property
func (m MaterializationTaskResponse) GetCreatedAt() time.Time {
	return m.CreatedAt
}

// SetCreatedAt sets the CreatedAt property
func (m MaterializationTaskResponse) SetCreatedAt(val time.Time) {
	m.CreatedAt = val
}

// GetFinishedAt returns the FinishedAt property
func (m MaterializationTaskResponse) GetFinishedAt() *time.Time {
	return m.FinishedAt
}

// SetFinishedAt sets the FinishedAt property
func (m MaterializationTaskResponse) SetFinishedAt(val *time.Time) {
	m.FinishedAt = val
}

// GetStartedAt returns the StartedAt property
func (m MaterializationTaskResponse) GetStartedAt() *time.Time {
	return m.StartedAt
}

// SetStartedAt sets the StartedAt property
func (m MaterializationTaskResponse) SetStartedAt(val *time.Time) {
	m.StartedAt = val
}

// GetStatus returns the Status property
func (m MaterializationTaskResponse) GetStatus() TaskStatus {
	return m.Status
}

// SetStatus sets the Status property
func (m MaterializationTaskResponse) SetStatus(val TaskStatus) {
	m.Status = val
}

// GetStatusData returns the StatusData property
func (m MaterializationTaskResponse) GetStatusData() TaskStatusData {
	return m.StatusData
}

// SetStatusData sets the StatusData property
func (m MaterializationTaskResponse) SetStatusData(val TaskStatusData) {
	m.StatusData = val
}

// GetTaskId returns the TaskId property
func (m MaterializationTaskResponse) GetTaskId() string {
	return m.TaskId
}

// SetTaskId sets the TaskId property
func (m MaterializationTaskResponse) SetTaskId(val string) {
	m.TaskId = val
}
