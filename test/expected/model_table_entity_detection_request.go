// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// TableEntityDetectionRequest is an object.
type TableEntityDetectionRequest struct {
	// SampleSize: Amount of rows to take for the entity detection session
	SampleSize int32 `json:"sampleSize,omitempty"`
}

// GetSampleSize returns the SampleSize property
func (m TableEntityDetectionRequest) GetSampleSize() int32 {
	return m.SampleSize
}

// SetSampleSize sets the SampleSize property
func (m TableEntityDetectionRequest) SetSampleSize(val int32) {
	m.SampleSize = val
}
