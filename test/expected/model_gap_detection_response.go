// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// GapDetectionResponse is an object.
type GapDetectionResponse struct {
	// Data:
	Data []struct {
		Count    float32 `json:"count,omitempty"`
		Datetime string  `json:"datetime,omitempty"`
	} `json:"data"`
	// Gaps:
	Gaps []struct {
		Duration float32 `json:"duration,omitempty"`
		End      string  `json:"end,omitempty"`
		Start    string  `json:"start,omitempty"`
	} `json:"gaps"`
	// Interval:
	Interval Interval `json:"interval"`
	// MaxCount:
	MaxCount float32 `json:"maxCount"`
	// MeanGap:
	MeanGap float32 `json:"meanGap,omitempty"`
	// MinCount:
	MinCount float32 `json:"minCount"`
}

// GetData returns the Data property
func (m GapDetectionResponse) GetData() []struct {
	Count    float32 `json:"count,omitempty"`
	Datetime string  `json:"datetime,omitempty"`
} {
	return m.Data
}

// SetData sets the Data property
func (m GapDetectionResponse) SetData(val []struct {
	Count    float32 `json:"count,omitempty"`
	Datetime string  `json:"datetime,omitempty"`
}) {
	m.Data = val
}

// GetGaps returns the Gaps property
func (m GapDetectionResponse) GetGaps() []struct {
	Duration float32 `json:"duration,omitempty"`
	End      string  `json:"end,omitempty"`
	Start    string  `json:"start,omitempty"`
} {
	return m.Gaps
}

// SetGaps sets the Gaps property
func (m GapDetectionResponse) SetGaps(val []struct {
	Duration float32 `json:"duration,omitempty"`
	End      string  `json:"end,omitempty"`
	Start    string  `json:"start,omitempty"`
}) {
	m.Gaps = val
}

// GetInterval returns the Interval property
func (m GapDetectionResponse) GetInterval() Interval {
	return m.Interval
}

// SetInterval sets the Interval property
func (m GapDetectionResponse) SetInterval(val Interval) {
	m.Interval = val
}

// GetMaxCount returns the MaxCount property
func (m GapDetectionResponse) GetMaxCount() float32 {
	return m.MaxCount
}

// SetMaxCount sets the MaxCount property
func (m GapDetectionResponse) SetMaxCount(val float32) {
	m.MaxCount = val
}

// GetMeanGap returns the MeanGap property
func (m GapDetectionResponse) GetMeanGap() float32 {
	return m.MeanGap
}

// SetMeanGap sets the MeanGap property
func (m GapDetectionResponse) SetMeanGap(val float32) {
	m.MeanGap = val
}

// GetMinCount returns the MinCount property
func (m GapDetectionResponse) GetMinCount() float32 {
	return m.MinCount
}

// SetMinCount sets the MinCount property
func (m GapDetectionResponse) SetMinCount(val float32) {
	m.MinCount = val
}
