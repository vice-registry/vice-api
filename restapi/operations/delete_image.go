// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/vice-registry/vice-util/models"
)

// DeleteImageHandlerFunc turns a function with the right signature into a delete image handler
type DeleteImageHandlerFunc func(DeleteImageParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteImageHandlerFunc) Handle(params DeleteImageParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// DeleteImageHandler interface for that can handle valid delete image params
type DeleteImageHandler interface {
	Handle(DeleteImageParams, *models.User) middleware.Responder
}

// NewDeleteImage creates a new http.Handler for the delete image operation
func NewDeleteImage(ctx *middleware.Context, handler DeleteImageHandler) *DeleteImage {
	return &DeleteImage{Context: ctx, Handler: handler}
}

/*DeleteImage swagger:route DELETE /image/{imageId} deleteImage

Deletes an image

*/
type DeleteImage struct {
	Context *middleware.Context
	Handler DeleteImageHandler
}

func (o *DeleteImage) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteImageParams()

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
