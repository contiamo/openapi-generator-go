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

func TestGenerator(t *testing.T) {
	// generate models and enums from hub api
	bs, err := ioutil.ReadFile("./hub.yaml")
	require.NoError(t, err)
	dir, err := ioutil.TempDir(os.TempDir(), "generator-test")
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

	cases := []struct {
		name            string
		description     string
		fileName        string
		expectedContent string
	}{
		{
			name:     "enums work like before (but with code formatting applied :P)",
			fileName: "model_access_request_state.go",
		},
		{
			name:     "array handling without pointers now (children prop is omitempty but no pointer)",
			fileName: "model_resource_dump_response.go",
		},
		{
			name:     "no pointers for maps anymore (properties in this example)",
			fileName: "model_table_update_request.go",
		},
		{
			name:     "allof handling in the technology specific connection properties",
			fileName: "model_postgre_sql_data_source_connection_info.go",
		},
		{
			name:     "deeply nested types like this [][]*interface{} from the pantheon spec",
			fileName: "model_query_response.go",
		},
		{
			name:     "oneof translates to interface{} and embedded anonymous structs are working",
			fileName: "model_search_response.go",
		},
		{
			name:     "additionalProperties are correctly handled and typed",
			fileName: "model_resource_dump_response_body.go",
		},
		{
			name:     "format date-time results in a time.Time and the import is added properly",
			fileName: "model_access_request_response.go",
		},
		{
			name:     "nullable true + usage of allof with one entry results in a pointer to a specific type",
			fileName: "model_upload_response.go",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			bs, err := ioutil.ReadFile(filepath.Join(dir, tc.fileName))
			require.NoError(t, err)
			expected, err := ioutil.ReadFile(filepath.Join("./expected", tc.fileName))
			require.NoError(t, err)
			require.Equal(t, string(expected), string(bs))
		})
	}

}
