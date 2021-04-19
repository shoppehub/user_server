package suser

import (
	"context"

	"github.com/shoppehub/suser/errors"

	"github.com/shoppehub/suser/umod"
	"github.com/shoppehub/suser/umod/user"
)

const (
	R6001 = 6001
	R6002 = 6002
	R6003 = 6003
	R6004 = 6004
	R6005 = 6005
	// 邮箱查询错误
	R6006 = 6006
	// 用户名查询错误
	R6007 = 6007
	// 手机号查询错误
	R6008 = 6008
)

func (resource *UserResource) reg(u *umod.User) (*umod.User, error) {

	// 邮箱、用户名、手机号，必须要有一个
	hasLoginNameField := false

	if u.Email != "" {
		hasLoginNameField = true
		dbUser, err := resource.Client.User.Query().Where(user.EmailEQ(u.Email)).Only(context.Background())

		if err != nil {
			return nil, errors.CodeError{Code: R6006, Msg: err.Error()}
		}
		if dbUser != nil {
			return nil, errors.CodeError{Code: R6001, Msg: "email has already used"}
		}
	}

	if u.Username != "" {
		hasLoginNameField = true
		dbUser, err := resource.Client.User.Query().Where(user.UsernameEQ(u.Username)).Only(context.Background())

		if err != nil {
			return nil, errors.CodeError{Code: R6007, Msg: err.Error()}
		}
		if dbUser != nil {
			return nil, errors.CodeError{Code: R6002, Msg: "username has already used"}
		}
	}

	if u.Mobile != "" {
		hasLoginNameField = true
		dbUser, err := resource.Client.User.Query().Where(user.MobileEQ(u.Mobile)).Only(context.Background())

		if err != nil {
			return nil, errors.CodeError{Code: R6008, Msg: err.Error()}
		}
		if dbUser != nil {
			return nil, errors.CodeError{Code: R6003, Msg: "mobile has already used"}
		}
	}

	if !hasLoginNameField {
		return nil, errors.CodeError{Code: R6004, Msg: "mobile|email|username must be one"}
	}

	// var u *umod.User
	u, err := resource.Client.User.Create().
		SetName(u.Name).
		SetNickName(u.NickName).
		Save(context.Background())
	if err != nil {
		return nil, errors.CodeError{Code: R6005, Msg: err.Error()}
	}
	return u, nil
}
