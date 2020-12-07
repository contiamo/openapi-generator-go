// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// Histogram is an object.
type Histogram struct {
	// Bins:
	Bins []float32 `json:"bins"`
	// Counts:
	Counts []float32 `json:"counts"`
}

// GetBins returns the Bins property
func (m Histogram) GetBins() []float32 {
	return m.Bins
}

// SetBins sets the Bins property
func (m Histogram) SetBins(val []float32) {
	m.Bins = val
}

// GetCounts returns the Counts property
func (m Histogram) GetCounts() []float32 {
	return m.Counts
}

// SetCounts sets the Counts property
func (m Histogram) SetCounts(val []float32) {
	m.Counts = val
}
