package generator_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	generatortest "github.com/contiamo/openapi-generator-go/pkg/generators/models/testdata/cases/object_with_additional_properties/expected"
)

func TestUnmarshalFooWithDefaultValues(t *testing.T) {
	var foo generatortest.Foo
	err := json.Unmarshal([]byte(`{}`), &foo)
	require.NoError(t, err)
	require.EqualValues(t, generatortest.NewFoo(), &foo)
}

func TestUnmarshalFooFooWithDefaultValues(t *testing.T) {
	var fooFoo generatortest.FooFoo
	err := json.Unmarshal([]byte(`{}`), &fooFoo)
	require.NoError(t, err)
	require.EqualValues(t, generatortest.NewFooFoo(), fooFoo)
}

func TestMarshalUnmarshalFoo(t *testing.T) {
	foo := &generatortest.Foo{
		Foo: generatortest.FooFoo{
			"key": "value",
		},
	}
	data, err := json.Marshal(foo)
	require.NoError(t, err)
	require.Equal(t, `{"foo":{"key":"value"}}`, string(data))

	var foo2 generatortest.Foo
	err = json.Unmarshal(data, &foo2)
	require.NoError(t, err)
	require.EqualValues(t, foo, &foo2)
}
