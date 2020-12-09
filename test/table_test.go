package test

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/contiamo/openapi-generator-go/pkg/generators/models"
	"github.com/stretchr/testify/require"
)

func TestTableTest(t *testing.T) {
	cases := []struct {
		name      string
		directory string
	}{
		{
			name:      "simple model",
			directory: "./cases/simple_model",
		},
		{
			name:      "embedded type",
			directory: "./cases/embedded_type",
		},
		{
			name:      "required property removes omitempty",
			directory: "./cases/required_properties",
		},
		{
			name:      "nullable converts to pointer",
			directory: "./cases/nullable_properties",
		},
		{
			name:      "untyped object converts to map[string]interface{}",
			directory: "./cases/untyped_object",
		},
		{
			name:      "oneOf converts to interface{}",
			directory: "./cases/oneof",
		},
		{
			name:      "allOf merges multiple inlined object definitions on property level",
			directory: "./cases/allof1",
		},
		{
			name:      "allOf merges multiple inlined object definitions on top level",
			directory: "./cases/allof2",
		},
		{
			name:      "allOf merges refs and inlined object definitions",
			directory: "./cases/allof3",
		},
		{
			name:      "allOf can be used in conjunction with nullable to produce pointers to other types",
			directory: "./cases/allof4",
		},
		{
			name:      "typed arrays generated typed arrays in go",
			directory: "./cases/typed_arrays",
		},
		{
			name:      "if the array prop is nullable it is NOT converted to a pointer (slices are already nullable in go)",
			directory: "./cases/nullable_arrays",
		},
		{
			name:      "if an untyped object prop is nullable it is NOT converted to a pointer (maps are already nullable in go)",
			directory: "./cases/nullable_untyped_object",
		},
		{
			name:      "an untyped object may have additional properties of a specific type",
			directory: "./cases/object_with_additional_properties",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			dir := filepath.Join(tc.directory, "generated")
			os.MkdirAll(dir, 0755)
			bs, err := ioutil.ReadFile(filepath.Join(tc.directory, "api.yaml"))
			require.NoError(t, err)
			reader := bytes.NewReader(bs)
			err = models.GenerateModels(reader, dir, models.Options{
				PackageName: "generatortest",
			})
			require.NoError(t, err)
			reader = bytes.NewReader(bs)
			err = models.GenerateEnums(reader, dir, models.Options{
				PackageName: "generatortest",
			})
			require.NoError(t, err)
			files, err := filepath.Glob(filepath.Join(dir, "*"))
			require.NoError(t, err)
			for _, f := range files {
				equalFiles(t,
					filepath.Join(tc.directory, "expected", filepath.Base(f)),
					filepath.Join(tc.directory, "generated", filepath.Base(f)),
				)
			}
		})
	}
}

func equalFiles(t *testing.T, expected, actual string) {
	bs1, err := ioutil.ReadFile(expected)
	require.NoError(t, err)
	bs2, err := ioutil.ReadFile(actual)
	require.NoError(t, err)
	require.Equal(t, string(bs1), string(bs2))
}
