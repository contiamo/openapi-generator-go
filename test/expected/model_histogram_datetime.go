// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// HistogramDatetime is an object.
type HistogramDatetime struct {
	// Bins:
	Bins []string `json:"bins"`
	// Counts:
	Counts []float32 `json:"counts"`
}

// GetBins returns the Bins property
func (m HistogramDatetime) GetBins() []string {
	return m.Bins
}

// SetBins sets the Bins property
func (m HistogramDatetime) SetBins(val []string) {
	m.Bins = val
}

// GetCounts returns the Counts property
func (m HistogramDatetime) GetCounts() []float32 {
	return m.Counts
}

// SetCounts sets the Counts property
func (m HistogramDatetime) SetCounts(val []float32) {
	m.Counts = val
}
