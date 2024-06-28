package templates

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMapFunc(t *testing.T) {
	m, err := MapFunc("a", "b")
	require.NoError(t, err)
	require.EqualValues(t, map[string]any{"a": "b"}, m)

	m, err = MapFunc("a", "b", "c", "d")
	require.NoError(t, err)
	require.EqualValues(t, map[string]any{"a": "b", "c": "d"}, m)

	// Test odd number of arguments
	m, err = MapFunc("a", "b", "c")
	require.Error(t, err)
	require.Equal(t, "expecting even number of arguments, but got: 3", err.Error())
	require.Nil(t, m)

	// Test non-string key
	type Person struct {
		Name string
	}
	m, err = MapFunc("a", "b", Person{Name: "Joe"}, "d")
	require.Error(t, err)
	require.Equal(t, "expecting type string as map key, got: templates.Person{Name:\"Joe\"} as map key", err.Error())
	require.Nil(t, m)

	// Test non-string key
	m, err = MapFunc("a", "b", 1, "d")
	require.Error(t, err)
	require.Equal(t, "expecting type string as map key, got: 1 as map key", err.Error())
	require.Nil(t, m)
}
