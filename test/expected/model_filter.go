// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package generatortest

// Filter is an object. A filter to reduce the output of list requests
type Filter struct {
	// Operation:
	Operation OperationType `json:"operation"`
	// Property:
	Property string `json:"property"`
	// Value:
	Value interface{} `json:"value"`
}

// GetOperation returns the Operation property
func (m Filter) GetOperation() OperationType {
	return m.Operation
}

// SetOperation sets the Operation property
func (m Filter) SetOperation(val OperationType) {
	m.Operation = val
}

// GetProperty returns the Property property
func (m Filter) GetProperty() string {
	return m.Property
}

// SetProperty sets the Property property
func (m Filter) SetProperty(val string) {
	m.Property = val
}

// GetValue returns the Value property
func (m Filter) GetValue() interface{} {
	return m.Value
}

// SetValue sets the Value property
func (m Filter) SetValue(val interface{}) {
	m.Value = val
}
