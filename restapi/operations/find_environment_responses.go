package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vice-registry/vice-api/models"
)

// FindEnvironmentOKCode is the HTTP code returned for type FindEnvironmentOK
const FindEnvironmentOKCode int = 200

/*FindEnvironmentOK An array of accessible execution environments.

swagger:response findEnvironmentOK
*/
type FindEnvironmentOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Environment `json:"body,omitempty"`
}

// NewFindEnvironmentOK creates FindEnvironmentOK with default headers values
func NewFindEnvironmentOK() *FindEnvironmentOK {
	return &FindEnvironmentOK{}
}

// WithPayload adds the payload to the find environment o k response
func (o *FindEnvironmentOK) WithPayload(payload []*models.Environment) *FindEnvironmentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find environment o k response
func (o *FindEnvironmentOK) SetPayload(payload []*models.Environment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindEnvironmentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.Environment, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// FindEnvironmentUnauthorizedCode is the HTTP code returned for type FindEnvironmentUnauthorized
const FindEnvironmentUnauthorizedCode int = 401

/*FindEnvironmentUnauthorized Authentication information is missing or invalid

swagger:response findEnvironmentUnauthorized
*/
type FindEnvironmentUnauthorized struct {
	/*
	  Required: true
	*/
	WWWAuthenticate string `json:"WWW_Authenticate"`
}

// NewFindEnvironmentUnauthorized creates FindEnvironmentUnauthorized with default headers values
func NewFindEnvironmentUnauthorized() *FindEnvironmentUnauthorized {
	return &FindEnvironmentUnauthorized{}
}

// WithWWWAuthenticate adds the wWWAuthenticate to the find environment unauthorized response
func (o *FindEnvironmentUnauthorized) WithWWWAuthenticate(wWWAuthenticate string) *FindEnvironmentUnauthorized {
	o.WWWAuthenticate = wWWAuthenticate
	return o
}

// SetWWWAuthenticate sets the wWWAuthenticate to the find environment unauthorized response
func (o *FindEnvironmentUnauthorized) SetWWWAuthenticate(wWWAuthenticate string) {
	o.WWWAuthenticate = wWWAuthenticate
}

// WriteResponse to the client
func (o *FindEnvironmentUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header WWW_Authenticate

	wWWAuthenticate := o.WWWAuthenticate
	if wWWAuthenticate != "" {
		rw.Header().Set("WWW_Authenticate", wWWAuthenticate)
	}

	rw.WriteHeader(401)
}

// FindEnvironmentInternalServerErrorCode is the HTTP code returned for type FindEnvironmentInternalServerError
const FindEnvironmentInternalServerErrorCode int = 500

/*FindEnvironmentInternalServerError Internal Server Error

swagger:response findEnvironmentInternalServerError
*/
type FindEnvironmentInternalServerError struct {
}

// NewFindEnvironmentInternalServerError creates FindEnvironmentInternalServerError with default headers values
func NewFindEnvironmentInternalServerError() *FindEnvironmentInternalServerError {
	return &FindEnvironmentInternalServerError{}
}

// WriteResponse to the client
func (o *FindEnvironmentInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
}
