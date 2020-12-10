// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// Foo is an enum.
type Foo string

var (
	FooBar Foo = "bar"
	FooFoo Foo = "foo"

	// KnownFoo is the list of valid Foo
	KnownFoo = []Foo{
		FooBar,
		FooFoo,
	}
	// KnownFooString is the list of valid Foo as string
	KnownFooString = []string{
		string(FooBar),
		string(FooFoo),
	}

	// InKnownFoo is an ozzo-validator for Foo
	InKnownFoo = validation.In(
		FooBar,
		FooFoo,
	)
)
