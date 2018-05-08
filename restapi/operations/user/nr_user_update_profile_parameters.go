// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"Passport/models"
)

// NewNrUserUpdateProfileParams creates a new NrUserUpdateProfileParams object
// with the default values initialized.
func NewNrUserUpdateProfileParams() NrUserUpdateProfileParams {
	var ()
	return NrUserUpdateProfileParams{}
}

// NrUserUpdateProfileParams contains all the bound params for the user update profile operation
// typically these are obtained from a http.Request
//
// swagger:parameters /user/updateProfile
type NrUserUpdateProfileParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*修改用户资料JSON
	  In: body
	*/
	Body *models.UserUpdateProfileParamsBody
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *NrUserUpdateProfileParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.UserUpdateProfileParamsBody
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("body", "body", "", err))
		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
