// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"go-swagger-hotels/models"
)

// GetAllHotelsOKCode is the HTTP code returned for type GetAllHotelsOK
const GetAllHotelsOKCode int = 200

/*GetAllHotelsOK OK

swagger:response getAllHotelsOK
*/
type GetAllHotelsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Hotel `json:"body,omitempty"`
}

// NewGetAllHotelsOK creates GetAllHotelsOK with default headers values
func NewGetAllHotelsOK() *GetAllHotelsOK {
	return &GetAllHotelsOK{}
}

// WithPayload adds the payload to the get all hotels o k response
func (o *GetAllHotelsOK) WithPayload(payload []*models.Hotel) *GetAllHotelsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all hotels o k response
func (o *GetAllHotelsOK) SetPayload(payload []*models.Hotel) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllHotelsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.Hotel, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*GetAllHotelsDefault error

swagger:response getAllHotelsDefault
*/
type GetAllHotelsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAllHotelsDefault creates GetAllHotelsDefault with default headers values
func NewGetAllHotelsDefault(code int) *GetAllHotelsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetAllHotelsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get all hotels default response
func (o *GetAllHotelsDefault) WithStatusCode(code int) *GetAllHotelsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get all hotels default response
func (o *GetAllHotelsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get all hotels default response
func (o *GetAllHotelsDefault) WithPayload(payload *models.Error) *GetAllHotelsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all hotels default response
func (o *GetAllHotelsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllHotelsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
