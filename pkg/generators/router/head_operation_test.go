package router

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHeadOperation(t *testing.T) {
	t.Run("generated file as expected", func(t *testing.T) {
		basePath := "testdata/head_operation"
		genDir := filepath.Join(basePath, "generated")
		err := os.RemoveAll(genDir)
		require.NoError(t, err)

		err = os.MkdirAll(genDir, 0755)
		require.NoError(t, err)
		bs, err := ioutil.ReadFile(filepath.Join(basePath, "api.yaml"))
		require.NoError(t, err)
		reader := bytes.NewReader(bs)

		outFile, err := os.Create(filepath.Join(genDir, "router.go"))
		require.NoError(t, err)
		defer outFile.Close()

		writer := bufio.NewWriter(outFile)
		err = Generate(reader, writer, Options{})
		require.NoError(t, err)
		writer.Flush()

		outFile.Close()

		equalFiles(t, filepath.Join(genDir, "router.go"), filepath.Join(basePath, "expected", "router.go"))
	})
}

func equalFiles(t *testing.T, actual, expected string) {
	bs1, err := ioutil.ReadFile(expected)
	require.NoError(t, err)
	bs2, err := ioutil.ReadFile(actual)
	require.NoError(t, err)
	require.Equal(t, string(bs1), string(bs2))
}
