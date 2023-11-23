package router

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"net/http"
	"sort"
	"strings"
	"text/template"

	tpl "github.com/contiamo/openapi-generator-go/v2/pkg/generators/templates"
	"github.com/rs/zerolog/log"
	yaml "gopkg.in/yaml.v2"
)

// DefaultPackageName used in the router's source code
const DefaultPackageName = "openapi"

// Options represent all the possible options of the generator
type Options struct {
	// PackageName of the generated router source code (`DefaultPackageName` by default)
	PackageName string

	// FailNoGroup if true the generator returns an error if an endpoint without
	// `x-handler-group` attribute was found. Otherwise, this endpoint will be skipped silently.
	FailNoGroup bool

	// FailNoOperationID if true the generator returns an error if an endpoint without
	// `operationId` attribute was found. Otherwise, this endpoint will be skipped silently.
	FailNoOperationID bool
}

// Generate writes a chi router source code into `router` reading the YAML definition of the
// Open API 3.0 spec from the `specFile`. It supports options `opts` (see `Options`).
//
// All included into generation endpoints must have `operationId` and `x-handler-group`
// attributes. Depending on the `opts` generator will either produce an error or skip
// endpoints without these attributes.
func Generate(specFile io.Reader, router io.Writer, opts Options) (err error) {
	decoder := yaml.NewDecoder(specFile)
	spec := spec{}
	err = decoder.Decode(&spec)
	if err != nil {
		return err
	}

	if opts.PackageName == "" {
		opts.PackageName = DefaultPackageName
	}

	tctx, err := createTemplateCtx(spec, opts)
	if err != nil {
		return err
	}

	buf := &bytes.Buffer{}
	err = routerTemplate.Execute(buf, tctx)
	if err != nil {
		log.Error().Err(err).Str("component", "router").Msg("failed to execute template")
		return err
	}

	content, err := format.Source(buf.Bytes())
	if err != nil {
		log.Error().Str("component", "router").Err(err).Msg("failed to format source code")
		return err
	}

	_, err = io.Copy(router, bytes.NewReader(content))
	if err != nil {
		log.Error().Str("component", "router").Err(err).Msg("failed to write to target output")
		return err
	}
	return nil
}

type templateCtx struct {
	PackageName        string
	Spec               spec
	GroupsSorted       []string
	Groups             map[string]*handlerGroup
	PathsByGroups      map[string]*pathsInGroup
	UniqueOperationIDs map[string]struct{}
}

type pathsInGroup struct {
	PathsSorted           []string
	AllowedMethodsByPaths map[string]*methodsInPath
}

type methodsInPath struct {
	MethodsSorted       []string
	OperationsByMethods map[string]string
}

type handlerGroup struct {
	Endpoints []endpoint
}

type endpoint struct {
	Summary     string `yaml:"summary"`
	Description string `yaml:"description"`
	OperationID string `yaml:"operationId"`
	Group       string `yaml:"x-handler-group"`
}

type path struct {
	GET    *endpoint `yaml:"get"`
	POST   *endpoint `yaml:"post"`
	PUT    *endpoint `yaml:"put"`
	PATCH  *endpoint `yaml:"patch"`
	DELETE *endpoint `yaml:"delete"`
	HEAD   *endpoint `yaml:"head"`
}

type info struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Version     string `yaml:"version"`
}

type spec struct {
	Info  info            `yaml:"info"`
	Paths map[string]path `yaml:"paths"`
}

func createTemplateCtx(spec spec, opts Options) (out templateCtx, err error) {
	out.PackageName = opts.PackageName
	out.Spec = spec
	out.Groups = make(map[string]*handlerGroup)
	out.PathsByGroups = make(map[string]*pathsInGroup)
	out.UniqueOperationIDs = make(map[string]struct{})

	for path, definition := range spec.Paths {
		if definition.GET != nil {
			err = setEndpoint(&out, opts, http.MethodGet, path, *definition.GET)
			if err != nil {
				return out, err
			}
		}

		if definition.POST != nil {
			err = setEndpoint(&out, opts, http.MethodPost, path, *definition.POST)
			if err != nil {
				return out, err
			}
		}

		if definition.PUT != nil {
			err = setEndpoint(&out, opts, http.MethodPut, path, *definition.PUT)
			if err != nil {
				return out, err
			}
		}

		if definition.PATCH != nil {
			err = setEndpoint(&out, opts, http.MethodPatch, path, *definition.PATCH)
			if err != nil {
				return out, err
			}
		}

		if definition.DELETE != nil {
			err = setEndpoint(&out, opts, http.MethodDelete, path, *definition.DELETE)
			if err != nil {
				return out, err
			}
		}

		if definition.HEAD != nil {
			err = setEndpoint(&out, opts, http.MethodHead, path, *definition.HEAD)
			if err != nil {
				return out, err
			}
		}
	}

	// sorting the whole template context to have deterministic results
	sortContext(&out)

	return out, nil
}

