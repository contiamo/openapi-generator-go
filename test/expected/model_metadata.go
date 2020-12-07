// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// Metadata is an object.
type Metadata struct {
	// IsSample:
	IsSample bool `json:"isSample"`
	// SampleSize:
	SampleSize float32 `json:"sampleSize,omitempty"`
	// SampleType:
	SampleType string `json:"sampleType,omitempty"`
}

// GetIsSample returns the IsSample property
func (m Metadata) GetIsSample() bool {
	return m.IsSample
}

// SetIsSample sets the IsSample property
func (m Metadata) SetIsSample(val bool) {
	m.IsSample = val
}

// GetSampleSize returns the SampleSize property
func (m Metadata) GetSampleSize() float32 {
	return m.SampleSize
}

// SetSampleSize sets the SampleSize property
func (m Metadata) SetSampleSize(val float32) {
	m.SampleSize = val
}

// GetSampleType returns the SampleType property
func (m Metadata) GetSampleType() string {
	return m.SampleType
}

// SetSampleType sets the SampleType property
func (m Metadata) SetSampleType(val string) {
	m.SampleType = val
}
