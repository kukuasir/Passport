// Code generated by go-swagger; DO NOT EDIT.

package passport

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"Passport/models"
)

// NewNrPassportQuickLoginParams creates a new NrPassportQuickLoginParams object
// with the default values initialized.
func NewNrPassportQuickLoginParams() NrPassportQuickLoginParams {
	var ()
	return NrPassportQuickLoginParams{}
}

// NrPassportQuickLoginParams contains all the bound params for the passport quick login operation
// typically these are obtained from a http.Request
//
// swagger:parameters /passport/quickLogin
type NrPassportQuickLoginParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*快捷登录JSON
	  In: body
	*/
	Body *models.PassportQuickLoginParamsBody
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *NrPassportQuickLoginParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.PassportQuickLoginParamsBody
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
