// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	_"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"github.com/go-openapi/runtime/middleware"
	"Passport/models"
	"Passport/utils"
	"fmt"
)

// NrUserGetOwnerAccountHandlerFunc turns a function with the right signature into a user get owner account handler
type NrUserGetOwnerAccountHandlerFunc func(NrUserGetOwnerAccountParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrUserGetOwnerAccountHandlerFunc) Handle(params NrUserGetOwnerAccountParams) middleware.Responder {
	return fn(params)
}

// NrUserGetOwnerAccountHandler interface for that can handle valid user get owner account params
type NrUserGetOwnerAccountHandler interface {
	Handle(NrUserGetOwnerAccountParams) middleware.Responder
}

// NewNrUserGetOwnerAccount creates a new http.Handler for the user get owner account operation
func NewNrUserGetOwnerAccount(ctx *middleware.Context, handler NrUserGetOwnerAccountHandler) *NrUserGetOwnerAccount {
	return &NrUserGetOwnerAccount{Context: ctx, Handler: handler}
}

/*NrUserGetOwnerAccount swagger:route GET /user/getOwnerAccount User userGetOwnerAccount

查看用户账户信息

*/
type NrUserGetOwnerAccount struct {
	Context *middleware.Context
	Handler NrUserGetOwnerAccountHandler
}

func (o *NrUserGetOwnerAccount) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrUserGetOwnerAccountParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	//res := o.Handler.Handle(Params) // actually handle the request

	var res models.AccountState
	var state models.State

	/**
	  100💥
	*/
	if Params.Platform == 100 {
		state.UnmarshalBinary([]byte(utils.Response200(0, "查询成功")))
		res.State = &state
	} else {

		db, err := utils.OpenConnection()
		if err != nil {
			fmt.Println(err.Error())
		}
		defer db.Close()

		var data models.Account
		sql := "SELECT id, current_coins, current_points, current_rmb FROM btk_User WHERE id = ? AND status = 0"
		db.Raw(sql, utils.DecodeUserID(Params.Euid)).Find(&data)
		res.Data = &data
		state.UnmarshalBinary([]byte(utils.Response200(200, "查询成功")))
		res.State = &state
	}

	o.Context.Respond(rw, r, route.Produces, route, res)

}
