package generator_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	generatortest "github.com/contiamo/openapi-generator-go/pkg/generators/models/testdata/cases/defaults/expected"
)

func TestDefaultValues_Type(t *testing.T) {
	require.Equal(t, generatortest.EnumInteger2, generatortest.DefaultEnumInteger)
	require.Equal(t, generatortest.EnumStringBar, generatortest.DefaultEnumString)
	require.Equal(t, generatortest.String("a string"), generatortest.DefaultString)
	require.Equal(t, generatortest.Email("no@email.com"), generatortest.DefaultEmail)
}

func TestDefaultValues_Factory(t *testing.T) {
	wrapper := generatortest.NewWrapper()
	require.NotNil(t, wrapper)
	require.NotZero(t, *wrapper)

	require.NotZero(t, wrapper.Email)
	require.NotZero(t, wrapper.EnumInteger)
	require.NotZero(t, wrapper.EnumString)
	require.NotZero(t, wrapper.Person)
	require.NotZero(t, wrapper.String)

	require.Equal(t, generatortest.DefaultEmail, wrapper.Email)
	require.Equal(t, generatortest.DefaultEnumInteger, wrapper.EnumInteger)
	require.Equal(t, generatortest.DefaultEnumString, wrapper.EnumString)
	require.Equal(t, generatortest.DefaultString, wrapper.String)

	require.Equal(t, generatortest.NewPerson(), wrapper.Person)
	require.Equal(t, int32(13), wrapper.Person.Age)

	require.Equal(t, generatortest.NewMap(), generatortest.Map(map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}))

	require.Equal(t, generatortest.NewMapWithRef(), generatortest.MapWithRef(map[string]generatortest.Person{
		"key1": {
			Address: generatortest.Address{
				City:   "Springfield",
				Number: 123,
				State:  "IL",
				Street: "123 Main St",
			},
			Age:  25,
			Name: "John",
		},
		"key2": {
			Address: generatortest.Address{
				City:   "Springfield",
				Number: 123,
				State:  "IL",
				Street: "123 Main St",
			},
			Age:  13,
			Name: "Jane",
		},
	}))
}

func TestDefaultValues_UnmarshalJSON(t *testing.T) {
	var wrapper generatortest.Wrapper
	err := json.Unmarshal([]byte(`{}`), &wrapper)
	require.NoError(t, err)
	require.EqualValues(t, generatortest.NewWrapper(), &wrapper)
}

func TestDefaultValuesDoNotOverwrite_UnmarshalJSON(t *testing.T) {
	var wrapper generatortest.Wrapper
	err := json.Unmarshal([]byte(`{"email":"other@email.com"}`), &wrapper)
	require.NoError(t, err)
	require.EqualValues(t, "other@email.com", wrapper.Email)
	wrapper.Email = generatortest.DefaultEmail
	require.EqualValues(t, generatortest.NewWrapper(), &wrapper)
}
