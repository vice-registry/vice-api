package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/vice-registry/vice-api/models"
)

// NewViceAPI creates a new Vice instance
func NewViceAPI(spec *loads.Document) *ViceAPI {
	return &ViceAPI{
		handlers:        make(map[string]map[string]http.Handler),
		formats:         strfmt.Default,
		defaultConsumes: "application/json",
		defaultProduces: "application/json",
		ServerShutdown:  func() {},
		spec:            spec,
		ServeError:      errors.ServeError,
		JSONConsumer:    runtime.JSONConsumer(),
		XMLConsumer:     runtime.XMLConsumer(),
		JSONProducer:    runtime.JSONProducer(),
		CreateEnvironmentHandler: CreateEnvironmentHandlerFunc(func(params CreateEnvironmentParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation CreateEnvironment has not yet been implemented")
		}),
		CreateImageHandler: CreateImageHandlerFunc(func(params CreateImageParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation CreateImage has not yet been implemented")
		}),
		DeleteDeploymentHandler: DeleteDeploymentHandlerFunc(func(params DeleteDeploymentParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation DeleteDeployment has not yet been implemented")
		}),
		DeleteEnvironmentHandler: DeleteEnvironmentHandlerFunc(func(params DeleteEnvironmentParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation DeleteEnvironment has not yet been implemented")
		}),
		DeleteImageHandler: DeleteImageHandlerFunc(func(params DeleteImageParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation DeleteImage has not yet been implemented")
		}),
		DeployImageHandler: DeployImageHandlerFunc(func(params DeployImageParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation DeployImage has not yet been implemented")
		}),
		FindDeploymentsHandler: FindDeploymentsHandlerFunc(func(params FindDeploymentsParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation FindDeployments has not yet been implemented")
		}),
		FindEnvironmentHandler: FindEnvironmentHandlerFunc(func(params FindEnvironmentParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation FindEnvironment has not yet been implemented")
		}),
		FindImagesHandler: FindImagesHandlerFunc(func(params FindImagesParams) middleware.Responder {
			return middleware.NotImplemented("operation FindImages has not yet been implemented")
		}),
		GetDeploymentHandler: GetDeploymentHandlerFunc(func(params GetDeploymentParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation GetDeployment has not yet been implemented")
		}),
		GetEnvironmentHandler: GetEnvironmentHandlerFunc(func(params GetEnvironmentParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation GetEnvironment has not yet been implemented")
		}),
		GetImageHandler: GetImageHandlerFunc(func(params GetImageParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation GetImage has not yet been implemented")
		}),
		GetRuntimeStatsHandler: GetRuntimeStatsHandlerFunc(func(params GetRuntimeStatsParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation GetRuntimeStats has not yet been implemented")
		}),
		UpdateEnvironmentHandler: UpdateEnvironmentHandlerFunc(func(params UpdateEnvironmentParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation UpdateEnvironment has not yet been implemented")
		}),
		UpdateImageHandler: UpdateImageHandlerFunc(func(params UpdateImageParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation UpdateImage has not yet been implemented")
		}),

		// Applies when the Authorization header is set with the Basic scheme
		ViceAuthAuth: func(user string, pass string) (*models.User, error) {
			return nil, errors.NotImplemented("basic auth  (vice_auth) has not yet been implemented")
		},
	}
}

/*ViceAPI API endpoint of the ViCE Image Registry: search, deploy, and import virtual environments from supported execution environments like OpenStack, Docker, or bwLehrpool. */
type ViceAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer
	// XMLConsumer registers a consumer for a "application/xml" mime type
	XMLConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// ViceAuthAuth registers a function that takes username and password and returns a principal
	// it performs authentication with basic auth
	ViceAuthAuth func(string, string) (*models.User, error)

	// CreateEnvironmentHandler sets the operation handler for the create environment operation
	CreateEnvironmentHandler CreateEnvironmentHandler
	// CreateImageHandler sets the operation handler for the create image operation
	CreateImageHandler CreateImageHandler
	// DeleteDeploymentHandler sets the operation handler for the delete deployment operation
	DeleteDeploymentHandler DeleteDeploymentHandler
	// DeleteEnvironmentHandler sets the operation handler for the delete environment operation
	DeleteEnvironmentHandler DeleteEnvironmentHandler
	// DeleteImageHandler sets the operation handler for the delete image operation
	DeleteImageHandler DeleteImageHandler
	// DeployImageHandler sets the operation handler for the deploy image operation
	DeployImageHandler DeployImageHandler
	// FindDeploymentsHandler sets the operation handler for the find deployments operation
	FindDeploymentsHandler FindDeploymentsHandler
	// FindEnvironmentHandler sets the operation handler for the find environment operation
	FindEnvironmentHandler FindEnvironmentHandler
	// FindImagesHandler sets the operation handler for the find images operation
	FindImagesHandler FindImagesHandler
	// GetDeploymentHandler sets the operation handler for the get deployment operation
	GetDeploymentHandler GetDeploymentHandler
	// GetEnvironmentHandler sets the operation handler for the get environment operation
	GetEnvironmentHandler GetEnvironmentHandler
	// GetImageHandler sets the operation handler for the get image operation
	GetImageHandler GetImageHandler
	// GetRuntimeStatsHandler sets the operation handler for the get runtime stats operation
	GetRuntimeStatsHandler GetRuntimeStatsHandler
	// UpdateEnvironmentHandler sets the operation handler for the update environment operation
	UpdateEnvironmentHandler UpdateEnvironmentHandler
	// UpdateImageHandler sets the operation handler for the update image operation
	UpdateImageHandler UpdateImageHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *ViceAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *ViceAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *ViceAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *ViceAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *ViceAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *ViceAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *ViceAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the ViceAPI
