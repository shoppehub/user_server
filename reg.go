package suser

import (
	"context"

	"github.com/shoppehub/suser/umod"
)

func (resource *UserResource) reg(user *umod.User) (*umod.User, error) {

	// var u *umod.User
	ctx := context.Background()
	u, err := resource.Client.User.Create().
		SetName(user.Name).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return u, nil
}
