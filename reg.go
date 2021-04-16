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
)

func (resource *UserResource) reg(u *umod.User) (*umod.User, error) {

	hasField := false
	if u.Email != "" {
		hasField = true
		dbUser, err := resource.Client.User.Query().Where(user.EmailEQ(u.Email)).Only(context.Background())

		if err != nil {
			return nil, err
		}
		if dbUser != nil {
			return nil, errors.CodeError{Code: R6001, Msg: "email has already used"}
		}
	}

	if u.Username != "" {
		hasField = true
		dbUser, err := resource.Client.User.Query().Where(user.UsernameEQ(u.Username)).Only(context.Background())

		if err != nil {
			return nil, err
		}
		if dbUser != nil {
			return nil, errors.CodeError{Code: R6002, Msg: "username has already used"}
		}
	}

	if u.Mobile != "" {
		hasField = true
		dbUser, err := resource.Client.User.Query().Where(user.MobileEQ(u.Mobile)).Only(context.Background())

		if err != nil {
			return nil, err
		}
		if dbUser != nil {
			return nil, errors.CodeError{Code: R6003, Msg: "mobile has already used"}
		}
	}

	if !hasField {
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
