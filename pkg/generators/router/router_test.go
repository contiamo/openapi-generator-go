package router

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
)

func TestHttpMethod(t *testing.T) {
	require.Equal(t, "Get", httpMethod("GET"))
	require.Equal(t, "Get", httpMethod("get"))
	require.Equal(t, "Get", httpMethod("gEt"))
}

func TestSetEndpoint(t *testing.T) {
	valid := endpoint{
		OperationID: "TestOperation",
		Group:       "TestGroup",
	}

	notPascal := endpoint{
		OperationID: "test_operation",
		Group:       "test_group",
	}

	path := "/some/path"
	method := http.MethodGet
	cases := []struct {
		name   string
		e      endpoint
		opts   Options
		exp    templateCtx
		expErr string
	}{
		{
			name: "returns populated dictionaries for a valid endpoint",
			e:    valid,
			exp: templateCtx{
				UniqueOperationIDs: map[string]struct{}{
					valid.OperationID: {},
				},
				Groups: map[string]*handlerGroup{
					valid.Group: {Endpoints: []endpoint{valid}},
				},
				PathsByGroups: map[string]*pathsInGroup{
					valid.Group: {
						AllowedMethodsByPaths: map[string]*methodsInPath{
							path: {
								OperationsByMethods: map[string]string{
									method: valid.OperationID,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "returns populated dictionaries for a valid endpoint with a not pascalifed name",
			e:    notPascal,
			exp: templateCtx{
				// the endpoint should be successfully converted into the `valid`
				UniqueOperationIDs: map[string]struct{}{
					valid.OperationID: {},
				},
				Groups: map[string]*handlerGroup{
					valid.Group: {Endpoints: []endpoint{valid}},
				},
				PathsByGroups: map[string]*pathsInGroup{
					valid.Group: {
						AllowedMethodsByPaths: map[string]*methodsInPath{
							path: {
								OperationsByMethods: map[string]string{
									method: valid.OperationID,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "does not populate if the endpoint is nil",
			exp: templateCtx{
				UniqueOperationIDs: make(map[string]struct{}),
				Groups:             make(map[string]*handlerGroup),
				PathsByGroups:      make(map[string]*pathsInGroup),
			},
		},
		{
			name: "does not populate if the endpoint has no group",
			e:    endpoint{OperationID: "some"},
			exp: templateCtx{
				UniqueOperationIDs: make(map[string]struct{}),
				Groups:             make(map[string]*handlerGroup),
				PathsByGroups:      make(map[string]*pathsInGroup),
			},
		},
		{
			name: "does not populate if the endpoint has no operation ID",
			e:    endpoint{Group: "some"},
			exp: templateCtx{
				UniqueOperationIDs: make(map[string]struct{}),
				Groups:             make(map[string]*handlerGroup),
				PathsByGroups:      make(map[string]*pathsInGroup),
			},
		},
		{
			name:   "returns error if opts.FailNoGroup = true and the endpoint has no group",
			e:      endpoint{OperationID: "some"},
			opts:   Options{FailNoGroup: true},
			expErr: fmt.Sprintf("`%s %s` does not have the `x-handler-group` value", method, path),
		},
		{
			name:   "returns error if opts.FailNoOperationID = true and the endpoint has no operation ID",
			e:      endpoint{Group: "some"},
			opts:   Options{FailNoOperationID: true},
			expErr: fmt.Sprintf("`%s %s` does not have the `operationId` value", method, path),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			out := &templateCtx{
				UniqueOperationIDs: make(map[string]struct{}),
				Groups:             make(map[string]*handlerGroup),
				PathsByGroups:      make(map[string]*pathsInGroup),
			}

			err := setEndpoint(out, tc.opts, method, path, tc.e)
			if tc.expErr != "" {
				require.Error(t, err)
				require.Equal(t, tc.expErr, err.Error())
				return
			}

			require.EqualValues(t, tc.exp, *out)
		})
	}

	t.Run("returns error if pascalified OperationID is not unique", func(t *testing.T) {
		out := &templateCtx{
			UniqueOperationIDs: map[string]struct{}{
				valid.OperationID: {},
			},
			Groups:        make(map[string]*handlerGroup),
			PathsByGroups: make(map[string]*pathsInGroup),
		}
		err := setEndpoint(out, Options{}, method, path, notPascal)
		require.Error(t, err)
		require.Equal(t, "converted operation ID value \"TestOperation\" is not unique (original value \"test_operation\")", err.Error())
	})
}

func TestCreateTemplateCtx(t *testing.T) {
	spec := spec{
		Info: info{
			Title:       "Title",
			Description: "Description",
			Version:     "Version",
		},
		Paths: map[string]path{
			"/some/path": {
				GET: &endpoint{
					Summary:     "GET Summary",
					Description: "GET Description",
					OperationID: "GETOperationId",
					Group:       "Group",
				},
				POST: &endpoint{
					Summary:     "POST Summary",
					Description: "POST Description",
					OperationID: "POST_operation_id",
					Group:       "_group",
				},
				PUT: &endpoint{
					Summary:     "PUT Summary",
					Description: "PUT Description",
					OperationID: "PUTOperationId",
					Group:       "Group",
				},
				PATCH: &endpoint{
					Summary:     "PATCH Summary",
					Description: "PATCH Description",
					OperationID: "PATCHOperationId",
					Group:       "Group",
				},
				DELETE: &endpoint{
					Summary:     "DELETE Summary",
					Description: "DELETE Description",
					OperationID: "DELETEOperationId",
					Group:       "Group",
				},
			},
			"/another/path": {},
		},
	}

	out, err := createTemplateCtx(spec, Options{PackageName: "Test"})
	require.NoError(t, err)
	expected := templateCtx{
		PackageName: "Test",
		Spec:        spec,
		GroupsSorted: []string{
			"Group",
		},
		UniqueOperationIDs: map[string]struct{}{
			"GETOperationId":    {},
			"POSTOperationId":   {},
			"PUTOperationId":    {},
			"PATCHOperationId":  {},
			"DELETEOperationId": {},
		},
		Groups: map[string]*handlerGroup{
			"Group": {
				Endpoints: []endpoint{
					*spec.Paths["/some/path"].DELETE,
					*spec.Paths["/some/path"].GET,
					*spec.Paths["/some/path"].PATCH,
					{
						Summary:     "POST Summary",
						Description: "POST Description",
						OperationID: "POSTOperationId",
						Group:       "Group",
					},
					*spec.Paths["/some/path"].PUT,
				},
			},
		},
		PathsByGroups: map[string]*pathsInGroup{
			"Group": {
				PathsSorted: []string{
					"/some/path",
				},
				AllowedMethodsByPaths: map[string]*methodsInPath{
					"/some/path": {
						MethodsSorted: []string{
							"DELETE",
							"GET",
							"PATCH",
							"POST",
							"PUT",
						},
						OperationsByMethods: map[string]string{
							"GET":    "GETOperationId",
							"POST":   "POSTOperationId",
							"PUT":    "PUTOperationId",
							"PATCH":  "PATCHOperationId",
							"DELETE": "DELETEOperationId",
						},
					},
				},
			},
		},
	}
	require.EqualValues(t, expected, out)
}

func TestGenerate(t *testing.T) {
	spec := spec{
		Info: info{
			Title:       "Title",
			Description: "Description",
			Version:     "Version",
		},
		Paths: map[string]path{
			"/some/path": {
				GET: &endpoint{
					Summary:     "GET Summary",
					Description: "GET Description",
					OperationID: "GETOperationID",
					Group:       "Group",
				},
				POST: &endpoint{
					Summary:     "POST Summary",
					Description: "POST Description",
					OperationID: "POSTOperationID",
					Group:       "Group",
				},
				PUT: &endpoint{
					Summary:     "PUT Summary",
					Description: "PUT Description",
					OperationID: "PUTOperationID",
					Group:       "Group",
				},
				PATCH: &endpoint{
					Summary:     "PATCH Summary",
					Description: "PATCH Description",
					OperationID: "PATCHOperationID",
					Group:       "Group",
				},
				DELETE: &endpoint{
					Summary:     "DELETE Summary",
					Description: "DELETE Description",
					OperationID: "DELETEOperationID",
					Group:       "Group",
				},
			},
			"/another/path": path{
				GET: &endpoint{
					Summary:     "GET Summary",
					Description: "GET Description",
					OperationID: "GETAnotherOperationID",
					Group:       "Another",
				},
				PATCH: &endpoint{
					Summary:     "PATCH Summary",
					Description: "PATCH Description",
					OperationID: "PATCHAnotherOperationID",
					Group:       "Another",
				},
				DELETE: &endpoint{
					Summary:     "DELETE Summary",
					Description: "DELETE Description\nas multiline comment",
					OperationID: "DELETEAnotherOperationID",
					Group:       "Another",
				},
			},
		},
	}

	specBytes, err := yaml.Marshal(spec)
	require.NoError(t, err)

	t.Run("produces expected results", func(t *testing.T) {
		specReader := bytes.NewBuffer(specBytes)
		outputWriter := bytes.NewBuffer(nil)
		err = Generate(specReader, outputWriter, Options{})
		require.NoError(t, err)

		expected := `package openapi

// This file is auto-generated, don't modify it manually

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi"
)

// AnotherHandler handles the operations of the 'Another' handler group.
type AnotherHandler interface {
	// DELETEAnotherOperationID DELETE Description
	// as multiline comment
	DELETEAnotherOperationID(w http.ResponseWriter, r *http.Request)
	// GETAnotherOperationID GET Description
	GETAnotherOperationID(w http.ResponseWriter, r *http.Request)
	// PATCHAnotherOperationID PATCH Description
	PATCHAnotherOperationID(w http.ResponseWriter, r *http.Request)
}

// GroupHandler handles the operations of the 'Group' handler group.
type GroupHandler interface {
	// DELETEOperationID DELETE Description
	DELETEOperationID(w http.ResponseWriter, r *http.Request)
	// GETOperationID GET Description
	GETOperationID(w http.ResponseWriter, r *http.Request)
	// PATCHOperationID PATCH Description
	PATCHOperationID(w http.ResponseWriter, r *http.Request)
	// POSTOperationID POST Description
	POSTOperationID(w http.ResponseWriter, r *http.Request)
	// PUTOperationID PUT Description
	PUTOperationID(w http.ResponseWriter, r *http.Request)
}

// NewRouter creates a new router for the spec and the given handlers.
// Title
//
// # Description
//
// Version
func NewRouter(
	anotherHandler AnotherHandler,
	groupHandler GroupHandler,
) http.Handler {

	r := chi.NewRouter()

	// 'Another' group

	// '/another/path'
	r.Options("/another/path", optionsHandlerFunc(
		http.MethodDelete,
		http.MethodGet,
		http.MethodPatch,
	))
	r.Delete("/another/path", anotherHandler.DELETEAnotherOperationID)
	r.Get("/another/path", anotherHandler.GETAnotherOperationID)
	r.Patch("/another/path", anotherHandler.PATCHAnotherOperationID)

	// 'Group' group

	// '/some/path'
	r.Options("/some/path", optionsHandlerFunc(
		http.MethodDelete,
		http.MethodGet,
		http.MethodPatch,
		http.MethodPost,
		http.MethodPut,
	))
	r.Delete("/some/path", groupHandler.DELETEOperationID)
	r.Get("/some/path", groupHandler.GETOperationID)
	r.Patch("/some/path", groupHandler.PATCHOperationID)
	r.Post("/some/path", groupHandler.POSTOperationID)
	r.Put("/some/path", groupHandler.PUTOperationID)

	return r
}

func optionsHandlerFunc(allowedMethods ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Allow", strings.Join(allowedMethods, ", "))
	}
}
`

		require.Equal(t, expected, outputWriter.String())
	})

	t.Run("produces deterministic results", func(t *testing.T) {
		specReader := bytes.NewBuffer(specBytes)
		outputWriter := bytes.NewBuffer(nil)
		err = Generate(specReader, outputWriter, Options{})
		require.NoError(t, err)

		for i := 0; i < 100; i++ {
			newReader := bytes.NewBuffer(specBytes)
			newWriter := bytes.NewBuffer(nil)
			err = Generate(newReader, newWriter, Options{})
			require.NoError(t, err)
			require.Equal(t, outputWriter.String(), newWriter.String())
		}
	})
}
