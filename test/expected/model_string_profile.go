// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// StringProfile is an object.
type StringProfile struct {
	// Cardinality:
	Cardinality string `json:"cardinality,omitempty"`
	// DataType:
	DataType string `json:"dataType"`
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
	// Top20:
	Top20 []struct {
		Key   string  `json:"key,omitempty"`
		Value float32 `json:"value,omitempty"`
	} `json:"top20"`
	// Type:
	Type string `json:"type,omitempty"`
	// Unique:
	Unique float32 `json:"unique"`
}

// GetCardinality returns the Cardinality property
func (m StringProfile) GetCardinality() string {
	return m.Cardinality
}

// SetCardinality sets the Cardinality property
func (m StringProfile) SetCardinality(val string) {
	m.Cardinality = val
}

// GetDataType returns the DataType property
func (m StringProfile) GetDataType() string {
	return m.DataType
}

// SetDataType sets the DataType property
func (m StringProfile) SetDataType(val string) {
	m.DataType = val
}

// GetMax returns the Max property
func (m StringProfile) GetMax() float32 {
	return m.Max
}

// SetMax sets the Max property
func (m StringProfile) SetMax(val float32) {
	m.Max = val
}

// GetMean returns the Mean property
func (m StringProfile) GetMean() float32 {
	return m.Mean
}

// SetMean sets the Mean property
func (m StringProfile) SetMean(val float32) {
	m.Mean = val
}

// GetMin returns the Min property
func (m StringProfile) GetMin() float32 {
	return m.Min
}

// SetMin sets the Min property
func (m StringProfile) SetMin(val float32) {
	m.Min = val
}

// GetNullCount returns the NullCount property
func (m StringProfile) GetNullCount() float32 {
	return m.NullCount
}

// SetNullCount sets the NullCount property
func (m StringProfile) SetNullCount(val float32) {
	m.NullCount = val
}

// GetNullPercentage returns the NullPercentage property
func (m StringProfile) GetNullPercentage() float32 {
	return m.NullPercentage
}

// SetNullPercentage sets the NullPercentage property
func (m StringProfile) SetNullPercentage(val float32) {
	m.NullPercentage = val
}

// GetTop20 returns the Top20 property
func (m StringProfile) GetTop20() []struct {
	Key   string  `json:"key,omitempty"`
	Value float32 `json:"value,omitempty"`
} {
	return m.Top20
}

// SetTop20 sets the Top20 property
func (m StringProfile) SetTop20(val []struct {
	Key   string  `json:"key,omitempty"`
	Value float32 `json:"value,omitempty"`
}) {
	m.Top20 = val
}

// GetType returns the Type property
func (m StringProfile) GetType() string {
	return m.Type
}

// SetType sets the Type property
func (m StringProfile) SetType(val string) {
	m.Type = val
}

// GetUnique returns the Unique property
func (m StringProfile) GetUnique() float32 {
	return m.Unique
}

// SetUnique sets the Unique property
func (m StringProfile) SetUnique(val float32) {
	m.Unique = val
}
