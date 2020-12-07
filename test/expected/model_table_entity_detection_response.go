// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// TableEntityDetectionResponse is an object.
type TableEntityDetectionResponse struct {
	// Entities:
	Entities *EntityDetectionResponse `json:"entities,omitempty"`
	// Progress:
	Progress ProfilingProgress `json:"progress"`
}

// GetEntities returns the Entities property
func (m TableEntityDetectionResponse) GetEntities() *EntityDetectionResponse {
	return m.Entities
}

// SetEntities sets the Entities property
func (m TableEntityDetectionResponse) SetEntities(val *EntityDetectionResponse) {
	m.Entities = val
}

// GetProgress returns the Progress property
func (m TableEntityDetectionResponse) GetProgress() ProfilingProgress {
	return m.Progress
}

// SetProgress sets the Progress property
func (m TableEntityDetectionResponse) SetProgress(val ProfilingProgress) {
	m.Progress = val
}
