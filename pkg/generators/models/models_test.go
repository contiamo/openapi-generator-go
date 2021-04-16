package models

import (
	"bytes"
	"context"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
)

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
			name:      "validation",
			directory: "testdata/cases/validation",
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
		{
			name:      "an untyped toplevel object may have additional properties of a specific type",
			directory: "testdata/cases/toplevel_object_with_additional_properties",
		},
		{
			name:      "handles string enums",
			directory: "testdata/cases/enum",
		},
		{
			name:      "enum with one item",
			directory: "testdata/cases/enum_with_one_item",
		},
		{
			name:      "model from path body",
			directory: "testdata/cases/model_from_path_body",
		},
		{
			name:      "enum with name collision",
			directory: "testdata/cases/enum_with_name_collision",
		},
		{
			name:      "model from path parameters",
			directory: "testdata/cases/parameter_model",
		},
		{
			name:      "model with recursive definition within an allof",
			directory: "testdata/cases/allof_self_reference",
		},
		{
			name:      "allof merges required list",
			directory: "testdata/cases/allof_merges_required_list",
		},
		{
			name:      "allof merges enum list",
			directory: "testdata/cases/allof_enum",
		},
		{
			name:      "allof supports intermediate array types",
			directory: "testdata/cases/allof_arrays",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ctx, cancel := context.
				WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			dir := filepath.Join(tc.directory, "generated")
			err := os.MkdirAll(dir, 0755)
			require.NoError(t, err)
			bs, err := ioutil.ReadFile(filepath.Join(tc.directory, "api.yaml"))
			require.NoError(t, err)
			reader := bytes.NewReader(bs)

			g, err := NewGenerator(reader, Options{
				PackageName: "generatortest",
				Destination: dir,
				// Logger:      log.Logger.Output(zerolog.ConsoleWriter{Out: os.Stderr}),
				Logger: log.Output(ioutil.Discard), // swap for debugging
			})
			require.NoError(t, err)

			err = g.Generate(ctx)
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
	t.Run("everything can be built", func(t *testing.T) {
		cmd := exec.Command("go", "build", "./...")
		cmd.Dir = "./testdata"
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		require.NoError(t, cmd.Run())
	})
}

func equalFiles(t *testing.T, expected, actual string) {
	bs1, err := ioutil.ReadFile(expected)
	require.NoError(t, err)
	bs2, err := ioutil.ReadFile(actual)
	require.NoError(t, err)
	require.Equal(t, string(bs1), string(bs2))
}
