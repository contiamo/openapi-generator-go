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
		if path.Connect != nil {
			operationID := path.Connect.OperationID
			schema := requestBodySchemaFromOperation(path.Connect)
			if operationID != "" && schema != nil && schema.Ref == "" {
				swagger.Components.Schemas[operationID+"Body"] = schema
			}
			path.Connect.Parameters = append(path.Parameters, path.Connect.Parameters...)
			if len(path.Connect.Parameters) > 0 {
				model, err := NewModelFromParameters(path.Connect.Parameters)
				if err != nil {
					continue
				}
				model.Name = tpl.ToPascalCase(path.Connect.OperationID + "QueryParameters")
				model.PackageName = opts.PackageName
				model.SpecTitle = swagger.Info.Title
				model.SpecVersion = swagger.Info.Version
				err = writeModelToFile(model, dst)
				if err != nil {
					continue
				}
			}
		}
		if path.Head != nil {
			operationID := path.Connect.OperationID
			schema := requestBodySchemaFromOperation(path.Head)
			if operationID != "" && schema != nil && schema.Ref == "" {
				swagger.Components.Schemas[operationID+"Body"] = schema
			}
			path.Head.Parameters = append(path.Parameters, path.Head.Parameters...)
			if len(path.Head.Parameters) > 0 {
				model, err := NewModelFromParameters(path.Head.Parameters)
				if err != nil {
					continue
				}
				model.Name = tpl.ToPascalCase(path.Head.OperationID + "QueryParameters")
				model.PackageName = opts.PackageName
				model.SpecTitle = swagger.Info.Title
				model.SpecVersion = swagger.Info.Version
				err = writeModelToFile(model, dst)
				if err != nil {
					continue
				}
			}
		}
		if path.Get != nil {
			operationID := path.Get.OperationID
			schema := requestBodySchemaFromOperation(path.Get)
			if operationID != "" && schema != nil && schema.Ref == "" {
				swagger.Components.Schemas[operationID+"Body"] = schema
			}
			path.Get.Parameters = append(path.Parameters, path.Get.Parameters...)
			if len(path.Get.Parameters) > 0 {
				model, err := NewModelFromParameters(path.Get.Parameters)
				if err != nil {
					continue
				}
				model.Name = tpl.ToPascalCase(path.Get.OperationID + "QueryParameters")
				model.PackageName = opts.PackageName
				model.SpecTitle = swagger.Info.Title
				model.SpecVersion = swagger.Info.Version
				err = writeModelToFile(model, dst)
				if err != nil {
					continue
				}
			}
		}
		if path.Options != nil {
			operationID := path.Options.OperationID
			schema := requestBodySchemaFromOperation(path.Options)
			if operationID != "" && schema != nil && schema.Ref == "" {
				swagger.Components.Schemas[operationID+"Body"] = schema
			}
			path.Options.Parameters = append(path.Parameters, path.Options.Parameters...)
			if len(path.Options.Parameters) > 0 {
				model, err := NewModelFromParameters(path.Options.Parameters)
				if err != nil {
					continue
				}
				model.Name = tpl.ToPascalCase(path.Options.OperationID + "QueryParameters")
				model.PackageName = opts.PackageName
				model.SpecTitle = swagger.Info.Title
				model.SpecVersion = swagger.Info.Version
				err = writeModelToFile(model, dst)
				if err != nil {
					continue
				}
			}
		}
		if path.Delete != nil {
			operationID := path.Delete.OperationID
			schema := requestBodySchemaFromOperation(path.Delete)
			if operationID != "" && schema != nil && schema.Ref == "" {
				swagger.Components.Schemas[operationID+"Body"] = schema
			}
			path.Delete.Parameters = append(path.Parameters, path.Delete.Parameters...)
			if len(path.Delete.Parameters) > 0 {
				model, err := NewModelFromParameters(path.Delete.Parameters)
				if err != nil {
					continue
				}
				model.Name = tpl.ToPascalCase(path.Delete.OperationID + "QueryParameters")
				model.PackageName = opts.PackageName
				model.SpecTitle = swagger.Info.Title
				model.SpecVersion = swagger.Info.Version
				err = writeModelToFile(model, dst)
				if err != nil {
					continue
				}
			}
		}
		if path.Post != nil {
			operationID := path.Post.OperationID
			schema := requestBodySchemaFromOperation(path.Post)
			if operationID != "" && schema != nil && schema.Ref == "" {
				swagger.Components.Schemas[operationID+"Body"] = schema
			}
			path.Post.Parameters = append(path.Parameters, path.Post.Parameters...)
			if len(path.Post.Parameters) > 0 {
				model, err := NewModelFromParameters(path.Post.Parameters)
				if err != nil {
					continue
				}
				model.Name = tpl.ToPascalCase(path.Post.OperationID + "QueryParameters")
				model.PackageName = opts.PackageName
				model.SpecTitle = swagger.Info.Title
				model.SpecVersion = swagger.Info.Version
				err = writeModelToFile(model, dst)
				if err != nil {
					continue
				}
			}
		}
		if path.Put != nil {
			operationID := path.Put.OperationID
			schema := requestBodySchemaFromOperation(path.Put)
			if operationID != "" && schema != nil && schema.Ref == "" {
				swagger.Components.Schemas[operationID+"Body"] = schema
			}
			path.Put.Parameters = append(path.Parameters, path.Put.Parameters...)
			if len(path.Put.Parameters) > 0 {
				model, err := NewModelFromParameters(path.Put.Parameters)
				if err != nil {
					continue
				}
				model.Name = tpl.ToPascalCase(path.Put.OperationID + "QueryParameters")
				model.PackageName = opts.PackageName
				model.SpecTitle = swagger.Info.Title
				model.SpecVersion = swagger.Info.Version
				err = writeModelToFile(model, dst)
				if err != nil {
					continue
				}
			}
		}
		if path.Patch != nil {
			operationID := path.Patch.OperationID
			schema := requestBodySchemaFromOperation(path.Patch)
			if operationID != "" && schema != nil && schema.Ref == "" {
				swagger.Components.Schemas[operationID+"Body"] = schema
			}
			path.Patch.Parameters = append(path.Parameters, path.Patch.Parameters...)
			if len(path.Patch.Parameters) > 0 {
				model, err := NewModelFromParameters(path.Patch.Parameters)
				if err != nil {
					continue
				}
				model.Name = tpl.ToPascalCase(path.Patch.OperationID + "QueryParameters")
				model.PackageName = opts.PackageName
				model.SpecTitle = swagger.Info.Title
				model.SpecVersion = swagger.Info.Version
				err = writeModelToFile(model, dst)
				if err != nil {
					continue
				}
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

		err = writeModelToFile(model, dst)
		if err != nil {
			return err
		}
	}

	return nil

}

func requestBodySchemaFromOperation(op *openapi3.Operation) *openapi3.SchemaRef {
	if op.RequestBody != nil && op.RequestBody.Value != nil && op.RequestBody.Value.Content != nil {
		content, ok := op.RequestBody.Value.Content["application/json"]
		if ok {
			return content.Schema
		}
	}
	return nil
}

func writeModelToFile(model *Model, dst string) error {
	modelName := strings.ToLower(strings.ReplaceAll(fmt.Sprintf("model_%s.go", tpl.ToSnakeCase(model.Name)), " ", "_"))

	buf := &bytes.Buffer{}
	err := model.Render(context.TODO(), buf)
	if err != nil {
		log.Error().Str("name", model.Name).Err(err).Msg("failed to render model")
		return err
	}

	content, err := format.Source(buf.Bytes())
	if err != nil {
		log.Error().Str("name", model.Name).Err(err).Msg("failed to format source code")
		fmt.Println(string(buf.Bytes()))
		return err
	}

	target := filepath.Join(dst, modelName)
	err = ioutil.WriteFile(target, content, 0644)
	if err != nil {
		log.Error().Str("name", model.Name).Str("path", target).Err(err).Msg("failed to write model to file")
		return err
	}
	return nil
}
