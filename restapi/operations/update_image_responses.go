// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vice-registry/vice-api/models"
)

// UpdateImageCreatedCode is the HTTP code returned for type UpdateImageCreated
const UpdateImageCreatedCode int = 201

/*UpdateImageCreated Updated

swagger:response updateImageCreated
*/
type UpdateImageCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Image `json:"body,omitempty"`
}

// NewUpdateImageCreated creates UpdateImageCreated with default headers values
func NewUpdateImageCreated() *UpdateImageCreated {
	return &UpdateImageCreated{}
}

// WithPayload adds the payload to the update image created response
func (o *UpdateImageCreated) WithPayload(payload *models.Image) *UpdateImageCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update image created response
func (o *UpdateImageCreated) SetPayload(payload *models.Image) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateImageCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateImageUnauthorizedCode is the HTTP code returned for type UpdateImageUnauthorized
const UpdateImageUnauthorizedCode int = 401

/*UpdateImageUnauthorized Authentication information is missing or invalid

swagger:response updateImageUnauthorized
*/
type UpdateImageUnauthorized struct {
	/*
	  Required: true
	*/
	WWWAuthenticate string `json:"WWW_Authenticate"`
}

// NewUpdateImageUnauthorized creates UpdateImageUnauthorized with default headers values
func NewUpdateImageUnauthorized() *UpdateImageUnauthorized {
	return &UpdateImageUnauthorized{}
}

// WithWWWAuthenticate adds the wWWAuthenticate to the update image unauthorized response
func (o *UpdateImageUnauthorized) WithWWWAuthenticate(wWWAuthenticate string) *UpdateImageUnauthorized {
	o.WWWAuthenticate = wWWAuthenticate
	return o
}

// SetWWWAuthenticate sets the wWWAuthenticate to the update image unauthorized response
func (o *UpdateImageUnauthorized) SetWWWAuthenticate(wWWAuthenticate string) {
	o.WWWAuthenticate = wWWAuthenticate
}

// WriteResponse to the client
func (o *UpdateImageUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header WWW_Authenticate

	wWWAuthenticate := o.WWWAuthenticate
	if wWWAuthenticate != "" {
		rw.Header().Set("WWW_Authenticate", wWWAuthenticate)
	}

	rw.WriteHeader(401)
}

// UpdateImageNotFoundCode is the HTTP code returned for type UpdateImageNotFound
const UpdateImageNotFoundCode int = 404

/*UpdateImageNotFound Element not found

swagger:response updateImageNotFound
*/
type UpdateImageNotFound struct {
}

// NewUpdateImageNotFound creates UpdateImageNotFound with default headers values
func NewUpdateImageNotFound() *UpdateImageNotFound {
	return &UpdateImageNotFound{}
}

// WriteResponse to the client
func (o *UpdateImageNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}

// UpdateImageMethodNotAllowedCode is the HTTP code returned for type UpdateImageMethodNotAllowed
const UpdateImageMethodNotAllowedCode int = 405

/*UpdateImageMethodNotAllowed Invalid input

swagger:response updateImageMethodNotAllowed
*/
type UpdateImageMethodNotAllowed struct {
}

// NewUpdateImageMethodNotAllowed creates UpdateImageMethodNotAllowed with default headers values
func NewUpdateImageMethodNotAllowed() *UpdateImageMethodNotAllowed {
	return &UpdateImageMethodNotAllowed{}
}

// WriteResponse to the client
func (o *UpdateImageMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
}

// UpdateImageInternalServerErrorCode is the HTTP code returned for type UpdateImageInternalServerError
const UpdateImageInternalServerErrorCode int = 500

/*UpdateImageInternalServerError Internal Server Error

swagger:response updateImageInternalServerError
*/
type UpdateImageInternalServerError struct {
}

// NewUpdateImageInternalServerError creates UpdateImageInternalServerError with default headers values
func NewUpdateImageInternalServerError() *UpdateImageInternalServerError {
	return &UpdateImageInternalServerError{}
}

// WriteResponse to the client
func (o *UpdateImageInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
}
