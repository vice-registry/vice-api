package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	graceful "github.com/tylerb/graceful"

	"omi-gitlab.e-technik.uni-ulm.de/vice/vice-api/restapi/operations"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --name vice-api --spec ../swagger.yaml

func configureFlags(api *operations.ViceAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ViceAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// s.api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.XMLConsumer = runtime.XMLConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.ViceAuthAuth = func(token string, scopes []string) (interface{}, error) {
		return nil, errors.NotImplemented("oauth2 bearer auth (vice_auth) has not yet been implemented")
	}

	api.CreateExecutionEnvironmentHandler = operations.CreateExecutionEnvironmentHandlerFunc(func(params operations.CreateExecutionEnvironmentParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation .CreateExecutionEnvironment has not yet been implemented")
	})
	api.DeleteExecutionEnvironmentHandler = operations.DeleteExecutionEnvironmentHandlerFunc(func(params operations.DeleteExecutionEnvironmentParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation .DeleteExecutionEnvironment has not yet been implemented")
	})
	api.FindExecutionEnvironmentHandler = operations.FindExecutionEnvironmentHandlerFunc(func(params operations.FindExecutionEnvironmentParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation .FindExecutionEnvironment has not yet been implemented")
	})
	api.GetExecutionEnvironmentHandler = operations.GetExecutionEnvironmentHandlerFunc(func(params operations.GetExecutionEnvironmentParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation .GetExecutionEnvironment has not yet been implemented")
	})
	api.UpdateExecutionEnvironmentHandler = operations.UpdateExecutionEnvironmentHandlerFunc(func(params operations.UpdateExecutionEnvironmentParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation .UpdateExecutionEnvironment has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
