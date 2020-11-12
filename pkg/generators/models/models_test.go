package models

import (
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
	err = GenerateModels(specReader, dname, opts)
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
