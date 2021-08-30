package merge

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// OpenAPISpec is highlevel container for an OpenAPI spec. This can be used for safely
// decode and encode an openapi spec.
type OpenAPISpec struct {
	OpenAPI      string                 `json:"openapi" yaml:"openapi"` // Required
	Info         interface{}            `json:"info" yaml:"info"`       // Required
	Tags         interface{}            `json:"tags,omitempty" yaml:"tags,omitempty"`
	Security     interface{}            `json:"security,omitempty" yaml:"security,omitempty"`
	Servers      interface{}            `json:"servers,omitempty" yaml:"servers,omitempty"`
	ExternalDocs interface{}            `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	Paths        map[string]interface{} `json:"paths" yaml:"paths"` // Required
	Components   components             `json:"components,omitempty" yaml:"components,omitempty"`
}

type components struct {
	Schemas         map[string]interface{} `json:"schemas,omitempty" yaml:"schemas,omitempty"`
	Parameters      map[string]interface{} `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	Headers         map[string]interface{} `json:"headers,omitempty" yaml:"headers,omitempty"`
	RequestBodies   map[string]interface{} `json:"requestBodies,omitempty" yaml:"requestBodies,omitempty"`
	Responses       map[string]interface{} `json:"responses,omitempty" yaml:"responses,omitempty"`
	SecuritySchemes map[string]interface{} `json:"securitySchemes,omitempty" yaml:"securitySchemes,omitempty"`
	Examples        map[string]interface{} `json:"examples,omitempty" yaml:"examples,omitempty"`
	Links           map[string]interface{} `json:"links,omitempty" yaml:"links,omitempty"`
	Callbacks       map[string]interface{} `json:"callbacks,omitempty" yaml:"callbacks,omitempty"`
}

// OpenAPI merges any OpenAPI schema files from path into the spec from
// the supplied file.
//
// The OpenAPI and Info values from the schemas will not be merged into the
// base specFile.
//
// We expect a Reader here because it allows us to parse the spec with minimal
// validation. The github.com/getkin/kin-openapi/openapi3 package enforces
// certain schema correctness that may not be true until after we finish merging
// all of the schemas.
//
// Example:
//
// 		baseSpec, err := os.Open(path)
// 		if err != nil {
// 			log.Fatalf("failed to open the OpenAPI base spec file at `%s`: %s", path, err.Error())
// 		}

// 		spec, err := merge.OpenAPI(baseSpec, "./openapi/parts")
// 		if err != nil {
// 			log.Fatal(err)
// 		}
func OpenAPI(baseSpec io.Reader, path string) (*OpenAPISpec, error) {
	data, err := ioutil.ReadAll(baseSpec)
	if err != nil {
		return nil, fmt.Errorf("can not read spec file: %w", err)
	}

	spec := &OpenAPISpec{}
	err = yaml.Unmarshal(data, spec)
	if err != nil {
		return nil, fmt.Errorf("can not parse the OpenAPI spec: %w", err)
	}

	err = filepath.Walk(path, walkerMergeSpec(spec))

	return spec, err
}

func walkerMergeSpec(dst *OpenAPISpec) filepath.WalkFunc {
	if dst.Paths == nil {
		dst.Paths = map[string]interface{}{}
	}

	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if ext := strings.ToLower(filepath.Ext(path)); ext != ".yaml" && ext != ".yml" {
			// skip because this is not a YAML file
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("can not open spec file: %w", err)
		}

		data, err := ioutil.ReadAll(f)
		if err != nil {
			return fmt.Errorf("can not read spec file: %w", err)
		}

		var spec OpenAPISpec
		err = yaml.Unmarshal(data, &spec)
		if err != nil {
			return fmt.Errorf("can not read parse spec: %w", err)
		}

		err = mergeComponents(&dst.Components, spec)
		if err != nil {
			return fmt.Errorf("can not merge file '%s': %w", info.Name(), err)
		}

		err = mergePaths(dst.Paths, spec)
		if err != nil {
			return fmt.Errorf("can not merge file '%s': %w", info.Name(), err)
		}

		return nil
	}
}

func mergeComponents(dst *components, spec OpenAPISpec) error {
	if spec.Components.Schemas != nil && dst.Schemas == nil {
		dst.Schemas = map[string]interface{}{}
	}

	for name, value := range spec.Components.Schemas {
		_, exists := dst.Schemas[name]
		if exists {
			return fmt.Errorf("duplicate components.schemas: %s", name)
		}
		dst.Schemas[name] = value
	}

	if spec.Components.Parameters != nil && dst.Parameters == nil {
		dst.Parameters = map[string]interface{}{}
	}

	for name, value := range spec.Components.Parameters {
		_, exists := dst.Parameters[name]
		if exists {
			return fmt.Errorf("duplicate components.parameters: %s", name)
		}
		dst.Parameters[name] = value
	}

	if spec.Components.Headers != nil && dst.Headers == nil {
		dst.Headers = map[string]interface{}{}
	}

	for name, value := range spec.Components.Headers {
		_, exists := dst.Headers[name]
		if exists {
			return fmt.Errorf("duplicate components.headers: %s", name)
		}
		dst.Headers[name] = value
	}

	if spec.Components.RequestBodies != nil && dst.RequestBodies == nil {
		dst.RequestBodies = map[string]interface{}{}
	}

	for name, value := range spec.Components.RequestBodies {
		_, exists := dst.RequestBodies[name]
		if exists {
			return fmt.Errorf("duplicate components.requestBodies: %s", name)
		}
		dst.RequestBodies[name] = value
	}

	if spec.Components.Responses != nil && dst.Responses == nil {
		dst.Responses = map[string]interface{}{}
	}

	for name, value := range spec.Components.Responses {
		_, exists := dst.Responses[name]
		if exists {
			return fmt.Errorf("duplicate components.responses: %s", name)
		}
		dst.Responses[name] = value
	}

	if spec.Components.SecuritySchemes != nil && dst.SecuritySchemes == nil {
		dst.SecuritySchemes = map[string]interface{}{}
	}

	for name, value := range spec.Components.SecuritySchemes {
		_, exists := dst.SecuritySchemes[name]
		if exists {
			return fmt.Errorf("duplicate components.securitySchema: %s", name)
		}
		dst.SecuritySchemes[name] = value
	}

	if spec.Components.Examples != nil && dst.Examples == nil {
		dst.Examples = map[string]interface{}{}
	}

	for name, value := range spec.Components.Examples {
		_, exists := dst.Examples[name]
		if exists {
			return fmt.Errorf("duplicate components.examples: %s", name)
		}
		dst.Examples[name] = value
	}

	if spec.Components.Links != nil && dst.Links == nil {
		dst.Links = map[string]interface{}{}
	}

	for name, value := range spec.Components.Links {
		_, exists := dst.Links[name]
		if exists {
			return fmt.Errorf("duplicate components.links: %s", name)
		}
		dst.Links[name] = value
	}

	if spec.Components.Callbacks != nil && dst.Callbacks == nil {
		dst.Callbacks = map[string]interface{}{}
	}

	for name, value := range spec.Components.Callbacks {
		_, exists := dst.Callbacks[name]
		if exists {
			return fmt.Errorf("duplicate components.callbacks: %s", name)
		}
		dst.Callbacks[name] = value
	}

	return nil
}

func mergePaths(dst map[string]interface{}, spec OpenAPISpec) error {
	for name, value := range spec.Paths {
		_, exists := dst[name]
		if exists {
			return fmt.Errorf("duplicate path: %s", name)
		}
		dst[name] = value
	}

	return nil
}
