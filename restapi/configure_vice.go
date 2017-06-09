package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	graceful "github.com/tylerb/graceful"

	cors "github.com/rs/cors"

	"omi-gitlab.e-technik.uni-ulm.de/vice/vice-api/actions"
	"omi-gitlab.e-technik.uni-ulm.de/vice/vice-api/models"
	"omi-gitlab.e-technik.uni-ulm.de/vice/vice-api/persistence"
	"omi-gitlab.e-technik.uni-ulm.de/vice/vice-api/restapi/operations"
)

// CouchbaseFlags cli Configuration options for couchbase connection
var CouchbaseFlags = struct {
	Location string `short:"c" long:"couchbase-location" description:"Location of the Couchbase cluster to connect to (e.g. localhost)"`
	Username string `short:"u" long:"couchbase-user" description:"Username to log in to Couchbase cluster"`
	Password string `short:"p" long:"couchbase-pass" description:"Password to log in to Couchbase cluster"`
}{}

// RabbitmqFlags cli Configuration options for rabbitmq connection
var RabbitmqFlags = struct {
	Location string `short:"r" long:"rabbitmq-location" description:"Location of the RabbitMQ to connect to (e.g. localhost)"`
	Username string `short:"" long:"rabbitmq-user" description:"Username to log in to RabbitMQ"`
	Password string `short:"" long:"rabbitmq-pass" description:"Password to log in to RabbitMQ"`
}{}

func configureFlags(api *operations.ViceAPI) {
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		swag.CommandLineOptionsGroup{
			ShortDescription: "Couchbase Connection",
			LongDescription:  "Configuration options for couchbase connection",
			Options:          &CouchbaseFlags,
		},
		swag.CommandLineOptionsGroup{
			ShortDescription: "RabbitMQ Connection",
			LongDescription:  "Configuration options for RabbitMQ connection",
			Options:          &RabbitmqFlags,
		},
	}
}

