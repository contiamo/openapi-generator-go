package generator_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	generatortest "github.com/contiamo/openapi-generator-go/pkg/generators/models/testdata/cases/validation/generated"
)

func TestMapWithValue(t *testing.T) {
	m := generatortest.NewMapWithValue()
	err := m.Validate()
	require.NoError(t, err)

	// Set a valid value
	m["key_valid"] = nil
	err = m.Validate()
	require.NoError(t, err)

	// Set another valid value
	m["key_valid"] = new(string)
	*m["key_valid"] = "value"
	err = m.Validate()
	require.NoError(t, err)

	// Set an invalid value
	m["key_invalid"] = new(string)
	*m["key_invalid"] = ""
	err = m.Validate()
	require.Error(t, err)
	require.Equal(t, "key_invalid: cannot be blank.", err.Error())
}

func TestUnmarshalMapWithValue(t *testing.T) {
	m := generatortest.NewMapWithValue()
	err := m.UnmarshalJSON([]byte(`{"key_valid": "value", "key_invalid": "", "key_default": null}`))
	require.NoError(t, err)
	require.Len(t, m, 3)
	require.NotNil(t, m["key_valid"])
	require.Equal(t, "value", *m["key_valid"])
	require.NotNil(t, m["key_invalid"])
	require.Equal(t, "", *m["key_invalid"])
	require.NotNil(t, m["key_default"])
	require.Equal(t, "value", *m["key_default"])

	err = m.Validate()
	require.Error(t, err)
	require.Equal(t, "key_invalid: cannot be blank.", err.Error())
}

func TestUnmarshalObjectWithAdditionalStrings(t *testing.T) {
	m := generatortest.NewObjectWithAdditionalStrings()
	err := m.UnmarshalJSON([]byte(`{"field1": "value", "key1": "value", "key2": null}`))
	require.NoError(t, err)
	require.Equal(t, "value", m.Field1)
	require.Len(t, m.AdditionalProperties, 2)
	fmt.Printf("%+v\n", m.AdditionalProperties)
	require.NotNil(t, m.AdditionalProperties["key1"])
	require.Equal(t, "value", *m.AdditionalProperties["key1"])
	require.NotNil(t, m.AdditionalProperties["key2"])
	require.Equal(t, "value", *m.AdditionalProperties["key2"])
}
