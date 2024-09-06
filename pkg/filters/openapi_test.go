package filters

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFilterByPath(t *testing.T) {

	cases := []struct {
		name    string
		source  string
		filters []string
	}{
		{
			name:   "empty filter list",
			source: "source.yaml",
		},
		{
			name:    "filter with single path",
			source:  "source.yaml",
			filters: []string{"/api/v2/{projectId}/fields"},
		},
		{
			name:    "filter with multiple paths",
			source:  "source.yaml",
			filters: []string{"/api/v2/{projectId}/fields", "/api/v2/{projectId}/field-sets/{resourceKind}"},
		},
		{
			name:    "filter without responses component",
			source:  "source_without_responses_component.yaml",
			filters: []string{"/api/v2/{projectId}/fields"},
		},
		{
			name:    "source with request bodies component",
			source:  "source_with_request_bodies_component.yaml",
			filters: []string{"/api/v2/{projectId}/fields"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			sourceFile, err := os.Open(filepath.Join("testdata", tc.source))
			require.NoError(t, err)

			out, err := ByPath(sourceFile, tc.filters)
			require.NoError(t, err)

			testFolder := strings.ReplaceAll(tc.name, " ", "_")
			exp, err := os.ReadFile(filepath.Join("testdata", testFolder, "out.yaml"))
			require.NoError(t, err)

			require.YAMLEq(t, string(exp), string(out), string(out))
		})
	}
}
