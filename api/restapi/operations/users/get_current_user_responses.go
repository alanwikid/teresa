package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/luizalabs/paas/api/models"
)

/*GetCurrentUserOK User info

swagger:response getCurrentUserOK
*/
type GetCurrentUserOK struct {

	// In: body
	Payload *models.User `json:"body,omitempty"`
}

// NewGetCurrentUserOK creates GetCurrentUserOK with default headers values
func NewGetCurrentUserOK() *GetCurrentUserOK {
	return &GetCurrentUserOK{}
}

// WithPayload adds the payload to the get current user o k response
func (o *GetCurrentUserOK) WithPayload(payload *models.User) *GetCurrentUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get current user o k response
func (o *GetCurrentUserOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCurrentUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetCurrentUserNotFound Resource not found

swagger:response getCurrentUserNotFound
*/
type GetCurrentUserNotFound struct {

	// In: body
	Payload *models.NotFound `json:"body,omitempty"`
}

// NewGetCurrentUserNotFound creates GetCurrentUserNotFound with default headers values
func NewGetCurrentUserNotFound() *GetCurrentUserNotFound {
	return &GetCurrentUserNotFound{}
}

// WithPayload adds the payload to the get current user not found response
func (o *GetCurrentUserNotFound) WithPayload(payload *models.NotFound) *GetCurrentUserNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get current user not found response
func (o *GetCurrentUserNotFound) SetPayload(payload *models.NotFound) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCurrentUserNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetCurrentUserDefault Error

swagger:response getCurrentUserDefault
*/
type GetCurrentUserDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetCurrentUserDefault creates GetCurrentUserDefault with default headers values
func NewGetCurrentUserDefault(code int) *GetCurrentUserDefault {
	if code <= 0 {
		code = 500
	}

	return &GetCurrentUserDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get current user default response
func (o *GetCurrentUserDefault) WithStatusCode(code int) *GetCurrentUserDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get current user default response
func (o *GetCurrentUserDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get current user default response
func (o *GetCurrentUserDefault) WithPayload(payload *models.Error) *GetCurrentUserDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get current user default response
func (o *GetCurrentUserDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCurrentUserDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}