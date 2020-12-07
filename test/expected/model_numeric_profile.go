// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// NumericProfile is an object.
type NumericProfile struct {
	// DataType:
	DataType string `json:"dataType"`
	// Histogram:
	Histogram Histogram `json:"histogram"`
	// Max:
	Max float32 `json:"max,omitempty"`
	// Mean:
	Mean float32 `json:"mean,omitempty"`
	// Min:
	Min float32 `json:"min,omitempty"`
	// NullCount:
	NullCount float32 `json:"nullCount"`
	// NullPercentage:
	NullPercentage float32 `json:"nullPercentage"`
	// Std:
	Std float32 `json:"std,omitempty"`
	// Type:
	Type string `json:"type,omitempty"`
}

// GetDataType returns the DataType property
func (m NumericProfile) GetDataType() string {
	return m.DataType
}

// SetDataType sets the DataType property
func (m NumericProfile) SetDataType(val string) {
	m.DataType = val
}

// GetHistogram returns the Histogram property
func (m NumericProfile) GetHistogram() Histogram {
	return m.Histogram
}

// SetHistogram sets the Histogram property
func (m NumericProfile) SetHistogram(val Histogram) {
	m.Histogram = val
}

// GetMax returns the Max property
func (m NumericProfile) GetMax() float32 {
	return m.Max
}

// SetMax sets the Max property
func (m NumericProfile) SetMax(val float32) {
	m.Max = val
}

// GetMean returns the Mean property
func (m NumericProfile) GetMean() float32 {
	return m.Mean
}

// SetMean sets the Mean property
func (m NumericProfile) SetMean(val float32) {
	m.Mean = val
}

// GetMin returns the Min property
func (m NumericProfile) GetMin() float32 {
	return m.Min
}

// SetMin sets the Min property
func (m NumericProfile) SetMin(val float32) {
	m.Min = val
}

// GetNullCount returns the NullCount property
func (m NumericProfile) GetNullCount() float32 {
	return m.NullCount
}

// SetNullCount sets the NullCount property
func (m NumericProfile) SetNullCount(val float32) {
	m.NullCount = val
}

// GetNullPercentage returns the NullPercentage property
func (m NumericProfile) GetNullPercentage() float32 {
	return m.NullPercentage
}

// SetNullPercentage sets the NullPercentage property
func (m NumericProfile) SetNullPercentage(val float32) {
	m.NullPercentage = val
}

// GetStd returns the Std property
func (m NumericProfile) GetStd() float32 {
	return m.Std
}

// SetStd sets the Std property
func (m NumericProfile) SetStd(val float32) {
	m.Std = val
}

// GetType returns the Type property
func (m NumericProfile) GetType() string {
	return m.Type
}

// SetType sets the Type property
func (m NumericProfile) SetType(val string) {
	m.Type = val
}
