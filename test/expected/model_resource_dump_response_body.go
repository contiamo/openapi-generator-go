// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// ResourceDumpResponseBody is an object.
type ResourceDumpResponseBody struct {
	// Data:
	Data ResourceDumpResponse `json:"data"`
	// FieldSpecs: Map from resource kind to fieldspec array
	FieldSpecs map[string][]FieldSpec `json:"fieldSpecs"`
}

// GetData returns the Data property
func (m ResourceDumpResponseBody) GetData() ResourceDumpResponse {
	return m.Data
}

// SetData sets the Data property
func (m ResourceDumpResponseBody) SetData(val ResourceDumpResponse) {
	m.Data = val
}

// GetFieldSpecs returns the FieldSpecs property
func (m ResourceDumpResponseBody) GetFieldSpecs() map[string][]FieldSpec {
	return m.FieldSpecs
}

// SetFieldSpecs sets the FieldSpecs property
func (m ResourceDumpResponseBody) SetFieldSpecs(val map[string][]FieldSpec) {
	m.FieldSpecs = val
}
