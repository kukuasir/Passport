// Code generated by go-swagger; DO NOT EDIT.

package passport

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"github.com/go-openapi/runtime/middleware"
	"Passport/models"
	"Passport/utils"
	"fmt"
	"time"
)

// NrPassportGetbackPwdHandlerFunc turns a function with the right signature into a passport getback pwd handler
type NrPassportGetbackPwdHandlerFunc func(NrPassportGetbackPwdParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrPassportGetbackPwdHandlerFunc) Handle(params NrPassportGetbackPwdParams) middleware.Responder {
	return fn(params)
}

// NrPassportGetbackPwdHandler interface for that can handle valid passport getback pwd params
type NrPassportGetbackPwdHandler interface {
	Handle(NrPassportGetbackPwdParams) middleware.Responder
}

// NewNrPassportGetbackPwd creates a new http.Handler for the passport getback pwd operation
func NewNrPassportGetbackPwd(ctx *middleware.Context, handler NrPassportGetbackPwdHandler) *NrPassportGetbackPwd {
	return &NrPassportGetbackPwd{Context: ctx, Handler: handler}
}

/*NrPassportGetbackPwd swagger:route POST /passport/getbackPwd Passport passportGetbackPwd

找回密码接口

找回密码接口

*/
type NrPassportGetbackPwd struct {
	Context *middleware.Context
	Handler NrPassportGetbackPwdHandler
}

func (o *NrPassportGetbackPwd) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrPassportGetbackPwdParams()

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

	var res models.RespState
	var state models.State

	var user models.UserBase

	// 定义错误信息
	var code int64
	var message string

	// 先判断验证码是否正确
	var sms SMSRecord
	db.Table(utils.T_SMS).Where("phone=?", *Params.Body.Phone).Where("code=?", *Params.Body.ValidCode).Where("type=3").Order("create_at DESC").First(&sms)
	if len(sms.Code) == 0 {
		code = 401
		message = "验证码不正确"
	} else {
		if time.Now().Unix() - utils.T_EXPIRED_SECONDS > sms.CreateAt {
			code = 402
			message = "验证码已失效"
		} else {

			db.Table(utils.T_USER).Where("phone=?", *Params.Body.Phone).Find(&user)

			// 用户ID不存在
			if user.ID == 0 {
				code = 403
				message = "手机号不存在"
			} else {
				sql := "UPDATE btk_User SET password = ? WHERE euid = ? AND status = 0"
				db.Exec(sql, utils.MD5Encrypt(*Params.Body.Password), user.Euid)
				code = 200
				message = "修改成功"
			}
		}
	}

	state.UnmarshalBinary([]byte(utils.Response200(code, message)))
	res.State = &state

	o.Context.Respond(rw, r, route.Produces, route, res)

}
