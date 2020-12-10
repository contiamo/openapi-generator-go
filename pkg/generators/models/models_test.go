package models

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var modelsSpec = `
openapi: 3.0.0
info:
  version: 0.1.0
  title: Hub Service
components:
    schemas:
        TestType:
            type: object
            properties:
              foo:
                type: int
              bar:
                type: string
              baz:
                type: array
                items:
                  $ref: "#/components/schemas/SubType"
        SubType:
            type: object
            properties:
                foo:
                    type: string
`

func TestGenerateModels(t *testing.T) {

	dname, err := ioutil.TempDir("", "modeldir")
	require.NoError(t, err)

	defer os.RemoveAll(dname)

	opts := Options{PackageName: "testpkg"}

	specReader := strings.NewReader(modelsSpec)
	err = Generate(specReader, dname, opts)
	require.NoError(t, err)

	testTypeContent, err := ioutil.ReadFile(filepath.Join(dname, "model_test_type.go"))
	require.NoError(t, err)
	expectedTestType, err := ioutil.ReadFile("testdata/model_test_type.go")
	require.NoError(t, err)
	require.Equal(t, string(expectedTestType), string(testTypeContent))

	subTypeContent, err := ioutil.ReadFile(filepath.Join(dname, "model_sub_type.go"))
	require.NoError(t, err)
	expectedSubType, err := ioutil.ReadFile("testdata/model_sub_type.go")
	require.NoError(t, err)
	require.Equal(t, string(expectedSubType), string(subTypeContent))
}

func TestModels(t *testing.T) {
	cases := []struct {
		name      string
		directory string
	}{
		{
			name:      "simple model",
			directory: "testdata/cases/simple_model",
		},
		{
			name:      "embedded type",
			directory: "testdata/cases/embedded_type",
		},
		{
			name:      "required property removes omitempty",
			directory: "testdata/cases/required_properties",
		},
		{
			name:      "nullable converts to pointer",
			directory: "testdata/cases/nullable_properties",
		},
		{
			name:      "untyped object converts to map[string]interface{}",
			directory: "testdata/cases/untyped_object",
		},
		{
			name:      "oneOf converts to interface{}",
			directory: "testdata/cases/oneof",
		},
		{
			name:      "allOf merges multiple inlined object definitions on property level",
			directory: "testdata/cases/allof1",
		},
		{
			name:      "allOf merges multiple inlined object definitions on top level",
			directory: "testdata/cases/allof2",
		},
		{
			name:      "allOf merges refs and inlined object definitions",
			directory: "testdata/cases/allof3",
		},
		{
			name:      "allOf can be used in conjunction with nullable to produce pointers to other types",
			directory: "testdata/cases/allof4",
		},
		{
			name:      "typed arrays generated typed arrays in go",
			directory: "testdata/cases/typed_arrays",
		},
		{
			name:      "if the array prop is nullable it is NOT converted to a pointer (slices are already nullable in go)",
			directory: "testdata/cases/nullable_arrays",
		},
		{
			name:      "if an untyped object prop is nullable it is NOT converted to a pointer (maps are already nullable in go)",
			directory: "testdata/cases/nullable_untyped_object",
		},
		{
			name:      "an untyped object may have additional properties of a specific type",
			directory: "testdata/cases/object_with_additional_properties",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			dir := filepath.Join(tc.directory, "generated")
			err := os.MkdirAll(dir, 0755)
			require.NoError(t, err)
			bs, err := ioutil.ReadFile(filepath.Join(tc.directory, "api.yaml"))
			require.NoError(t, err)
			reader := bytes.NewReader(bs)
			err = Generate(reader, dir, Options{
				PackageName: "generatortest",
			})
			require.NoError(t, err)
			reader = bytes.NewReader(bs)
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
