package api

// This file is auto-generated, don't modify it manually

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi"
)


// EndpointsHandler handles the operations of the 'Endpoints' handler group.
type EndpointsHandler interface {
	// ServeAEndpoint Returns "ok!" message when the service is running
	ServeAEndpoint(w http.ResponseWriter, r *http.Request)
}

// NewRouter creates a new router for the spec and the given handlers.
// Test Service
//
// 
//
// 0.1.0
//
func NewRouter(
	endpointsHandler EndpointsHandler,
) http.Handler {

	r := chi.NewRouter()

// 'Endpoints' group

// '/aEndpoint'
r.Options("/aEndpoint", optionsHandlerFunc(
	http.MethodGet,
))
r.Get("/aEndpoint", endpointsHandler.ServeAEndpoint)

	return r
}

func optionsHandlerFunc(allowedMethods ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Allow", strings.Join(allowedMethods, ", "))
	}
}
