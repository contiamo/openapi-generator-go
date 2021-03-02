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

	if swagger.Components.Schemas == nil {
		swagger.Components.Schemas = make(map[string]*openapi3.SchemaRef)
	}

	for _, path := range swagger.Paths {
		if path.Post != nil {
			operationID := path.Post.OperationID
			schema := schemaFromOperation(path.Post)
			if operationID != "" && schema != nil && schema.Ref == "" {
				swagger.Components.Schemas[operationID+"Body"] = schema
			}
		}
		if path.Put != nil {
			operationID := path.Put.OperationID
			schema := schemaFromOperation(path.Put)
			if operationID != "" && schema != nil && schema.Ref == "" {
				swagger.Components.Schemas[operationID+"Body"] = schema
			}
		}
		if path.Patch != nil {
			operationID := path.Patch.OperationID
			schema := schemaFromOperation(path.Patch)
			if operationID != "" && schema != nil && schema.Ref == "" {
				swagger.Components.Schemas[operationID+"Body"] = schema
			}
		}
	}

	for name, ref := range swagger.Components.Schemas {
		if ref.Value.Type == "array" {
			// We dont generate toplevel arrays. If they are referenced within another object, they will translate to []ItemType
			continue
		}
		model, err := NewModelFromRef(ref)
		if err != nil {
			log.Debug().Str("name", name).Err(err).Msg("failed to parse model from openapi spec")
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

func schemaFromOperation(op *openapi3.Operation) *openapi3.SchemaRef {
	content, ok := op.RequestBody.Value.Content["application/json"]
	if ok {
		return content.Schema
	}
	return nil
}
