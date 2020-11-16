// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test Service
//     Version: 0.1.0
package api

// TopLevelAllOf is an object.
type TopLevelAllOf struct {
	// Extra
	Extra string `json:"extra,omitempty"`
	// Foo
	Foo string `json:"foo,omitempty"`
}

// GetExtra returns the Extra property
func (m TopLevelAllOf) GetExtra() string {
	return m.Extra
}

// SetExtra sets the Extra property
func (m TopLevelAllOf) SetExtra(val string) {
	m.Extra = val
}

// GetFoo returns the Foo property
func (m TopLevelAllOf) GetFoo() string {
	return m.Foo
}

// SetFoo sets the Foo property
func (m TopLevelAllOf) SetFoo(val string) {
	m.Foo = val
}
