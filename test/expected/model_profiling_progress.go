// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

import (
	"time"
)

// ProfilingProgress is an object.
type ProfilingProgress struct {
	// Error:
	Error string `json:"error,omitempty"`
	// FinishedAt:
	FinishedAt *time.Time `json:"finishedAt"`
	// Processed:
	Processed int32 `json:"processed"`
	// StartedAt:
	StartedAt *time.Time `json:"startedAt"`
	// Total:
	Total int32 `json:"total"`
}

// GetError returns the Error property
func (m ProfilingProgress) GetError() string {
	return m.Error
}

// SetError sets the Error property
func (m ProfilingProgress) SetError(val string) {
	m.Error = val
}

// GetFinishedAt returns the FinishedAt property
func (m ProfilingProgress) GetFinishedAt() *time.Time {
	return m.FinishedAt
}

// SetFinishedAt sets the FinishedAt property
func (m ProfilingProgress) SetFinishedAt(val *time.Time) {
	m.FinishedAt = val
}

// GetProcessed returns the Processed property
func (m ProfilingProgress) GetProcessed() int32 {
	return m.Processed
}

// SetProcessed sets the Processed property
func (m ProfilingProgress) SetProcessed(val int32) {
	m.Processed = val
}

// GetStartedAt returns the StartedAt property
func (m ProfilingProgress) GetStartedAt() *time.Time {
	return m.StartedAt
}

// SetStartedAt sets the StartedAt property
func (m ProfilingProgress) SetStartedAt(val *time.Time) {
	m.StartedAt = val
}

// GetTotal returns the Total property
func (m ProfilingProgress) GetTotal() int32 {
	return m.Total
}

// SetTotal sets the Total property
func (m ProfilingProgress) SetTotal(val int32) {
	m.Total = val
}
