package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"omi-gitlab.e-technik.uni-ulm.de/vice/vice-api/models"
)

// GetEnvironmentOKCode is the HTTP code returned for type GetEnvironmentOK
const GetEnvironmentOKCode int = 200

/*GetEnvironmentOK successful operation

swagger:response getEnvironmentOK
*/
type GetEnvironmentOK struct {

	/*
	  In: Body
	*/
	Payload *models.Environment `json:"body,omitempty"`
}

// NewGetEnvironmentOK creates GetEnvironmentOK with default headers values
func NewGetEnvironmentOK() *GetEnvironmentOK {
	return &GetEnvironmentOK{}
}

// WithPayload adds the payload to the get environment o k response
func (o *GetEnvironmentOK) WithPayload(payload *models.Environment) *GetEnvironmentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get environment o k response
func (o *GetEnvironmentOK) SetPayload(payload *models.Environment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEnvironmentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetEnvironmentUnauthorizedCode is the HTTP code returned for type GetEnvironmentUnauthorized
const GetEnvironmentUnauthorizedCode int = 401

/*GetEnvironmentUnauthorized Authentication information is missing or invalid

swagger:response getEnvironmentUnauthorized
*/
type GetEnvironmentUnauthorized struct {
	/*
	  Required: true
	*/
	WWWAuthenticate string `json:"WWW_Authenticate"`
}

// NewGetEnvironmentUnauthorized creates GetEnvironmentUnauthorized with default headers values
func NewGetEnvironmentUnauthorized() *GetEnvironmentUnauthorized {
	return &GetEnvironmentUnauthorized{}
}

// WithWWWAuthenticate adds the wWWAuthenticate to the get environment unauthorized response
func (o *GetEnvironmentUnauthorized) WithWWWAuthenticate(wWWAuthenticate string) *GetEnvironmentUnauthorized {
	o.WWWAuthenticate = wWWAuthenticate
	return o
}

// SetWWWAuthenticate sets the wWWAuthenticate to the get environment unauthorized response
func (o *GetEnvironmentUnauthorized) SetWWWAuthenticate(wWWAuthenticate string) {
	o.WWWAuthenticate = wWWAuthenticate
}

// WriteResponse to the client
func (o *GetEnvironmentUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header WWW_Authenticate

	wWWAuthenticate := o.WWWAuthenticate
	if wWWAuthenticate != "" {
		rw.Header().Set("WWW_Authenticate", wWWAuthenticate)
	}

	rw.WriteHeader(401)
}

// GetEnvironmentNotFoundCode is the HTTP code returned for type GetEnvironmentNotFound
const GetEnvironmentNotFoundCode int = 404

/*GetEnvironmentNotFound Element not found

swagger:response getEnvironmentNotFound
*/
type GetEnvironmentNotFound struct {
}

// NewGetEnvironmentNotFound creates GetEnvironmentNotFound with default headers values
func NewGetEnvironmentNotFound() *GetEnvironmentNotFound {
	return &GetEnvironmentNotFound{}
}

// WriteResponse to the client
func (o *GetEnvironmentNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}

// GetEnvironmentInternalServerErrorCode is the HTTP code returned for type GetEnvironmentInternalServerError
const GetEnvironmentInternalServerErrorCode int = 500

/*GetEnvironmentInternalServerError Internal Server Error

swagger:response getEnvironmentInternalServerError
*/
type GetEnvironmentInternalServerError struct {
}

// NewGetEnvironmentInternalServerError creates GetEnvironmentInternalServerError with default headers values
func NewGetEnvironmentInternalServerError() *GetEnvironmentInternalServerError {
	return &GetEnvironmentInternalServerError{}
}

// WriteResponse to the client
func (o *GetEnvironmentInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
}