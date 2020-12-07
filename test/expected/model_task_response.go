// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	"time"
)

// TaskResponse is an object.
type TaskResponse struct {
	// CollectionId:
	CollectionId string `json:"collectionId"`
	// CreatedAt:
	CreatedAt time.Time `json:"createdAt"`
	// FileId:
	FileId string `json:"fileId"`
	// FileName:
	FileName string `json:"fileName"`
	// FinishedAt:
	FinishedAt *time.Time `json:"finishedAt"`
	// InsertMode: Defines how to apply rows to a table.
	// upsert: insert or update rows from the upload to the table
	// append: insert rows from the upload to the table.
	InsertMode IngestInsertModeType `json:"insertMode"`
	// StartedAt:
	StartedAt *time.Time `json:"startedAt"`
	// Status: Current state of the task
	Status TaskStatus `json:"status"`
	// StatusData:
	StatusData TaskStatusData `json:"statusData"`
	// TableId:
	TableId string `json:"tableId"`
	// TableName:
	TableName string `json:"tableName"`
	// TaskId:
	TaskId string `json:"taskId"`
	// UpsertOn: rows used to match file data to a possibly existing row, only used and required if insertMode is `upsert`
	UpsertOn []string `json:"upsertOn"`
}

// GetCollectionId returns the CollectionId property
func (m TaskResponse) GetCollectionId() string {
	return m.CollectionId
}

// SetCollectionId sets the CollectionId property
func (m TaskResponse) SetCollectionId(val string) {
	m.CollectionId = val
}

// GetCreatedAt returns the CreatedAt property
func (m TaskResponse) GetCreatedAt() time.Time {
	return m.CreatedAt
}

// SetCreatedAt sets the CreatedAt property
func (m TaskResponse) SetCreatedAt(val time.Time) {
	m.CreatedAt = val
}

// GetFileId returns the FileId property
func (m TaskResponse) GetFileId() string {
	return m.FileId
}

// SetFileId sets the FileId property
func (m TaskResponse) SetFileId(val string) {
	m.FileId = val
}

// GetFileName returns the FileName property
func (m TaskResponse) GetFileName() string {
	return m.FileName
}

// SetFileName sets the FileName property
func (m TaskResponse) SetFileName(val string) {
	m.FileName = val
}

// GetFinishedAt returns the FinishedAt property
func (m TaskResponse) GetFinishedAt() *time.Time {
	return m.FinishedAt
}

// SetFinishedAt sets the FinishedAt property
func (m TaskResponse) SetFinishedAt(val *time.Time) {
	m.FinishedAt = val
}

// GetInsertMode returns the InsertMode property
func (m TaskResponse) GetInsertMode() IngestInsertModeType {
	return m.InsertMode
}

// SetInsertMode sets the InsertMode property
func (m TaskResponse) SetInsertMode(val IngestInsertModeType) {
	m.InsertMode = val
}

// GetStartedAt returns the StartedAt property
func (m TaskResponse) GetStartedAt() *time.Time {
	return m.StartedAt
}

// SetStartedAt sets the StartedAt property
func (m TaskResponse) SetStartedAt(val *time.Time) {
	m.StartedAt = val
}

// GetStatus returns the Status property
func (m TaskResponse) GetStatus() TaskStatus {
	return m.Status
}

// SetStatus sets the Status property
func (m TaskResponse) SetStatus(val TaskStatus) {
	m.Status = val
}

// GetStatusData returns the StatusData property
func (m TaskResponse) GetStatusData() TaskStatusData {
	return m.StatusData
}

// SetStatusData sets the StatusData property
func (m TaskResponse) SetStatusData(val TaskStatusData) {
	m.StatusData = val
}

// GetTableId returns the TableId property
func (m TaskResponse) GetTableId() string {
	return m.TableId
}

// SetTableId sets the TableId property
func (m TaskResponse) SetTableId(val string) {
	m.TableId = val
}

// GetTableName returns the TableName property
func (m TaskResponse) GetTableName() string {
	return m.TableName
}

// SetTableName sets the TableName property
func (m TaskResponse) SetTableName(val string) {
	m.TableName = val
}

// GetTaskId returns the TaskId property
func (m TaskResponse) GetTaskId() string {
	return m.TaskId
}

// SetTaskId sets the TaskId property
func (m TaskResponse) SetTaskId(val string) {
	m.TaskId = val
}

// GetUpsertOn returns the UpsertOn property
func (m TaskResponse) GetUpsertOn() []string {
	return m.UpsertOn
}

// SetUpsertOn sets the UpsertOn property
func (m TaskResponse) SetUpsertOn(val []string) {
	m.UpsertOn = val
}
