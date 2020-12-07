// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// TableProfileRequest is an object.
type TableProfileRequest struct {
	// SampleSize: Amount of rows to take for the profiling session
	SampleSize int32 `json:"sampleSize,omitempty"`
}

// GetSampleSize returns the SampleSize property
func (m TableProfileRequest) GetSampleSize() int32 {
	return m.SampleSize
}

// SetSampleSize sets the SampleSize property
func (m TableProfileRequest) SetSampleSize(val int32) {
	m.SampleSize = val
}
