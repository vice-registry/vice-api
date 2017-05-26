package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"omi-gitlab.e-technik.uni-ulm.de/vice/vice-api/models"
)

// CreateEnvironmentCreatedCode is the HTTP code returned for type CreateEnvironmentCreated
const CreateEnvironmentCreatedCode int = 201

/*CreateEnvironmentCreated Created

swagger:response createEnvironmentCreated
*/
type CreateEnvironmentCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Environment `json:"body,omitempty"`
}

// NewCreateEnvironmentCreated creates CreateEnvironmentCreated with default headers values
func NewCreateEnvironmentCreated() *CreateEnvironmentCreated {
	return &CreateEnvironmentCreated{}
}

// WithPayload adds the payload to the create environment created response
func (o *CreateEnvironmentCreated) WithPayload(payload *models.Environment) *CreateEnvironmentCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create environment created response
func (o *CreateEnvironmentCreated) SetPayload(payload *models.Environment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateEnvironmentCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateEnvironmentUnauthorizedCode is the HTTP code returned for type CreateEnvironmentUnauthorized
const CreateEnvironmentUnauthorizedCode int = 401

/*CreateEnvironmentUnauthorized Authentication information is missing or invalid

swagger:response createEnvironmentUnauthorized
*/
type CreateEnvironmentUnauthorized struct {
	/*
	  Required: true
	*/
	WWWAuthenticate string `json:"WWW_Authenticate"`
}

// NewCreateEnvironmentUnauthorized creates CreateEnvironmentUnauthorized with default headers values
func NewCreateEnvironmentUnauthorized() *CreateEnvironmentUnauthorized {
	return &CreateEnvironmentUnauthorized{}
}

// WithWWWAuthenticate adds the wWWAuthenticate to the create environment unauthorized response
func (o *CreateEnvironmentUnauthorized) WithWWWAuthenticate(wWWAuthenticate string) *CreateEnvironmentUnauthorized {
	o.WWWAuthenticate = wWWAuthenticate
	return o
}

// SetWWWAuthenticate sets the wWWAuthenticate to the create environment unauthorized response
func (o *CreateEnvironmentUnauthorized) SetWWWAuthenticate(wWWAuthenticate string) {
	o.WWWAuthenticate = wWWAuthenticate
}

// WriteResponse to the client
func (o *CreateEnvironmentUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header WWW_Authenticate

	wWWAuthenticate := o.WWWAuthenticate
	if wWWAuthenticate != "" {
		rw.Header().Set("WWW_Authenticate", wWWAuthenticate)
	}

	rw.WriteHeader(401)
}

// CreateEnvironmentMethodNotAllowedCode is the HTTP code returned for type CreateEnvironmentMethodNotAllowed
const CreateEnvironmentMethodNotAllowedCode int = 405

/*CreateEnvironmentMethodNotAllowed Invalid input

swagger:response createEnvironmentMethodNotAllowed
*/
type CreateEnvironmentMethodNotAllowed struct {
}

// NewCreateEnvironmentMethodNotAllowed creates CreateEnvironmentMethodNotAllowed with default headers values
func NewCreateEnvironmentMethodNotAllowed() *CreateEnvironmentMethodNotAllowed {
	return &CreateEnvironmentMethodNotAllowed{}
}

// WriteResponse to the client
func (o *CreateEnvironmentMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
}
