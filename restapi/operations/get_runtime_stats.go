package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/vice-registry/vice-api/models"
)

// GetRuntimeStatsHandlerFunc turns a function with the right signature into a get runtime stats handler
type GetRuntimeStatsHandlerFunc func(GetRuntimeStatsParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRuntimeStatsHandlerFunc) Handle(params GetRuntimeStatsParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// GetRuntimeStatsHandler interface for that can handle valid get runtime stats params
type GetRuntimeStatsHandler interface {
	Handle(GetRuntimeStatsParams, *models.User) middleware.Responder
}

// NewGetRuntimeStats creates a new http.Handler for the get runtime stats operation
func NewGetRuntimeStats(ctx *middleware.Context, handler GetRuntimeStatsHandler) *GetRuntimeStats {
	return &GetRuntimeStats{Context: ctx, Handler: handler}
}

/*GetRuntimeStats swagger:route GET /runtimestats getRuntimeStats

Provide internal statistics about import/export queue length and workers

*/
type GetRuntimeStats struct {
	Context *middleware.Context
	Handler GetRuntimeStatsHandler
}

func (o *GetRuntimeStats) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetRuntimeStatsParams()

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
