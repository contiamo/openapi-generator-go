// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// SchemaPreviewCell is an object. SchemaPreviewCell represents a specific table cell in a file, it shows how that specific value will be parsed with a given schema
type SchemaPreviewCell struct {
	// Error:
	Error string `json:"error"`
	// Name:
	Name string `json:"name"`
	// Value:
	Value string `json:"value"`
}

// GetError returns the Error property
func (m SchemaPreviewCell) GetError() string {
	return m.Error
}

// SetError sets the Error property
func (m SchemaPreviewCell) SetError(val string) {
	m.Error = val
}

// GetName returns the Name property
func (m SchemaPreviewCell) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m SchemaPreviewCell) SetName(val string) {
	m.Name = val
}

// GetValue returns the Value property
func (m SchemaPreviewCell) GetValue() string {
	return m.Value
}

// SetValue sets the Value property
func (m SchemaPreviewCell) SetValue(val string) {
	m.Value = val
}