func (o *ViceAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.XMLConsumer == nil {
		unregistered = append(unregistered, "XMLConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.ViceAuthAuth == nil {
		unregistered = append(unregistered, "ViceAuthAuth")
	}

	if o.CreateEnvironmentHandler == nil {
		unregistered = append(unregistered, "CreateEnvironmentHandler")
	}

	if o.CreateImageHandler == nil {
		unregistered = append(unregistered, "CreateImageHandler")
	}

	if o.DeleteDeploymentHandler == nil {
		unregistered = append(unregistered, "DeleteDeploymentHandler")
	}

	if o.DeleteEnvironmentHandler == nil {
		unregistered = append(unregistered, "DeleteEnvironmentHandler")
	}

	if o.DeleteImageHandler == nil {
		unregistered = append(unregistered, "DeleteImageHandler")
	}

	if o.DeployImageHandler == nil {
		unregistered = append(unregistered, "DeployImageHandler")
	}

	if o.FindDeploymentsHandler == nil {
		unregistered = append(unregistered, "FindDeploymentsHandler")
	}

	if o.FindEnvironmentHandler == nil {
		unregistered = append(unregistered, "FindEnvironmentHandler")
	}

	if o.FindImagesHandler == nil {
		unregistered = append(unregistered, "FindImagesHandler")
	}

	if o.GetDeploymentHandler == nil {
		unregistered = append(unregistered, "GetDeploymentHandler")
	}

	if o.GetEnvironmentHandler == nil {
		unregistered = append(unregistered, "GetEnvironmentHandler")
	}

	if o.GetImageHandler == nil {
		unregistered = append(unregistered, "GetImageHandler")
	}

	if o.GetRuntimeStatsHandler == nil {
		unregistered = append(unregistered, "GetRuntimeStatsHandler")
	}

	if o.UpdateEnvironmentHandler == nil {
		unregistered = append(unregistered, "UpdateEnvironmentHandler")
	}

	if o.UpdateImageHandler == nil {
		unregistered = append(unregistered, "UpdateImageHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *ViceAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *ViceAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	result := make(map[string]runtime.Authenticator)
	for name, scheme := range schemes {
		switch name {

		case "vice_auth":
			_ = scheme
			result[name] = security.BasicAuth(func(username, password string) (interface{}, error) {
				return o.ViceAuthAuth(username, password)
			})

		}
	}
	return result

}

// ConsumersFor gets the consumers for the specified media types
func (o *ViceAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		case "application/xml":
			result["application/xml"] = o.XMLConsumer

		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *ViceAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *ViceAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the vice API
func (o *ViceAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *ViceAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/environments"] = NewCreateEnvironment(o.context, o.CreateEnvironmentHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/images"] = NewCreateImage(o.context, o.CreateImageHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/deployment/{deploymentId}"] = NewDeleteDeployment(o.context, o.DeleteDeploymentHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/environment/{environmentId}"] = NewDeleteEnvironment(o.context, o.DeleteEnvironmentHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/image/{imageId}"] = NewDeleteImage(o.context, o.DeleteImageHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/deploy"] = NewDeployImage(o.context, o.DeployImageHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/deployments"] = NewFindDeployments(o.context, o.FindDeploymentsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/environments"] = NewFindEnvironment(o.context, o.FindEnvironmentHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/images"] = NewFindImages(o.context, o.FindImagesHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/deployment/{deploymentId}"] = NewGetDeployment(o.context, o.GetDeploymentHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/environment/{environmentId}"] = NewGetEnvironment(o.context, o.GetEnvironmentHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/image/{imageId}"] = NewGetImage(o.context, o.GetImageHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/runtimestats"] = NewGetRuntimeStats(o.context, o.GetRuntimeStatsHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/environments"] = NewUpdateEnvironment(o.context, o.UpdateEnvironmentHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/images"] = NewUpdateImage(o.context, o.UpdateImageHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *ViceAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middelware as you see fit
func (o *ViceAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}
