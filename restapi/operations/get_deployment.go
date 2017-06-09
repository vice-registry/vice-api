package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	"omi-gitlab.e-technik.uni-ulm.de/vice/vice-api/models"
)

// GetDeploymentHandlerFunc turns a function with the right signature into a get deployment handler
type GetDeploymentHandlerFunc func(GetDeploymentParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn GetDeploymentHandlerFunc) Handle(params GetDeploymentParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// GetDeploymentHandler interface for that can handle valid get deployment params
type GetDeploymentHandler interface {
	Handle(GetDeploymentParams, *models.User) middleware.Responder
}

// NewGetDeployment creates a new http.Handler for the get deployment operation
func NewGetDeployment(ctx *middleware.Context, handler GetDeploymentHandler) *GetDeployment {
	return &GetDeployment{Context: ctx, Handler: handler}
}

/*GetDeployment swagger:route GET /deployment/{deploymentId} getDeployment

Get a deployment

*/
type GetDeployment struct {
	Context *middleware.Context
	Handler GetDeploymentHandler
}

func (o *GetDeployment) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetDeploymentParams()

	uprinc, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	var principal *models.User
	if uprinc != nil {
		principal = uprinc.(*models.User) // this is really a models.User, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}