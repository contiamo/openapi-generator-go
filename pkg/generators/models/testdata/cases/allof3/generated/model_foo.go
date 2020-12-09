// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test
//     Version: 0.1.0
package generatortest

// Foo is an object.
type Foo struct {
	// Mixin:
	Mixin string `json:"mixin,omitempty"`
	// Sub:
	Sub string `json:"sub,omitempty"`
}

// GetMixin returns the Mixin property
func (m Foo) GetMixin() string {
	return m.Mixin
}

// SetMixin sets the Mixin property
func (m Foo) SetMixin(val string) {
	m.Mixin = val
}

// GetSub returns the Sub property
func (m Foo) GetSub() string {
	return m.Sub
}

// SetSub sets the Sub property
func (m Foo) SetSub(val string) {
	m.Sub = val
}
