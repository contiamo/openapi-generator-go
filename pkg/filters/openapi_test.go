package filters

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFilterByPath(t *testing.T) {

	cases := []struct {
		name    string
		filters []string
	}{
		{
			name: "empty filter list",
		},
		{
			name:    "filter with single path",
			filters: []string{"/api/v2/{projectId}/fields"},
		},
		{
			name:    "filter with multiple paths",
			filters: []string{"/api/v2/{projectId}/fields", "/api/v2/{projectId}/field-sets/{resourceKind}"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			sourceFile, err := os.Open(filepath.Join("testdata", "source.yaml"))
			require.NoError(t, err)

			out, err := ByPath(sourceFile, tc.filters)
			require.NoError(t, err)

			testFolder := strings.ReplaceAll(tc.name, " ", "_")
			exp, err := ioutil.ReadFile(filepath.Join("testdata", testFolder, "out.yaml"))
			require.NoError(t, err)

			require.YAMLEq(t, string(exp), string(out))
		})
	}
}
