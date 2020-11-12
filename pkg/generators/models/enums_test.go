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

var spec = `
openapi: 3.0.0
info:
  version: 0.1.0
  title: Hub Service
components:
    schemas:
        FilterType:
            type: string
            enum:
            - showAll
            - hideAll
            - showOnlyExpr
            - hideOnlyExpr
        HighlightIndicatorStart:
            type: string
            enum:
            - '{{{'
`

func TestGenerate(t *testing.T) {

	dname, err := ioutil.TempDir("", "enumedir")
	require.NoError(t, err)

	defer os.RemoveAll(dname)

	opts := Options{PackageName: "testpkg"}

	specReader := strings.NewReader(spec)
	err = GenerateEnums(specReader, dname, opts)
	require.NoError(t, err)

	content, err := ioutil.ReadFile(filepath.Join(dname, "model_filter_type.go"))
	require.NoError(t, err)

	expectedFilterType, err := ioutil.ReadFile("testdata/model_filter_type.go")
	require.NoError(t, err)
	expectedFilterType = bytes.TrimSpace(expectedFilterType)
	content = bytes.TrimSpace(content)
	require.Equal(t, string(expectedFilterType), string(content))

	content, err = ioutil.ReadFile(filepath.Join(dname, "model_highlight_indicator_start.go"))
	require.NoError(t, err)

	expectedHighlight, err := ioutil.ReadFile("testdata/model_highlight_indicator_start.go")
	require.NoError(t, err)
	expectedHighlight = bytes.TrimSpace(expectedHighlight)
	content = bytes.TrimSpace(content)
	require.Equal(t, string(expectedHighlight), string(content))
}
