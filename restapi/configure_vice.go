package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	graceful "github.com/tylerb/graceful"

	"omi-gitlab.e-technik.uni-ulm.de/vice/vice-api/models"
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

	var executionEnvironments = make([]*models.ExecutionEnvironment, 1)
	var credentials = make([]*models.Credentials, 1)
	//var images = make([]*models.Image, 1)
	var managementLayers = make([]*models.ManagementLayer, 1)
	var runtimeTechnologies = make([]*models.RuntimeTechnology, 1)
	var users = make([]*models.User, 1)

	users[0] = &models.User{1, "geheim", "user"}
	credentials[0] = &models.Credentials{"endpoint", 1, "pass", "user"}
	managementLayers[0] = &models.ManagementLayer{"OpenStack", "cloudcomputing", "Kilo"}
	runtimeTechnologies[0] = &models.RuntimeTechnology{"KVM", "virtualmachine", "4.10.13"}
	executionEnvironments[0] = &models.ExecutionEnvironment{credentials[0], 1, managementLayers[0], runtimeTechnologies[0], users[0]}

	// Applies when the Authorization header is set with the Basic scheme
	api.ViceAuthAuth = func(user string, pass string) (interface{}, error) {
		for _, userentry := range users {
			if user == userentry.Username && pass == userentry.Password {
				// allow
				return userentry, nil
			}
		}
		// deny
		api.Logger("Access attempt with incorrect user credentials.")
		return nil, errors.Unauthenticated("basic")
	}

	api.CreateExecutionEnvironmentHandler = operations.CreateExecutionEnvironmentHandlerFunc(func(params operations.CreateExecutionEnvironmentParams, principal interface{}) middleware.Responder {
		executionEnvironments = append(executionEnvironments, params.Body)
		return operations.NewCreateExecutionEnvironmentCreated().WithPayload(params.Body)
	})
	api.DeleteExecutionEnvironmentHandler = operations.DeleteExecutionEnvironmentHandlerFunc(func(params operations.DeleteExecutionEnvironmentParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation .DeleteExecutionEnvironment has not yet been implemented")
	})
	api.FindExecutionEnvironmentHandler = operations.FindExecutionEnvironmentHandlerFunc(func(params operations.FindExecutionEnvironmentParams, principal interface{}) middleware.Responder {
		return operations.NewFindExecutionEnvironmentOK().WithPayload(executionEnvironments)
	})
	api.FindImagesHandler = operations.FindImagesHandlerFunc(func(params operations.FindImagesParams) middleware.Responder {
		return middleware.NotImplemented("operation .FindImages has not yet been implemented")
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
