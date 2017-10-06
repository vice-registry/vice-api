// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/vice-registry/vice-util/models"
)

// DeleteDeploymentHandlerFunc turns a function with the right signature into a delete deployment handler
type DeleteDeploymentHandlerFunc func(DeleteDeploymentParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteDeploymentHandlerFunc) Handle(params DeleteDeploymentParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// DeleteDeploymentHandler interface for that can handle valid delete deployment params
type DeleteDeploymentHandler interface {
	Handle(DeleteDeploymentParams, *models.User) middleware.Responder
}

// NewDeleteDeployment creates a new http.Handler for the delete deployment operation
func NewDeleteDeployment(ctx *middleware.Context, handler DeleteDeploymentHandler) *DeleteDeployment {
	return &DeleteDeployment{Context: ctx, Handler: handler}
}

/*DeleteDeployment swagger:route DELETE /deployment/{deploymentId} deleteDeployment

Removes an image deployment from the environment

*/
type DeleteDeployment struct {
	Context *middleware.Context
	Handler DeleteDeploymentHandler
}

func (o *DeleteDeployment) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteDeploymentParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
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
