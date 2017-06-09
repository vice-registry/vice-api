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

//go:generate swagger generate server --target .. --name vice-api --spec ../swagger.yaml --principal models.User

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

	// Applies when the Authorization header is set with the Basic scheme
	api.ViceAuthAuth = func(user string, pass string) (*models.User, error) {
		return nil, errors.NotImplemented("basic auth  (vice_auth) has not yet been implemented")
	}

	api.CreateEnvironmentHandler = operations.CreateEnvironmentHandlerFunc(func(params operations.CreateEnvironmentParams, principal *models.User) middleware.Responder {
		return middleware.NotImplemented("operation .CreateEnvironment has not yet been implemented")
	})
	api.CreateImageHandler = operations.CreateImageHandlerFunc(func(params operations.CreateImageParams, principal *models.User) middleware.Responder {
		return middleware.NotImplemented("operation .CreateImage has not yet been implemented")
	})
	api.DeleteDeploymentHandler = operations.DeleteDeploymentHandlerFunc(func(params operations.DeleteDeploymentParams, principal *models.User) middleware.Responder {
		return middleware.NotImplemented("operation .DeleteDeployment has not yet been implemented")
	})
	api.DeleteEnvironmentHandler = operations.DeleteEnvironmentHandlerFunc(func(params operations.DeleteEnvironmentParams, principal *models.User) middleware.Responder {
		return middleware.NotImplemented("operation .DeleteEnvironment has not yet been implemented")
	})
	api.DeleteImageHandler = operations.DeleteImageHandlerFunc(func(params operations.DeleteImageParams, principal *models.User) middleware.Responder {
		return middleware.NotImplemented("operation .DeleteImage has not yet been implemented")
	})
	api.DeployImageHandler = operations.DeployImageHandlerFunc(func(params operations.DeployImageParams, principal *models.User) middleware.Responder {
		return middleware.NotImplemented("operation .DeployImage has not yet been implemented")
	})
	api.FindDeploymentsHandler = operations.FindDeploymentsHandlerFunc(func(params operations.FindDeploymentsParams, principal *models.User) middleware.Responder {
		return middleware.NotImplemented("operation .FindDeployments has not yet been implemented")
	})
	api.FindEnvironmentHandler = operations.FindEnvironmentHandlerFunc(func(params operations.FindEnvironmentParams, principal *models.User) middleware.Responder {
		return middleware.NotImplemented("operation .FindEnvironment has not yet been implemented")
	})
	api.FindImagesHandler = operations.FindImagesHandlerFunc(func(params operations.FindImagesParams) middleware.Responder {
		return middleware.NotImplemented("operation .FindImages has not yet been implemented")
	})
	api.GetDeploymentHandler = operations.GetDeploymentHandlerFunc(func(params operations.GetDeploymentParams, principal *models.User) middleware.Responder {
		return middleware.NotImplemented("operation .GetDeployment has not yet been implemented")
	})
	api.GetEnvironmentHandler = operations.GetEnvironmentHandlerFunc(func(params operations.GetEnvironmentParams, principal *models.User) middleware.Responder {
		return middleware.NotImplemented("operation .GetEnvironment has not yet been implemented")
	})
	api.GetImageHandler = operations.GetImageHandlerFunc(func(params operations.GetImageParams, principal *models.User) middleware.Responder {
		return middleware.NotImplemented("operation .GetImage has not yet been implemented")
	})
	api.UpdateEnvironmentHandler = operations.UpdateEnvironmentHandlerFunc(func(params operations.UpdateEnvironmentParams, principal *models.User) middleware.Responder {
		return middleware.NotImplemented("operation .UpdateEnvironment has not yet been implemented")
	})
	api.UpdateImageHandler = operations.UpdateImageHandlerFunc(func(params operations.UpdateImageParams, principal *models.User) middleware.Responder {
		return middleware.NotImplemented("operation .UpdateImage has not yet been implemented")
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
