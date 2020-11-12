package templates

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFirstLower(t *testing.T) {
	require.Equal(t, "string", FirstLower("string"))
	require.Equal(t, "sTRING", FirstLower("STRING"))
	require.Equal(t, "string", FirstLower("String"))
}

func TestFirstUpper(t *testing.T) {
	require.Equal(t, "String", FirstUpper("string"))
	require.Equal(t, "STRING", FirstUpper("STRING"))
	require.Equal(t, "String", FirstUpper("String"))
}

func TestToSnakeCase(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "handle empty string",
			input:  "",
			output: "",
		},
		{
			name:   "noop if string is already snakey",
			input:  "already_snake",
			output: "already_snake",
		},
		{
			name:   "consecutive uppercase is converted to lower",
			input:  "AA",
			output: "aa",
		},
		{
			name:   "insert underscore only between last two consecutive upper case",
			input:  "HTTPRequest",
			output: "http_request",
		},
		{
			name:   "pascal case converts",
			input:  "BatteryLifeValue",
			output: "battery_life_value",
		},
		{
			name:   "insert underscore after number",
			input:  "Id0Value",
			output: "id0_value",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := ToSnakeCase(tc.input)
			require.Equal(t, tc.output, got)
		})
	}
}

func TestToPascalCase(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "Converts underscore casing to pascal casing",
			input:  "some_table_name",
			output: "SomeTableName",
		},
		{
			name:   "Converts mixed casing",
			input:  "some_table_nameSomeName",
			output: "SomeTableNameSomeName",
		},
		{
			name:   "Does not change the value if it's pascal casing",
			input:  "SomeName",
			output: "SomeName",
		},
		{
			name:   "Trims single bad character at the end",
			input:  "Some{",
			output: "Some",
		},
		{
			name:   "Trims multiple bad characters at the end",
			input:  "Some{{{",
			output: "Some",
		},
		{
			name: "Handles empty string",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.output, ToPascalCase(tc.input))
		})
	}
}
