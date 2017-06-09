package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	"omi-gitlab.e-technik.uni-ulm.de/vice/vice-api/models"
)

// GetImageHandlerFunc turns a function with the right signature into a get image handler
type GetImageHandlerFunc func(GetImageParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn GetImageHandlerFunc) Handle(params GetImageParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// GetImageHandler interface for that can handle valid get image params
type GetImageHandler interface {
	Handle(GetImageParams, *models.User) middleware.Responder
}

// NewGetImage creates a new http.Handler for the get image operation
func NewGetImage(ctx *middleware.Context, handler GetImageHandler) *GetImage {
	return &GetImage{Context: ctx, Handler: handler}
}

/*GetImage swagger:route GET /image/{imageId} getImage

Get an image by id

*/
type GetImage struct {
	Context *middleware.Context
	Handler GetImageHandler
}

func (o *GetImage) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetImageParams()

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
