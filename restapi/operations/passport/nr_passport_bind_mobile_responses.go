// Code generated by go-swagger; DO NOT EDIT.

package passport

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"Passport/models"
)

// PassportBindMobileOKCode is the HTTP code returned for type PassportBindMobileOK
const PassportBindMobileOKCode int = 200

/*PassportBindMobileOK 绑定成功，返回成功信息

swagger:response passportBindMobileOK
*/
type PassportBindMobileOK struct {

	/*
	  In: Body
	*/
	Payload *models.State `json:"body,omitempty"`
}

// NewPassportBindMobileOK creates PassportBindMobileOK with default headers values
func NewPassportBindMobileOK() *PassportBindMobileOK {
	return &PassportBindMobileOK{}
}

// WithPayload adds the payload to the passport bind mobile o k response
func (o *PassportBindMobileOK) WithPayload(payload *models.State) *PassportBindMobileOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the passport bind mobile o k response
func (o *PassportBindMobileOK) SetPayload(payload *models.State) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PassportBindMobileOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*NrPassportBindMobileDefault error

swagger:response passportBindMobileDefault
*/
type NrPassportBindMobileDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewNrPassportBindMobileDefault creates NrPassportBindMobileDefault with default headers values
func NewNrPassportBindMobileDefault(code int) *NrPassportBindMobileDefault {
	if code <= 0 {
		code = 500
	}

	return &NrPassportBindMobileDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the passport bind mobile default response
func (o *NrPassportBindMobileDefault) WithStatusCode(code int) *NrPassportBindMobileDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the passport bind mobile default response
func (o *NrPassportBindMobileDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the passport bind mobile default response
func (o *NrPassportBindMobileDefault) WithPayload(payload *models.Error) *NrPassportBindMobileDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the passport bind mobile default response
func (o *NrPassportBindMobileDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NrPassportBindMobileDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
