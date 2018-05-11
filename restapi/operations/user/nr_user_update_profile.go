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
	"strconv"
	"time"
)

// NrUserUpdateProfileHandlerFunc turns a function with the right signature into a user update profile handler
type NrUserUpdateProfileHandlerFunc func(NrUserUpdateProfileParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrUserUpdateProfileHandlerFunc) Handle(params NrUserUpdateProfileParams) middleware.Responder {
	return fn(params)
}

// NrUserUpdateProfileHandler interface for that can handle valid user update profile params
type NrUserUpdateProfileHandler interface {
	Handle(NrUserUpdateProfileParams) middleware.Responder
}

// NewNrUserUpdateProfile creates a new http.Handler for the user update profile operation
func NewNrUserUpdateProfile(ctx *middleware.Context, handler NrUserUpdateProfileHandler) *NrUserUpdateProfile {
	return &NrUserUpdateProfile{Context: ctx, Handler: handler}
}

/*NrUserUpdateProfile swagger:route POST /user/updateProfile User userUpdateProfile

修改用户资料接口

修改用户资料接口

*/
type NrUserUpdateProfile struct {
	Context *middleware.Context
	Handler NrUserUpdateProfileHandler
}

func (o *NrUserUpdateProfile) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrUserUpdateProfileParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	//res := o.Handler.Handle(Params) // actually handle the request

	db, err := utils.OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	var res models.UserState
	var state models.State
	var data models.UserInfo

	db.Table(utils.T_USER).Where("euid=?", Params.Body.Euid).Where("status=0").Find(&data)

	// 不存在的用户
	if data.Euid == nil {
		state.UnmarshalBinary([]byte(utils.Response200(402, "用户不存在")))
		res.State = &state
		o.Context.Respond(rw, r, route.Produces, route, res)
		return
	}

	if len(Params.Body.BirthDay) > 0 {
		data.BirthDay = Params.Body.BirthDay
	}
	if len(Params.Body.Blood) > 0 {
		data.Blood = Params.Body.Blood
	}
	if len(Params.Body.Degree) > 0 {
		data.Degree = Params.Body.Degree
	}
	if Params.Body.Gender > 0 {
		data.Gender = Params.Body.Gender
	}
	if len(Params.Body.HomeArea) > 0 {
		data.HomeArea = Params.Body.HomeArea
	}
	if len(Params.Body.Interest) > 0 {
		data.Interest = Params.Body.Interest
	}
	if len(Params.Body.Marriage) > 0 {
		data.Marriage = Params.Body.Marriage
	}
	if len(Params.Body.NickName) > 0 {
		data.NickName = Params.Body.NickName
	}
	if len(Params.Body.NowArea) > 0 {
		data.NowArea = Params.Body.NowArea
	}
	if len(Params.Body.Profession) > 0 {
		data.Profession = Params.Body.Profession
	}
	if len(Params.Body.Resume) > 0 {
		data.Resume = Params.Body.Resume
	}
	if len(Params.Body.Salary) > 0 {
		data.Salary = Params.Body.Salary
	}
	if len(Params.Body.School) > 0 {
		data.School = Params.Body.School
	}
	if len(Params.Body.Shape) > 0 {
		data.Shape = Params.Body.Shape
	}
	if Params.Body.Stature > 0 {
		data.Stature = strconv.FormatInt(Params.Body.Stature, 10) + "cm"
	}
	if Params.Body.Weight > 0 {
		data.Weight = strconv.FormatInt(Params.Body.Weight, 10) + "cm"
	}

	data.UpdateAt = time.Now().Unix()

	db.Table(utils.T_USER).Save(&data)

	state.UnmarshalBinary([]byte(utils.Response200(200, "修改成功")))
	res.State = &state
	res.Data = &data

	o.Context.Respond(rw, r, route.Produces, route, res)

}
