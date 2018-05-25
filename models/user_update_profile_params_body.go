// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserUpdateProfileParamsBody user update profile params body
// swagger:model userUpdateProfileParamsBody
type UserUpdateProfileParamsBody struct {

	// 生日(格式必须为:yyyy-MM-dd)
	BirthDay string `json:"birth_day,omitempty"`

	// 血型
	Blood string `json:"blood,omitempty"`

	// 最高学历
	Degree string `json:"degree,omitempty"`

	// 用户ID
	// Required: true
	Euid *string `json:"euid"`

	// 用户性别(0:保密 1:男 2:女)
	// Required: true
	Gender *int64 `json:"gender"`

	// 家乡
	HomeArea string `json:"home_area,omitempty"`

	// 兴趣爱好
	Interest string `json:"interest,omitempty"`

	// 婚姻状况
	Marriage string `json:"marriage,omitempty"`

	// 用户昵称
	NickName string `json:"nick_name,omitempty"`

	// 现居城市
	NowArea string `json:"now_area,omitempty"`

	// 职业信息
	Profession string `json:"profession,omitempty"`

	// 个人说明
	Resume string `json:"resume,omitempty"`

	// 目前月薪
	Salary string `json:"salary,omitempty"`

	// 毕业学校
	School string `json:"school,omitempty"`

	// 体型
	Shape string `json:"shape,omitempty"`

	// 身高
	Stature int64 `json:"stature,omitempty"`

	// 体重
	Weight int64 `json:"weight,omitempty"`
}

// Validate validates this user update profile params body
func (m *UserUpdateProfileParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEuid(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGender(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserUpdateProfileParamsBody) validateEuid(formats strfmt.Registry) error {

	if err := validate.Required("euid", "body", m.Euid); err != nil {
		return err
	}

	return nil
}

func (m *UserUpdateProfileParamsBody) validateGender(formats strfmt.Registry) error {

	/*
	if err := validate.Required("gender", "body", m.Gender); err != nil {
		return err
	}
	*/

	return nil
}

// MarshalBinary interface implementation
func (m *UserUpdateProfileParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserUpdateProfileParamsBody) UnmarshalBinary(b []byte) error {
	var res UserUpdateProfileParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
