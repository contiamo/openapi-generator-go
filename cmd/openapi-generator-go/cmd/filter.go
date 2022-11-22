/*
Copyright © 2021 Contiamo

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
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/contiamo/openapi-generator-go/pkg/filters"
)

// filterCmd reads a list of allowed paths and filters the spec so that it
// only includes those paths and the required schemas for the spec to remain
// valid.
var filterCmd = &cobra.Command{
	Use:   "filter",
	Short: "filter spec by a list of paths",
	Run: func(cmd *cobra.Command, args []string) {
		file, err := cmd.Flags().GetString("spec")
		if err != nil {
			log.Fatal().Err(err).Msg("can't read file flag")
		}

		destinationFile, err := cmd.Flags().GetString("output")
		if err != nil {
			log.Fatal().Err(err).Msg("can't read output flag")
		}
		bs, err := os.ReadFile(file)
		if err != nil {
			log.Fatal().Str("spec-file", file).Err(err).Msg("failed to read the spec file")
		}

		allowedPaths := getAllowedPaths(cmd.Flags())
		spec := bytes.NewReader(bs)
		filteredSpec, err := filters.ByPath(spec, allowedPaths)
		if err != nil {
			if zerolog.GlobalLevel() == zerolog.DebugLevel {
				//nolint:forbidigo  // printing in commands is allowed
				fmt.Println(string(filteredSpec))
			}
			log.Fatal().Err(err).Msg("failed to filter the api spec")
		}

		err = os.WriteFile(destinationFile, filteredSpec, 0644)
		if err != nil {
			log.Fatal().Str("output", destinationFile).Err(err).Msg("failed to create output file")
		}
	},
}

//nolint:gochecknoinits  // init is allowed for cobra commands
func init() {
	rootCmd.AddCommand(filterCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// modelsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	filterCmd.Flags().StringP("output", "o", "filtered-spec.yaml", "Output location for the filtered spec")
}
