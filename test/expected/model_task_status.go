// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// TaskStatus is an enum. Current state of the task
type TaskStatus string

var (
	TaskStatusCancelled TaskStatus = "cancelled"
	TaskStatusFailed    TaskStatus = "failed"
	TaskStatusFinished  TaskStatus = "finished"
	TaskStatusRunning   TaskStatus = "running"
	TaskStatusWaiting   TaskStatus = "waiting"

	// KnownTaskStatus is the list of valid TaskStatus
	KnownTaskStatus = []TaskStatus{
		TaskStatusCancelled,
		TaskStatusFailed,
		TaskStatusFinished,
		TaskStatusRunning,
		TaskStatusWaiting,
	}
	// KnownTaskStatusString is the list of valid TaskStatus as string
	KnownTaskStatusString = []string{
		string(TaskStatusCancelled),
		string(TaskStatusFailed),
		string(TaskStatusFinished),
		string(TaskStatusRunning),
		string(TaskStatusWaiting),
	}

	// InKnownTaskStatus is an ozzo-validator for TaskStatus
	InKnownTaskStatus = validation.In(
		TaskStatusCancelled,
		TaskStatusFailed,
		TaskStatusFinished,
		TaskStatusRunning,
		TaskStatusWaiting,
	)
)
