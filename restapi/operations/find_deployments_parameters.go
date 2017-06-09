package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
)

// NewFindDeploymentsParams creates a new FindDeploymentsParams object
// with the default values initialized.
func NewFindDeploymentsParams() FindDeploymentsParams {
	var ()
	return FindDeploymentsParams{}
}

// FindDeploymentsParams contains all the bound params for the find deployments operation
// typically these are obtained from a http.Request
//
// swagger:parameters findDeployments
type FindDeploymentsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *FindDeploymentsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}