func sortContext(out *templateCtx) {
	out.GroupsSorted = make([]string, 0, len(out.Groups))

	for g := range out.Groups {
		out.GroupsSorted = append(out.GroupsSorted, g)
		group := out.Groups[g]

		// sort the list of endpoints by their operation IDs
		sort.SliceStable(group.Endpoints, func(i, j int) bool {
			return group.Endpoints[i].OperationID < group.Endpoints[j].OperationID
		})

		paths := out.PathsByGroups[g]
		// get sorted paths in a group
		paths.PathsSorted = make(
			[]string,
			0,
			len(paths.AllowedMethodsByPaths),
		)
		for path := range paths.AllowedMethodsByPaths {
			paths.PathsSorted = append(paths.PathsSorted, path)

			// get sorted methods in a path
			paths.AllowedMethodsByPaths[path].MethodsSorted = make(
				[]string, 0,
				len(paths.AllowedMethodsByPaths[path].OperationsByMethods),
			)
			for method := range paths.AllowedMethodsByPaths[path].OperationsByMethods {
				paths.AllowedMethodsByPaths[path].MethodsSorted = append(paths.AllowedMethodsByPaths[path].MethodsSorted, method)
			}
			sort.Strings(paths.AllowedMethodsByPaths[path].MethodsSorted)
		}
		sort.Strings(out.PathsByGroups[g].PathsSorted)
	}

	sort.Strings(out.GroupsSorted)
}

func setEndpoint(out *templateCtx, opts Options, method, path string, e endpoint) error {
	if e.Group == "" {
		if opts.FailNoGroup {
			return fmt.Errorf("`%s %s` does not have the `x-handler-group` value", method, path)
		}
		return nil
	}
	if e.OperationID == "" {
		if opts.FailNoOperationID {
			return fmt.Errorf("`%s %s` does not have the `operationId` value", method, path)
		}
		return nil
	}

	casedOperationID := tpl.ToPascalCase(e.OperationID)
	if _, exists := out.UniqueOperationIDs[casedOperationID]; exists {
		return fmt.Errorf(
			"converted operation ID value %q is not unique (original value %q)",
			casedOperationID,
			e.OperationID,
		)
	}
	e.OperationID = casedOperationID
	out.UniqueOperationIDs[e.OperationID] = struct{}{}

	e.Group = tpl.ToPascalCase(e.Group)

	group := out.Groups[e.Group]
	if group == nil {
		group = &handlerGroup{}
		out.Groups[e.Group] = group
	}
	group.Endpoints = append(group.Endpoints, e)

	exPathsInGroup := out.PathsByGroups[e.Group]
	if exPathsInGroup == nil {
		exPathsInGroup = &pathsInGroup{
			AllowedMethodsByPaths: make(map[string]*methodsInPath),
		}
		out.PathsByGroups[e.Group] = exPathsInGroup
	}

	exMethodsInPath := exPathsInGroup.AllowedMethodsByPaths[path]
	if exMethodsInPath == nil {
		exMethodsInPath = &methodsInPath{
			OperationsByMethods: make(map[string]string, 5), // we have only 5 HTTP methods
		}
		exPathsInGroup.AllowedMethodsByPaths[path] = exMethodsInPath
	}

	exMethodsInPath.OperationsByMethods[method] = e.OperationID

	return nil
}

func httpMethod(str string) string {
	return tpl.FirstUpper(strings.ToLower(str))
}

// template definitions
var (
	fmap = template.FuncMap{
		"firstLower":   tpl.FirstLower,
		"firstUpper":   tpl.FirstUpper,
		"httpMethod":   httpMethod,
		"commentBlock": tpl.CommentBlock,
	}
	routerTemplateSource = `package {{ .PackageName }}

// Code generated by openapi-generator-go DO NOT EDIT., don't modify it manually

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

{{ $groups:=.Groups -}}
{{range $groupName := .GroupsSorted }}
{{- $group:=(index $groups $groupName) }}
// {{ $groupName }}Handler handles the operations of the '{{ $groupName }}' handler group.
type {{ $groupName }}Handler interface {
{{- range $e := $group.Endpoints }}
	{{ (printf "%s %s" $e.OperationID $e.Description) | commentBlock }}
	{{ $e.OperationID }}(w http.ResponseWriter, r *http.Request)
{{- end}}
}
{{end}}
// NewRouter creates a new router for the spec and the given handlers.
// {{ .Spec.Info.Title }}
//
{{ .Spec.Info.Description | commentBlock }}
//
// {{ .Spec.Info.Version }}
//
func NewRouter(
{{- $paths := .PathsByGroups -}}
{{- range $groupName := .GroupsSorted }}
{{- $def := (index $paths $groupName) }}
	{{ $groupName | firstLower}}Handler {{ $groupName | firstUpper }}Handler,
{{- end}}
) http.Handler {

	r := chi.NewRouter()
{{range $groupName := .GroupsSorted }}
{{- $def := (index $paths $groupName) }}
// '{{ $groupName }}' group
{{ range $path := $def.PathsSorted }}
{{- $methodsInPath := (index $def.AllowedMethodsByPaths $path) }}
// '{{ $path }}'
r.Options("{{ $path }}", optionsHandlerFunc(
{{- range $method := $methodsInPath.MethodsSorted }}
{{- $operation := (index $methodsInPath.OperationsByMethods $method) }}
	http.Method{{ $method | httpMethod }},
{{- end}}
))

{{- range $method := $methodsInPath.MethodsSorted }}
{{- $operation := (index $methodsInPath.OperationsByMethods $method) }}
r.{{ $method | httpMethod }}("{{ $path }}", {{ $groupName | firstLower }}Handler.{{ $operation }})
{{- end}}
{{end}}
{{- end}}
	return r
}

func optionsHandlerFunc(allowedMethods ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Allow", strings.Join(allowedMethods, ", "))
	}
}
`
	routerTemplate = template.Must(
		template.New("router").
			Funcs(fmap).
			Parse(routerTemplateSource),
	)
)
