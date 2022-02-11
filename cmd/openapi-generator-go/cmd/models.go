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
	"context"
	"os"
	"time"

	"github.com/contiamo/openapi-generator-go/pkg/filters"
	"github.com/contiamo/openapi-generator-go/pkg/generators/models"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// modelsCmd represents the models command
var modelsCmd = &cobra.Command{
	Use:   "models",
	Short: "Generate Go code for models",
	Long:  `This generates Go files for each model resolved from the given API spec.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		// human-readable
		log := log.Output(zerolog.ConsoleWriter{
			Out: cmd.OutOrStderr(),
		})

		file, err := cmd.Flags().GetString("spec")
		if err != nil {
			log.Fatal().Err(err).Msg("wrong value for `spec`")
		}
		output, err := cmd.Flags().GetString("output")
		if err != nil {
			log.Fatal().Err(err).Msg("wrong value for `output`")
		}
		packageName, err := cmd.Flags().GetString("package-name")
		if err != nil {
			log.Fatal().Err(err).Msg("wrong value for `package-name`")
		}

		log = log.
			With().
			Str("spec", file).
			Str("output", output).
			Str("package_name", packageName).
			Logger()

		log.Debug().Msg("Loading the spec file...")
		specFile, err := os.Open(file)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to read the spec file")
		}
		log.Debug().Msg("The spec file has been loaded.")

		log.Debug().Msg("Initializing the output directory...")
		err = os.MkdirAll(output, 0755)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to create output folder")
		}
		log.Debug().Msg("Output directory has been initialized.")

		log.Debug().Msg("Generating files...")

		allowedPaths := getAllowedPaths(cmd.Flags())
		bs, err := filters.ByPath(specFile, allowedPaths)
		if err != nil {
			log.Fatal().Err(err).Msg("can't filter by specified paths")
		}

		reader := bytes.NewReader(bs)
		g, err := models.NewGenerator(reader, models.Options{
			PackageName: packageName,
			Destination: output,
			Logger:      log,
		})
		if err != nil {
			log.Fatal().Err(err).Msg("failed to initialize the generator")
		}
		err = g.Generate(ctx)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to generate models")
		}

		log.Debug().Msg("Files have been generated.")
	},
}

//nolint:gochecknoinits  // init is allowed for cobra commands
func init() {
	generateCmd.AddCommand(modelsCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// modelsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// modelsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
