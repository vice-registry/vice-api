// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/vice-registry/vice-api/models"
)

// DeployImageHandlerFunc turns a function with the right signature into a deploy image handler
type DeployImageHandlerFunc func(DeployImageParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn DeployImageHandlerFunc) Handle(params DeployImageParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// DeployImageHandler interface for that can handle valid deploy image params
type DeployImageHandler interface {
	Handle(DeployImageParams, *models.User) middleware.Responder
}

// NewDeployImage creates a new http.Handler for the deploy image operation
func NewDeployImage(ctx *middleware.Context, handler DeployImageHandler) *DeployImage {
	return &DeployImage{Context: ctx, Handler: handler}
}

/*DeployImage swagger:route POST /deploy deployImage

Create an image into an environment

*/
type DeployImage struct {
	Context *middleware.Context
	Handler DeployImageHandler
}

func (o *DeployImage) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeployImageParams()

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
