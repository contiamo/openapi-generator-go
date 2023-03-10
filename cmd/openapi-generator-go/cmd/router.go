/*
Copyright © 2020 Tino Rusch

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"bytes"
	"os"
	"path/filepath"

	"github.com/contiamo/openapi-generator-go/v2/pkg/filters"
	"github.com/contiamo/openapi-generator-go/v2/pkg/generators/router"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// routerCmd represents the router command
var routerCmd = &cobra.Command{
	Use:   "router",
	Short: "generate a router",
	Long:  `generate a router.`,
	Run: func(cmd *cobra.Command, args []string) {
		file, _ := cmd.Flags().GetString("spec")
		outputDirectory, _ := cmd.Flags().GetString("output")
		packageName, _ := cmd.Flags().GetString("package-name")
		failNoGroup, _ := cmd.Flags().GetBool("fail-no-group")
		failNoOpID, _ := cmd.Flags().GetBool("fail-no-operation-id")
		specFile, err := os.Open(file)
		if err != nil {
			log.Fatal().Str("spec-file", file).Err(err).Msg("failed to read the spec file")
		}

		allowedPaths := getAllowedPaths(cmd.Flags())
		bs, err := filters.ByPath(specFile, allowedPaths)
		if err != nil {
			log.Fatal().Err(err).Msg("can't filter by specified paths")
		}

		reader := bytes.NewReader(bs)
		err = os.MkdirAll(outputDirectory, 0755)
		if err != nil {
			log.Fatal().Str("output", outputDirectory).Err(err).Msg("failed to create output folder")
		}
		outputFile := filepath.Join(outputDirectory, "router.go")

		output, err := os.Create(outputFile)
		if err != nil {
			log.Fatal().Str("output", outputDirectory).Err(err).Msg("failed to create router.go file")
		}
		defer output.Close()
		err = router.Generate(reader, output, router.Options{
			PackageName:       packageName,
			FailNoGroup:       failNoGroup,
			FailNoOperationID: failNoOpID,
		})
		if err != nil {
			log.Fatal().Err(err).Msg("failed to generate router")
		}
	},
}

//nolint:gochecknoinits  // init is allowed for cobra commands
func init() {
	routerCmd.Flags().Bool("fail-no-group", false, "fail when there is no x-handler-group defined for any of the endpoints")
	routerCmd.Flags().Bool("fail-no-operation-id", false, "fail when there is no operationId defined for any of the methods")
	generateCmd.AddCommand(routerCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// routerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// routerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
