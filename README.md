openapi-generator-go
====================

![test](https://github.com/contiamo/openapi-generator-go/workflows/test/badge.svg)

Generate go code from an openapi spec!

```bash
# generate structs, enums and a router from a spec file
openapi-generator-go generate --spec ./example/api.yaml --output ./example/generated
# generate models and router independently
openapi-generator-go generate models --spec ./example/api.yaml --output ./example/models --package-name models
openapi-generator-go generate router --spec ./example/api.yaml --output ./example/router --package-name router
```

## Why
The platform team at Contiamo has a schema first development flow. This means we write and review the OpenAPI spec before we write the implementation.

To make this easier we started using (or tried to use) various OpenAPI code generators to reduce busy work and ensure precise consistent implementations. Unfortunately, we never really enjoyed the output from these generators and have created our own generator that matches our preferred style and conventions.

## What
The generator consists of two parts

1. generator for the HTTP server mux using [go-chi](https://github.com/go-chi/chi), and
2. a model generator for the request and response objects.

### The router generator
We chose `go-chi` because it is fast, supports the standard library Handler interface, and is easy to add middlewares using `r.Use`.

The generator will create a method `NewRouter` and handler interfaces that you must then implement and pass to `NewRouter`. It is important to note that it does not generate any handler logic, just the required interfaces.  Handlers can be grouped into specific interfaces by adding `x-handler-group: <Name>` to your schema. This is one of the only schema extensions we have added to improve our code generation. The handler interface will be called `NameHandler` and will have a method matching the `operationId` from the spec.

For example, each of our projects will implement a project specific `NewRouter` that does the handler configuration and  initialization of the generated router. In the next example, `openapi.NewRouter` is generated using

```sh
openapi-generator-go generate router --spec ./example/api.yaml --output ./example/router --package-name openapi
```
The project specific router initialization:
```go
func NewRouter(ctx context.Context, db *sql.DB, cfg config.Config) (http.Handler, error) {
    // do setup stuff

    // and probably some more setup for each of the handler implementations
    userHandler := handlers.NewUserHandler(ctx, db)
    resourcesHandler := handlers.NewResource(ctx, db)

    // openapi-generator-go generate router --spec ./example/api.yaml --output ./example/router --package-name openapi
    apiRouter := openapi.NewRouter(
        // the UsersHandler interface contains a method for endpoint with
        // x-handler-group: Users
        userHandler,

        // ResourcesHandler interace contains a method for each endpoint with
        // x-handler-group: Resources
        resourcesHandler,
    )

    root.Route("/", func(api chi.Router) {
        // now  init and dd your middlewares
		r := middlewares.WithRecovery(os.Stderr, cfg.Debug)
		t := middlewares.WithTracing(config.ApplicationName, nil, middlewares.ChiRouteName)
		l := middlewares.WithLogging(config.ApplicationName)
		m := middlewares.WithMetrics(config.ApplicationName, nil)
		a := authorization.NewMiddleware(cfg.Authorization.HeaderName, publicKey)

		api.Use(r.WrapHandler)
		api.Use(t.WrapHandler)
		api.Use(l.WrapHandler)
		api.Use(m.WrapHandler)
		api.Use(a.WrapHandler)

        // mount the generate router
		api.Mount("/", apiRouter)
	})

	return root, nil

}
```

### The model generator
The model generator is a work in progress, but covers the most common cases we need in a REST API.

You can find various examples of the what is supported and the corresponding output in our [test fixtures](./pkg/generators/models/testdata/cases).

Our generator differs from the official OpenAPI generator tools by also providing getters for many fields, which makes it easier to define and work with interfaces of the models.

We also provide better support for :
* [enums](./pkg/generators/models/testdata/cases/enums/expected/model_filter_type.go), enum support also includes some utility methods that make validating the enums much easier,
  * Now including support for string, integer, and number (i.e. decimal) enums.
* [arrays](./pkg/generators/models/testdata/cases/typed_arrays/expected/model_foo.go),
* [allOf](./pkg/generators/models/testdata/cases/allof1/expected/model_foo.go),
* and [oneOf](./pkg/generators/models/testdata/cases/oneof/expected/model_foo.go)
that feel more natural in Go. Creating as many strong types as is possible and using `interface{}` otherwise. These cases often failed or generated non-compilable code with the official generator.
