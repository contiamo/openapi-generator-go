// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// EntityDetectionRequest is an object.
type EntityDetectionRequest struct {
	// ProjectId:
	ProjectId string `json:"projectId"`
	// SampleSize:
	SampleSize float32 `json:"sampleSize,omitempty"`
	// Sql:
	Sql string `json:"sql"`
}

// GetProjectId returns the ProjectId property
func (m EntityDetectionRequest) GetProjectId() string {
	return m.ProjectId
}

// SetProjectId sets the ProjectId property
func (m EntityDetectionRequest) SetProjectId(val string) {
	m.ProjectId = val
}

// GetSampleSize returns the SampleSize property
func (m EntityDetectionRequest) GetSampleSize() float32 {
	return m.SampleSize
}

// SetSampleSize sets the SampleSize property
func (m EntityDetectionRequest) SetSampleSize(val float32) {
	m.SampleSize = val
}

// GetSql returns the Sql property
func (m EntityDetectionRequest) GetSql() string {
	return m.Sql
}

// SetSql sets the Sql property
func (m EntityDetectionRequest) SetSql(val string) {
	m.Sql = val
}
