// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDownloadImageParams creates a new DownloadImageParams object
// with the default values initialized.
func NewDownloadImageParams() DownloadImageParams {
	var ()
	return DownloadImageParams{}
}

// DownloadImageParams contains all the bound params for the download image operation
// typically these are obtained from a http.Request
//
// swagger:parameters downloadImage
type DownloadImageParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*
	  Required: true
	  In: path
	*/
	ImageID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *DownloadImageParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rImageID, rhkImageID, _ := route.Params.GetOK("imageId")
	if err := o.bindImageID(rImageID, rhkImageID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DownloadImageParams) bindImageID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.ImageID = raw

	return nil
}
