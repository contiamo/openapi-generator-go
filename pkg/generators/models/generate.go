package models

import (
	"bytes"
	"context"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"path/filepath"
	"strings"

	tpl "github.com/contiamo/openapi-generator-go/pkg/generators/templates"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// DefaultPackageName used in the models source code
const DefaultPackageName = "openapi"

// Options represent all the possible options of the generator
type Options struct {
	// PackageName of the generated models source code (`DefaultPackageName` by default)
	PackageName string
}

func Generate(specFile io.Reader, dst string, opts Options) error {
	if opts.PackageName == "" {
		opts.PackageName = DefaultPackageName
	}

	data, err := ioutil.ReadAll(specFile)
	if err != nil {
		return errors.Wrap(err, "can not read spec file")
	}
	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(data)
	if err != nil {
		return errors.Wrap(err, "can not parse the OpenAPI spec")
	}

	for name, ref := range swagger.Components.Schemas {
		model, err := NewModelFromRef(ref)
		if err != nil {
			log.Error().Str("name", name).Err(err).Msg("fauled to parse model from openapi spec")
			continue
		}

		model.Name = tpl.ToPascalCase(name)
		model.PackageName = opts.PackageName
		model.SpecTitle = swagger.Info.Title
		model.SpecVersion = swagger.Info.Version

		modelName := strings.ToLower(strings.ReplaceAll(fmt.Sprintf("model_%s.go", tpl.ToSnakeCase(model.Name)), " ", "_"))

		buf := &bytes.Buffer{}
		err = model.Render(context.TODO(), buf)
		if err != nil {
			log.Error().Str("name", name).Err(err).Msg("failed to render model")
			continue
		}

		content, err := format.Source(buf.Bytes())
		if err != nil {
			log.Error().Str("name", name).Err(err).Msg("failed to format source code")
			fmt.Println(string(buf.Bytes()))
			continue
		}

		target := filepath.Join(dst, modelName)
		err = ioutil.WriteFile(target, content, 0644)
		if err != nil {
			log.Error().Str("name", name).Str("path", target).Err(err).Msg("failed to write model to file")
			continue
		}
	}

	return nil

}
