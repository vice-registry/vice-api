package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"omi-gitlab.e-technik.uni-ulm.de/vice/vice-api/models"
)

// GetImageOKCode is the HTTP code returned for type GetImageOK
const GetImageOKCode int = 200

/*GetImageOK successful operation

swagger:response getImageOK
*/
type GetImageOK struct {

	/*
	  In: Body
	*/
	Payload *models.Image `json:"body,omitempty"`
}

// NewGetImageOK creates GetImageOK with default headers values
func NewGetImageOK() *GetImageOK {
	return &GetImageOK{}
}

// WithPayload adds the payload to the get image o k response
func (o *GetImageOK) WithPayload(payload *models.Image) *GetImageOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get image o k response
func (o *GetImageOK) SetPayload(payload *models.Image) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetImageOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetImageBadRequestCode is the HTTP code returned for type GetImageBadRequest
const GetImageBadRequestCode int = 400

/*GetImageBadRequest Invalid ID supplied

swagger:response getImageBadRequest
*/
type GetImageBadRequest struct {
}

// NewGetImageBadRequest creates GetImageBadRequest with default headers values
func NewGetImageBadRequest() *GetImageBadRequest {
	return &GetImageBadRequest{}
}

// WriteResponse to the client
func (o *GetImageBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
}

// GetImageUnauthorizedCode is the HTTP code returned for type GetImageUnauthorized
const GetImageUnauthorizedCode int = 401

/*GetImageUnauthorized Authentication information is missing or invalid

swagger:response getImageUnauthorized
*/
type GetImageUnauthorized struct {
	/*
	  Required: true
	*/
	WWWAuthenticate string `json:"WWW_Authenticate"`
}

// NewGetImageUnauthorized creates GetImageUnauthorized with default headers values
func NewGetImageUnauthorized() *GetImageUnauthorized {
	return &GetImageUnauthorized{}
}

// WithWWWAuthenticate adds the wWWAuthenticate to the get image unauthorized response
func (o *GetImageUnauthorized) WithWWWAuthenticate(wWWAuthenticate string) *GetImageUnauthorized {
	o.WWWAuthenticate = wWWAuthenticate
	return o
}

// SetWWWAuthenticate sets the wWWAuthenticate to the get image unauthorized response
func (o *GetImageUnauthorized) SetWWWAuthenticate(wWWAuthenticate string) {
	o.WWWAuthenticate = wWWAuthenticate
}

// WriteResponse to the client
func (o *GetImageUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header WWW_Authenticate

	wWWAuthenticate := o.WWWAuthenticate
	if wWWAuthenticate != "" {
		rw.Header().Set("WWW_Authenticate", wWWAuthenticate)
	}

	rw.WriteHeader(401)
}

// GetImageNotFoundCode is the HTTP code returned for type GetImageNotFound
const GetImageNotFoundCode int = 404

/*GetImageNotFound Image not found

swagger:response getImageNotFound
*/
type GetImageNotFound struct {
}

// NewGetImageNotFound creates GetImageNotFound with default headers values
func NewGetImageNotFound() *GetImageNotFound {
	return &GetImageNotFound{}
}

// WriteResponse to the client
func (o *GetImageNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}
