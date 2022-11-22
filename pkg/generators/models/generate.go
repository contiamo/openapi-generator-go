package models

import (
	"bytes"
	"context"
	"fmt"
	"go/format"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	tpl "github.com/contiamo/openapi-generator-go/pkg/generators/templates"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// DefaultPackageName used in the models source code
const DefaultPackageName = "openapi"

// Options represent all the possible options of the generator
type Options struct {
	// PackageName of the generated models source code (`DefaultPackageName` by default)
	PackageName string
	// Destination is a path to a folder where create all the Go files
	Destination string
	// Logger instance to use for messaging
	Logger zerolog.Logger
}

// NewGenerator creates a standard generator implementation
func NewGenerator(specFile io.Reader, opts Options) (Generator, error) {
	log := opts.Logger.
		With().
		Str("dst", opts.Destination).
		Str("pkg", opts.PackageName).
		Logger()

	if opts.PackageName == "" {
		opts.PackageName = DefaultPackageName
		log = log.
			With().
			Str("dst", opts.PackageName).
			Logger()

		log.Warn().Msg("The package name is not set, default name is set")
	}

	log.Debug().Msg("Reading the spec file...")
	data, err := io.ReadAll(specFile)
	if err != nil {
		return nil, errors.Wrap(err, "can not read spec file")
	}
	log.Debug().Msg("File has been read.")

	log.Debug().Msg("Parsing the spec...")
	swagger, err := openapi3.NewLoader().LoadFromData(data)
	if err != nil {
		return nil, errors.Wrap(err, "can not parse the OpenAPI spec")
	}
	log.Debug().Msg("Spec has been parsed.")

	return generator{
		spec: swagger,
		opts: opts,
	}, nil
}

// Generator defines generator operations needed for generating Go code for the spec
type Generator interface {
	// Generate generates and writes model files that can be generated out of the given spec.
	Generate(ctx context.Context) error
}

// generator implements the whole generation process based on input parameters
type generator struct {
	// spec is the initial Open API spec read from a file
	spec *openapi3.T
	// opts are various options set for the generation process
	opts Options
}

// Generate generates and writes model files that can be generated out of the given spec.
func (g generator) Generate(ctx context.Context) (err error) {
	if g.spec.Components.Schemas == nil {
		g.spec.Components.Schemas = make(map[string]*openapi3.SchemaRef)
	}

	log := g.opts.Logger.
		With().
		Int("path_len", len(g.spec.Paths)).
		Logger()

	log.Debug().Msg("Processing paths...")
	for path, pathItem := range g.spec.Paths {
		log := log.
			With().
			Str("path", path).
			Logger()

		log.Debug().Msg("Processing path...")
		operations := map[string]*openapi3.Operation{
			http.MethodConnect: pathItem.Connect,
			http.MethodHead:    pathItem.Head,
			http.MethodGet:     pathItem.Get,
			http.MethodOptions: pathItem.Options,
			http.MethodDelete:  pathItem.Delete,
			http.MethodPost:    pathItem.Post,
			http.MethodPut:     pathItem.Put,
			http.MethodPatch:   pathItem.Patch,
		}

		for name, operation := range operations {
			err = ctx.Err()
			if err != nil {
				return err
			}
			if operation == nil {
				continue
			}
			log := log.
				With().
				Str("operation", name).
				Logger()

			log.Debug().Msg("Processing operation...")

			err = g.processOperation(ctx, pathItem, operation)
			if err != nil {
				return errors.Wrapf(err, "cannot process operation %q on path %q", name, path)
			}

			log.Debug().Msg("Operation has been processed.")
		}
		log.Debug().Msg("Path has been processed.")
	}
	log.Debug().Msg("Paths have been processed.")

	log.Debug().Msg("Processing schemas...")
	for _, name := range sortedKeys(g.spec.Components.Schemas) {
		ref := g.spec.Components.Schemas[name]
		err = ctx.Err()
		if err != nil {
			return err
		}
		err = g.processSchema(ctx, name, ref)
		if err != nil {
			return errors.Wrapf(err, "cannot process schema %q", name)
		}
	}
	log.Debug().Msg("Schemas have been processed.")

	return nil
}

// setSchema sets the given schema checking for name collisions
func (g generator) setSchema(name string, s *openapi3.SchemaRef) (err error) {
	if _, exists := g.spec.Components.Schemas[name]; exists {
		return fmt.Errorf(
			"name collision found: tried to add new schema %s that already exists",
			name,
		)
	}

	g.spec.Components.Schemas[name] = s

	return nil
}

// processParameters generates and writes to a file a model based on the given set of parameters.
// The model struct will contain all the operation parameters and validation for them.
func (g generator) processParameters(ctx context.Context, o *openapi3.Operation, p openapi3.Parameters) (err error) {
	if len(p) == 0 {
		return nil
	}

	model, err := NewModelFromParameters(p)
	if err != nil {
		return err
	}

	model.Name = tpl.ToPascalCase(o.OperationID + "QueryParameters")
	model.PackageName = g.opts.PackageName
	model.SpecTitle = g.spec.Info.Title
	model.SpecVersion = g.spec.Info.Version

	return g.writeModelToFile(ctx, model, g.opts.Destination)
}

// processOperation generates and writes to files all the models that can be extracted
// from the operation definitions.
// This includes: request body payload and query parameters for now.
func (g generator) processOperation(ctx context.Context, p *openapi3.PathItem, o *openapi3.Operation) (err error) {
	if o == nil {
		return nil
	}

	var reqSchema *openapi3.SchemaRef

	// we try to find a request body definition for a JSON payload
	if o.RequestBody != nil &&
		o.RequestBody.Value != nil &&
		o.RequestBody.Value.Content != nil {

		content, ok := o.RequestBody.Value.Content["application/json"]
		if ok {
			reqSchema = content.Schema
		}
	}

	// if the request body is not a reference to an existing schema,
	// we generate one with the operation-specific name
	if o.OperationID != "" && reqSchema != nil && reqSchema.Ref == "" {
		reqSchemaName := tpl.ToPascalCase(o.OperationID + "Body")
		err = g.setSchema(reqSchemaName, reqSchema)
		if err != nil {
			return err
		}
	}

	// append operation specific parameters to base path parameters
	o.Parameters = append(p.Parameters, o.Parameters...)
	err = g.processParameters(ctx, o, o.Parameters)
	if err != nil {
		return err
	}

	return nil
}

// processSchema generates and writes to a file a model out of an Open API schema definition
func (g generator) processSchema(ctx context.Context, name string, s *openapi3.SchemaRef) (err error) {
	if s.Value.Type == "array" {
		// We dont generate toplevel arrays.
		// If they are referenced within another object, they will translate to []ItemType
		return nil
	}
	model, err := NewModelFromRef(s)
	if err != nil {
		return err
	}

	model.Name = tpl.ToPascalCase(name)
	model.PackageName = g.opts.PackageName
	model.SpecTitle = g.spec.Info.Title
	model.SpecVersion = g.spec.Info.Version

	err = g.writeModelToFile(ctx, model, g.opts.Destination)
	if err != nil {
		return err
	}

	return err
}

// writeModelToFile renders, formats and writes Go code to a file for a given model in the `dst` folder.
func (g generator) writeModelToFile(ctx context.Context, model *Model, dst string) error {
	if model == nil {
		return nil
	}

	filename := fmt.Sprintf("model_%s.go", tpl.ToSnakeCase(model.Name))
	filename = strings.ReplaceAll(filename, " ", "_")
	filename = strings.ToLower(filename)

	log := g.opts.Logger.
		With().
		Str("filename", filename).
		Str("model_name", model.Name).
		Logger()

	log.Debug().Msg("Rendering the model template...")
	buf := &bytes.Buffer{}
	err := model.Render(ctx, buf)
	if err != nil {
		return errors.Wrap(err, "cannot render model template")
	}
	log.Debug().Msg("Model template has been rendered.")

	log.Debug().Msg("Formatting the rendered code...")
	content, err := format.Source(buf.Bytes())
	if err != nil {
		return errors.Wrap(err, "cannot format model code")
	}
	log.Debug().Msg("Code has been formatted.")

	log.Debug().Msg("Writing the model to file...")
	target := filepath.Join(dst, filename)
	err = os.WriteFile(target, content, 0644)
	if err != nil {
		return errors.Wrap(err, "cannot write model file")
	}
	log.Debug().Msg("File has been written.")

	return nil
}
