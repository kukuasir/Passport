// Code generated by go-swagger; DO NOT EDIT.

package file_upload

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"Passport/models"
)

// FileUploadOKCode is the HTTP code returned for type FileUploadOK
const FileUploadOKCode int = 200

/*FileUploadOK 成功，返回成功信息

swagger:response fileUploadOK
*/
type FileUploadOK struct {

	/*
	  In: Body
	*/
	Payload *models.FileUploadOKBody `json:"body,omitempty"`
}

// NewFileUploadOK creates FileUploadOK with default headers values
func NewFileUploadOK() *FileUploadOK {
	return &FileUploadOK{}
}

// WithPayload adds the payload to the file upload o k response
func (o *FileUploadOK) WithPayload(payload *models.FileUploadOKBody) *FileUploadOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the file upload o k response
func (o *FileUploadOK) SetPayload(payload *models.FileUploadOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FileUploadOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*NrFileUploadDefault 返回错误

swagger:response fileUploadDefault
*/
type NrFileUploadDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewNrFileUploadDefault creates NrFileUploadDefault with default headers values
func NewNrFileUploadDefault(code int) *NrFileUploadDefault {
	if code <= 0 {
		code = 500
	}

	return &NrFileUploadDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the file upload default response
func (o *NrFileUploadDefault) WithStatusCode(code int) *NrFileUploadDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the file upload default response
func (o *NrFileUploadDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the file upload default response
func (o *NrFileUploadDefault) WithPayload(payload *models.Error) *NrFileUploadDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the file upload default response
func (o *NrFileUploadDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NrFileUploadDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
