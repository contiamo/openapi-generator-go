// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// DatetimeProfile is an object.
type DatetimeProfile struct {
	// DataType:
	DataType string `json:"dataType"`
	// Histogram:
	Histogram HistogramDatetime `json:"histogram"`
	// Interval:
	Interval Interval `json:"interval,omitempty"`
	// Max: Max date
	Max string `json:"max"`
	// Min: Min date
	Min string `json:"min"`
	// NullCount:
	NullCount float32 `json:"nullCount"`
	// NullPercentage:
	NullPercentage float32 `json:"nullPercentage"`
}

// GetDataType returns the DataType property
func (m DatetimeProfile) GetDataType() string {
	return m.DataType
}

// SetDataType sets the DataType property
func (m DatetimeProfile) SetDataType(val string) {
	m.DataType = val
}

// GetHistogram returns the Histogram property
func (m DatetimeProfile) GetHistogram() HistogramDatetime {
	return m.Histogram
}

// SetHistogram sets the Histogram property
func (m DatetimeProfile) SetHistogram(val HistogramDatetime) {
	m.Histogram = val
}

// GetInterval returns the Interval property
func (m DatetimeProfile) GetInterval() Interval {
	return m.Interval
}

// SetInterval sets the Interval property
func (m DatetimeProfile) SetInterval(val Interval) {
	m.Interval = val
}

// GetMax returns the Max property
func (m DatetimeProfile) GetMax() string {
	return m.Max
}

// SetMax sets the Max property
func (m DatetimeProfile) SetMax(val string) {
	m.Max = val
}

// GetMin returns the Min property
func (m DatetimeProfile) GetMin() string {
	return m.Min
}

// SetMin sets the Min property
func (m DatetimeProfile) SetMin(val string) {
	m.Min = val
}

// GetNullCount returns the NullCount property
func (m DatetimeProfile) GetNullCount() float32 {
	return m.NullCount
}

// SetNullCount sets the NullCount property
func (m DatetimeProfile) SetNullCount(val float32) {
	m.NullCount = val
}

// GetNullPercentage returns the NullPercentage property
func (m DatetimeProfile) GetNullPercentage() float32 {
	return m.NullPercentage
}

// SetNullPercentage sets the NullPercentage property
func (m DatetimeProfile) SetNullPercentage(val float32) {
	m.NullPercentage = val
}
