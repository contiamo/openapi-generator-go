// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Hub Service
//     Version: 0.1.0
package testpkg

// GetFoo returns the Foo property
func (m SubType) GetFoo() string {
	return m.Foo
}

// GetBar returns the Bar property
func (m TestType) GetBar() string {
	return m.Bar
}

// GetBaz returns the Baz property
func (m TestType) GetBaz() SubType {
	return m.Baz
}

// GetFoo returns the Foo property
func (m TestType) GetFoo() int {
	return m.Foo
}

// SetFoo sets the Foo property
func (m SubType) SetFoo(val string) {
	m.Foo = val
}

// SetBar sets the Bar property
func (m TestType) SetBar(val string) {
	m.Bar = val
}

// SetBaz sets the Baz property
func (m TestType) SetBaz(val SubType) {
	m.Baz = val
}

// SetFoo sets the Foo property
func (m TestType) SetFoo(val int) {
	m.Foo = val
}
