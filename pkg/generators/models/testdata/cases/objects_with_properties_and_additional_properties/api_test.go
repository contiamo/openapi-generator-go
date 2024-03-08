package generator_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	generatortest "github.com/contiamo/openapi-generator-go/pkg/generators/models/testdata/cases/objects_with_properties_and_additional_properties/generated"
)

func TestUnmarshalFooWithDefaultValues(t *testing.T) {
	var foo generatortest.Foo
	err := json.Unmarshal([]byte(`{}`), &foo)
	require.NoError(t, err)
	require.EqualValues(t, generatortest.NewFoo(), &foo)
}

func TestUnmarshalFooBarWithDefaultValues(t *testing.T) {
	var fooBar generatortest.FooBar
	err := json.Unmarshal([]byte(`{}`), &fooBar)
	require.NoError(t, err)
	require.EqualValues(t, generatortest.NewFooBar(), &fooBar)
}

func TestMarshalUnmarshalFoo(t *testing.T) {
	foo := &generatortest.Foo{
		Bar: 1,
		Baz: 2.0,
		AdditionalProperties: map[string]interface{}{
			"key": "value",
		},
	}
	data, err := json.Marshal(foo)
	require.NoError(t, err)
	require.Equal(t, `{"bar":1,"baz":2,"key":"value"}`, string(data))

	var foo2 generatortest.Foo
	err = json.Unmarshal(data, &foo2)
	require.NoError(t, err)
	require.EqualValues(t, foo, &foo2)

	// Test that Baz is omitted when it's 0 and there are no more properties when additional properties is not nil
	foo3 := generatortest.NewFoo()
	foo3.Baz = 0
	foo3.AdditionalProperties = map[string]interface{}{}
	data, err = json.Marshal(foo3)
	require.NoError(t, err)
	require.Equal(t, `{"bar":0}`, string(data))
}

func TestMarshalUnmarshalFooBar(t *testing.T) {
	fooBar := &generatortest.FooBar{
		Baz: "boo",
		AdditionalProperties: map[string][]generatortest.FooBarAdditionalProperties{
			"hi": {{Test: "who?"}, {Test: "me?"}},
		},
	}
	data, err := json.Marshal(fooBar)
	require.NoError(t, err)
	require.Equal(t, `{"baz":"boo","hi":[{"test":"who?"},{"test":"me?"}]}`, string(data))

	var fooBar2 generatortest.FooBar
	err = json.Unmarshal(data, &fooBar2)
	require.NoError(t, err)
	require.EqualValues(t, fooBar, &fooBar2)

	// Test that Baz is omitted when it's empty and there is one item in the additional properties
	fooBar3 := generatortest.NewFooBar()
	fooBar3.Baz = ""
	fooBar3.AdditionalProperties = map[string][]generatortest.FooBarAdditionalProperties{
		"hi": {},
	}
	data, err = json.Marshal(fooBar3)
	require.NoError(t, err)
	require.Equal(t, `{"hi":[]}`, string(data))
}

func TestMarshallFooFailsOnDuplicateKeys(t *testing.T) {
	foo := &generatortest.Foo{
		Bar: 1,
		Baz: 2.0,
		AdditionalProperties: map[string]interface{}{
			"bar": "value",
		},
	}
	_, err := json.Marshal(foo)
	require.Error(t, err)
	var jsonError *json.MarshalerError
	require.ErrorAs(t, err, &jsonError)
	require.Equal(t, reflect.TypeOf(foo), jsonError.Type)
	require.Contains(t, jsonError.Unwrap().Error(), " bar ")
}

func TestMarshallFooBarFailsOnDuplicateKeys(t *testing.T) {
	fooBar := &generatortest.FooBar{
		Baz: "",
		AdditionalProperties: map[string][]generatortest.FooBarAdditionalProperties{
			"baz": {{Test: "who?"}, {Test: "me?"}},
		},
	}
	_, err := json.Marshal(fooBar)
	require.Error(t, err)
	var jsonError *json.MarshalerError
	require.ErrorAs(t, err, &jsonError)
	require.Equal(t, reflect.TypeOf(fooBar), jsonError.Type)
	require.Contains(t, jsonError.Unwrap().Error(), " baz ")
}
