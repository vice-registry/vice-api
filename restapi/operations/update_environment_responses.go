// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vice-registry/vice-util/models"
)

// UpdateEnvironmentCreatedCode is the HTTP code returned for type UpdateEnvironmentCreated
const UpdateEnvironmentCreatedCode int = 201

/*UpdateEnvironmentCreated Updated

swagger:response updateEnvironmentCreated
*/
type UpdateEnvironmentCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Environment `json:"body,omitempty"`
}

// NewUpdateEnvironmentCreated creates UpdateEnvironmentCreated with default headers values
func NewUpdateEnvironmentCreated() *UpdateEnvironmentCreated {
	return &UpdateEnvironmentCreated{}
}

// WithPayload adds the payload to the update environment created response
func (o *UpdateEnvironmentCreated) WithPayload(payload *models.Environment) *UpdateEnvironmentCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update environment created response
func (o *UpdateEnvironmentCreated) SetPayload(payload *models.Environment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateEnvironmentCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateEnvironmentUnauthorizedCode is the HTTP code returned for type UpdateEnvironmentUnauthorized
const UpdateEnvironmentUnauthorizedCode int = 401

/*UpdateEnvironmentUnauthorized Authentication information is missing or invalid

swagger:response updateEnvironmentUnauthorized
*/
type UpdateEnvironmentUnauthorized struct {
	/*
	  Required: true
	*/
	WWWAuthenticate string `json:"WWW_Authenticate"`
}

// NewUpdateEnvironmentUnauthorized creates UpdateEnvironmentUnauthorized with default headers values
func NewUpdateEnvironmentUnauthorized() *UpdateEnvironmentUnauthorized {
	return &UpdateEnvironmentUnauthorized{}
}

// WithWWWAuthenticate adds the wWWAuthenticate to the update environment unauthorized response
func (o *UpdateEnvironmentUnauthorized) WithWWWAuthenticate(wWWAuthenticate string) *UpdateEnvironmentUnauthorized {
	o.WWWAuthenticate = wWWAuthenticate
	return o
}

// SetWWWAuthenticate sets the wWWAuthenticate to the update environment unauthorized response
func (o *UpdateEnvironmentUnauthorized) SetWWWAuthenticate(wWWAuthenticate string) {
	o.WWWAuthenticate = wWWAuthenticate
}

// WriteResponse to the client
func (o *UpdateEnvironmentUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header WWW_Authenticate

	wWWAuthenticate := o.WWWAuthenticate
	if wWWAuthenticate != "" {
		rw.Header().Set("WWW_Authenticate", wWWAuthenticate)
	}

	rw.WriteHeader(401)
}

// UpdateEnvironmentNotFoundCode is the HTTP code returned for type UpdateEnvironmentNotFound
const UpdateEnvironmentNotFoundCode int = 404

/*UpdateEnvironmentNotFound Element not found

swagger:response updateEnvironmentNotFound
*/
type UpdateEnvironmentNotFound struct {
}

// NewUpdateEnvironmentNotFound creates UpdateEnvironmentNotFound with default headers values
func NewUpdateEnvironmentNotFound() *UpdateEnvironmentNotFound {
	return &UpdateEnvironmentNotFound{}
}

// WriteResponse to the client
func (o *UpdateEnvironmentNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}

// UpdateEnvironmentMethodNotAllowedCode is the HTTP code returned for type UpdateEnvironmentMethodNotAllowed
const UpdateEnvironmentMethodNotAllowedCode int = 405

/*UpdateEnvironmentMethodNotAllowed Invalid input

swagger:response updateEnvironmentMethodNotAllowed
*/
type UpdateEnvironmentMethodNotAllowed struct {
}

// NewUpdateEnvironmentMethodNotAllowed creates UpdateEnvironmentMethodNotAllowed with default headers values
func NewUpdateEnvironmentMethodNotAllowed() *UpdateEnvironmentMethodNotAllowed {
	return &UpdateEnvironmentMethodNotAllowed{}
}

// WriteResponse to the client
func (o *UpdateEnvironmentMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
}

// UpdateEnvironmentInternalServerErrorCode is the HTTP code returned for type UpdateEnvironmentInternalServerError
const UpdateEnvironmentInternalServerErrorCode int = 500

/*UpdateEnvironmentInternalServerError Internal Server Error

swagger:response updateEnvironmentInternalServerError
*/
type UpdateEnvironmentInternalServerError struct {
}

// NewUpdateEnvironmentInternalServerError creates UpdateEnvironmentInternalServerError with default headers values
func NewUpdateEnvironmentInternalServerError() *UpdateEnvironmentInternalServerError {
	return &UpdateEnvironmentInternalServerError{}
}

// WriteResponse to the client
func (o *UpdateEnvironmentInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
}