func configureAPI(api *operations.ViceAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// initialize couchbase
	persistence.SetCouchbaseCredentials(CouchbaseFlags.Location, CouchbaseFlags.Username, CouchbaseFlags.Password)
	persistence.InitViceCouchbase()

	// initialize rabbitmq
	actions.SetRabbitmqCredentials(RabbitmqFlags.Location, RabbitmqFlags.Username, RabbitmqFlags.Password)

	api.JSONConsumer = runtime.JSONConsumer()
	api.XMLConsumer = runtime.XMLConsumer()
	api.JSONProducer = runtime.JSONProducer()

	// Applies when the Authorization header is set with the Basic scheme
	api.ViceAuthAuth = func(user string, pass string) (*models.User, error) {
		userentry, err := persistence.GetUserByName(user)
		if err != nil {
			return nil, errors.New(500, "internal server error")
		}
		if userentry.Password == pass {
			// allow
			return userentry, nil
		}
		// deny
		return nil, errors.New(401, "invalid authentication")
	}

	// Environment
	api.FindEnvironmentHandler = operations.FindEnvironmentHandlerFunc(func(params operations.FindEnvironmentParams, principal *models.User) middleware.Responder {
		if principal == nil {
			return operations.NewFindEnvironmentUnauthorized()
		}
		environments, err := persistence.GetEnvironments()
		if err != nil {
			return operations.NewFindEnvironmentInternalServerError()
		}
		return operations.NewFindEnvironmentOK().WithPayload(environments)
	})
	api.CreateEnvironmentHandler = operations.CreateEnvironmentHandlerFunc(func(params operations.CreateEnvironmentParams, principal *models.User) middleware.Responder {
		if principal == nil {
			return operations.NewCreateEnvironmentUnauthorized()
		}
		environment, err := persistence.CreateEnvironment(params.Body)
		if err != nil {
			return operations.NewCreateEnvironmentInternalServerError()
		}
		return operations.NewCreateEnvironmentCreated().WithPayload(environment)
	})
	api.UpdateEnvironmentHandler = operations.UpdateEnvironmentHandlerFunc(func(params operations.UpdateEnvironmentParams, principal *models.User) middleware.Responder {
		if principal == nil {
			return operations.NewUpdateEnvironmentUnauthorized()
		}
		environment, err := persistence.GetEnvironment(params.Body.ID)
		if err != nil {
			return operations.NewUpdateEnvironmentNotFound()
		}
		if environment.Userid != principal.ID {
			return operations.NewUpdateEnvironmentUnauthorized()
		}

		environment, err = persistence.UpdateEnvironment(params.Body)
		if err != nil {
			return operations.NewUpdateEnvironmentInternalServerError()
		}
		return operations.NewUpdateEnvironmentCreated().WithPayload(environment)
	})
	api.GetEnvironmentHandler = operations.GetEnvironmentHandlerFunc(func(params operations.GetEnvironmentParams, principal *models.User) middleware.Responder {
		if principal == nil {
			return operations.NewGetEnvironmentUnauthorized()
		}
		environment, err := persistence.GetEnvironment(params.EnvironmentID)
		if err != nil {
			return operations.NewUpdateEnvironmentNotFound()
		}
		if environment.Userid != principal.ID {
			return operations.NewUpdateEnvironmentUnauthorized()
		}
		return operations.NewGetEnvironmentOK().WithPayload(environment)
	})
	api.DeleteEnvironmentHandler = operations.DeleteEnvironmentHandlerFunc(func(params operations.DeleteEnvironmentParams, principal *models.User) middleware.Responder {
		if principal == nil {
			return operations.NewDeleteEnvironmentUnauthorized()
		}
		environment, err := persistence.GetEnvironment(params.EnvironmentID)
		if err != nil {
			return operations.NewUpdateEnvironmentNotFound()
		}
		if environment.Userid != principal.ID {
			return operations.NewUpdateEnvironmentUnauthorized()
		}
		err = persistence.DeleteEnvironment(params.EnvironmentID)
		if err != nil {
			return operations.NewDeleteEnvironmentInternalServerError()
		}
		return operations.NewDeleteEnvironmentOK()
	})

	// Image
	api.FindImagesHandler = operations.FindImagesHandlerFunc(func(params operations.FindImagesParams) middleware.Responder {
		images, err := persistence.GetImages()
		if err != nil {
			return operations.NewFindImagesInternalServerError()
		}
		return operations.NewFindImagesOK().WithPayload(images)
	})
	api.CreateImageHandler = operations.CreateImageHandlerFunc(func(params operations.CreateImageParams, principal *models.User) middleware.Responder {
		if principal == nil {
			return operations.NewCreateImageUnauthorized()
		}
		// store image in couchbase
		image, err := persistence.CreateImage(params.Body)
		if err != nil {
			return operations.NewCreateImageInternalServerError()
		}
		// send import action
		actions.NewImportAction(image)
		return operations.NewCreateImageCreated().WithPayload(image)
	})
	api.UpdateImageHandler = operations.UpdateImageHandlerFunc(func(params operations.UpdateImageParams, principal *models.User) middleware.Responder {
		if principal == nil {
			return operations.NewUpdateImageUnauthorized()
		}
		image, err := persistence.GetImage(params.Body.ID)
		if err != nil {
			return operations.NewUpdateImageNotFound()
		}
		if image.Userid != principal.ID {
			return operations.NewUpdateImageUnauthorized()
		}
		// store image in couchbase
		image, err = persistence.UpdateImage(params.Body)
		if err != nil {
			return operations.NewUpdateImageInternalServerError()
		}
		// send import action
		actions.NewImportAction(image)
		return operations.NewUpdateImageCreated().WithPayload(image)
	})
	api.GetImageHandler = operations.GetImageHandlerFunc(func(params operations.GetImageParams, principal *models.User) middleware.Responder {
		if principal == nil {
			return operations.NewGetImageUnauthorized()
		}
		image, err := persistence.GetImage(params.ImageID)
		if err != nil {
			return operations.NewGetImageNotFound()
		}
		if image.Userid != principal.ID {
			return operations.NewGetImageUnauthorized()
		}
		return operations.NewGetImageOK().WithPayload(image)
	})
	api.DeleteImageHandler = operations.DeleteImageHandlerFunc(func(params operations.DeleteImageParams, principal *models.User) middleware.Responder {
		if principal == nil {
			return operations.NewDeleteImageUnauthorized()
		}
		image, err := persistence.GetImage(params.ImageID)
		if err != nil {
			return operations.NewDeleteImageNotFound()
		}
		if image.Userid != principal.ID {
			return operations.NewDeleteImageUnauthorized()
		}
		err = persistence.DeleteImage(params.ImageID)
		if err != nil {
			return operations.NewDeleteImageInternalServerError()
		}
		return operations.NewDeleteImageOK()
	})

	// Deployment
	api.DeployImageHandler = operations.DeployImageHandlerFunc(func(params operations.DeployImageParams, principal *models.User) middleware.Responder {
		if principal == nil {
			return operations.NewDeployImageUnauthorized()
		}
		// image valid?
		image, err := persistence.GetImage(params.Body.Imageid)
		if err != nil {
			return operations.NewDeployImageMethodNotAllowed()
		}
		if image.Userid != principal.ID {
			return operations.NewDeployImageUnauthorized()
		}
		// environment valid?
		environment, err := persistence.GetEnvironment(params.Body.EnvironmentID)
		if err != nil {
			return operations.NewDeployImageMethodNotAllowed()
		}
		if environment.Userid != principal.ID {
			return operations.NewDeployImageUnauthorized()
		}
		// store deployment in couchbase
		deployment, err := persistence.CreateDeployment(params.Body)
		if err != nil {
			return operations.NewDeployImageInternalServerError()
		}
		// send deployment action
		actions.NewDeployAction(deployment)

		return operations.NewDeployImageCreated().WithPayload(deployment)
	})
	api.FindDeploymentsHandler = operations.FindDeploymentsHandlerFunc(func(params operations.FindDeploymentsParams, principal *models.User) middleware.Responder {
		if principal == nil {
			return operations.NewFindDeploymentsUnauthorized()
		}
		deployments, err := persistence.GetDeployments()
		if err != nil {
			return operations.NewFindDeploymentsInternalServerError()
		}
		return operations.NewFindDeploymentsOK().WithPayload(deployments)
	})
	api.GetDeploymentHandler = operations.GetDeploymentHandlerFunc(func(params operations.GetDeploymentParams, principal *models.User) middleware.Responder {
		if principal == nil {
			return operations.NewGetDeploymentUnauthorized()
		}
		deployment, err := persistence.GetDeployment(params.DeploymentID)
		if err != nil {
			return operations.NewGetDeploymentInternalServerError()
		}
		if deployment.Userid != principal.ID {
			return operations.NewGetDeploymentUnauthorized()
		}
		return operations.NewGetDeploymentOK().WithPayload(deployment)
	})
	api.DeleteDeploymentHandler = operations.DeleteDeploymentHandlerFunc(func(params operations.DeleteDeploymentParams, principal *models.User) middleware.Responder {
		if principal == nil {
			return operations.NewDeleteDeploymentUnauthorized()
		}
		deployment, err := persistence.GetDeployment(params.DeploymentID)
		if err != nil {
			return operations.NewDeleteDeploymentNotFound()
		}
		if deployment.Userid != principal.ID {
			return operations.NewDeleteDeploymentUnauthorized()
		}
		err = persistence.DeleteDeployment(params.DeploymentID)
		if err != nil {
			return operations.NewDeleteDeploymentInternalServerError()
		}
		return operations.NewDeleteDeploymentOK()
	})

	// User
	// TODO

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
	//handleCORS := cors.Default().Handler
	handleCORS := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://petstore.swagger.io", "http://localhost"},
		AllowCredentials: true,
	}).Handler
	return handleCORS(handler)
}
