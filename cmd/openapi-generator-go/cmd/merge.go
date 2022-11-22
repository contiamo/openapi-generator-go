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
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/contiamo/openapi-generator-go/pkg/merge"
)

// mergeCmd merges a directory of openapi specs into a base spec.
var mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "merges a directory of openapi specs into a base spec",
	Run: func(cmd *cobra.Command, args []string) {
		baseFile, err := cmd.Flags().GetString("base")
		if err != nil {
			log.Fatal().Err(err).Msg("can't read base flag")
		}

		dir, err := cmd.Flags().GetString("dir")
		if err != nil {
			log.Fatal().Err(err).Msg("can't read dir flag")
		}

		destinationFile, err := cmd.Flags().GetString("merged-spec")
		if err != nil {
			log.Fatal().Err(err).Msg("can't read merged-spec flag")
		}

		f, err := os.Open(baseFile)
		if err != nil {
			log.Fatal().Err(err).Msg("can't read base spec")
		}

		mergedSpec, err := merge.OpenAPI(f, dir)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to merge the api spec")
		}

		bs, err := yaml.Marshal(mergedSpec)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to encode spec")
		}

		err = os.WriteFile(destinationFile, bs, 0644)
		if err != nil {
			log.Fatal().Str("output", destinationFile).Err(err).Msg("failed to create output file")
		}
	},
}

//nolint:gochecknoinits  // init is allowed for cobra commands
func init() {
	rootCmd.AddCommand(mergeCmd)

	mergeCmd.Flags().StringP("base", "b", "base.yaml", "base spec")
	mergeCmd.Flags().StringP("dir", "d", "./parts", "part directory")
	mergeCmd.Flags().StringP("merged-spec", "m", "api.yaml", "output spec file")
}
