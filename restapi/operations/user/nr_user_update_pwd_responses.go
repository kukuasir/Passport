// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"Passport/models"
)

// UserUpdatePwdOKCode is the HTTP code returned for type UserUpdatePwdOK
const UserUpdatePwdOKCode int = 200

/*UserUpdatePwdOK 修改成功，返回成功信息

swagger:response userUpdatePwdOK
*/
type UserUpdatePwdOK struct {

	/*
	  In: Body
	*/
	Payload *models.State `json:"body,omitempty"`
}

// NewUserUpdatePwdOK creates UserUpdatePwdOK with default headers values
func NewUserUpdatePwdOK() *UserUpdatePwdOK {
	return &UserUpdatePwdOK{}
}

// WithPayload adds the payload to the user update pwd o k response
func (o *UserUpdatePwdOK) WithPayload(payload *models.State) *UserUpdatePwdOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user update pwd o k response
func (o *UserUpdatePwdOK) SetPayload(payload *models.State) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserUpdatePwdOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*NrUserUpdatePwdDefault error

swagger:response userUpdatePwdDefault
*/
type NrUserUpdatePwdDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewNrUserUpdatePwdDefault creates NrUserUpdatePwdDefault with default headers values
func NewNrUserUpdatePwdDefault(code int) *NrUserUpdatePwdDefault {
	if code <= 0 {
		code = 500
	}

	return &NrUserUpdatePwdDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the user update pwd default response
func (o *NrUserUpdatePwdDefault) WithStatusCode(code int) *NrUserUpdatePwdDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the user update pwd default response
func (o *NrUserUpdatePwdDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the user update pwd default response
func (o *NrUserUpdatePwdDefault) WithPayload(payload *models.Error) *NrUserUpdatePwdDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user update pwd default response
func (o *NrUserUpdatePwdDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NrUserUpdatePwdDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
