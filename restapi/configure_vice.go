package restapi

import (
	"crypto/tls"
	"log"
	"net/http"

	gocb "gopkg.in/couchbase/gocb.v1"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	graceful "github.com/tylerb/graceful"

	cors "github.com/rs/cors"

	"omi-gitlab.e-technik.uni-ulm.de/vice/vice-api/persistence"
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

	cbCluster, err := gocb.Connect("couchbase://localhost")
	if err != nil {
		log.Fatalln("cannot connect to couchbase: ", err)
	}

	// Applies when the Authorization header is set with the Basic scheme
	api.ViceAuthAuth = func(user string, pass string) (interface{}, error) {
		userentry := persistence.GetUserByName(cbCluster, user)
		if userentry.Password == pass {
			// allow
			return userentry, nil
		}
		// deny
		return nil, errors.New(401, "invalid authentication")
	}

	api.CreateEnvironmentHandler = operations.CreateEnvironmentHandlerFunc(func(params operations.CreateEnvironmentParams, principal interface{}) middleware.Responder {
		environment := persistence.CreateEnvironment(cbCluster, params.Body)
		return operations.NewCreateEnvironmentCreated().WithPayload(environment)
	})
	api.CreateImageHandler = operations.CreateImageHandlerFunc(func(params operations.CreateImageParams, principal interface{}) middleware.Responder {
		image := persistence.CreateImage(cbCluster, params.Body)
		return operations.NewCreateImageCreated().WithPayload(image)
	})
	api.DeleteEnvironmentHandler = operations.DeleteEnvironmentHandlerFunc(func(params operations.DeleteEnvironmentParams, principal interface{}) middleware.Responder {
		persistence.DeleteEnvironment(cbCluster, params.EnvironmentID)
		return operations.NewDeleteEnvironmentOK()
	})
	api.DeleteImageHandler = operations.DeleteImageHandlerFunc(func(params operations.DeleteImageParams, principal interface{}) middleware.Responder {
		persistence.DeleteImage(cbCluster, params.ImageID)
		return operations.NewDeleteImageOK()
	})
	api.FindEnvironmentHandler = operations.FindEnvironmentHandlerFunc(func(params operations.FindEnvironmentParams, principal interface{}) middleware.Responder {
		environments := persistence.GetEnvironments(cbCluster)
		return operations.NewFindEnvironmentOK().WithPayload(environments)
	})
	api.FindImagesHandler = operations.FindImagesHandlerFunc(func(params operations.FindImagesParams) middleware.Responder {
		images := persistence.GetImages(cbCluster)
		return operations.NewFindImagesOK().WithPayload(images)
	})
	api.GetEnvironmentHandler = operations.GetEnvironmentHandlerFunc(func(params operations.GetEnvironmentParams, principal interface{}) middleware.Responder {
		environment := persistence.GetEnvironment(cbCluster, params.EnvironmentID)
		return operations.NewGetEnvironmentOK().WithPayload(environment)
	})
	api.UpdateEnvironmentHandler = operations.UpdateEnvironmentHandlerFunc(func(params operations.UpdateEnvironmentParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation .UpdateEnvironment has not yet been implemented")
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
	handleCORS := cors.Default().Handler

	return handleCORS(handler)
}
