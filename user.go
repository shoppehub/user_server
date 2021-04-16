package suser

import (
	entsql "entgo.io/ent/dialect/sql"

	"github.com/shoppehub/suser/umod"
)

type UserResource struct {
	Client *umod.Client
}

func New(drv *entsql.Driver) *UserResource {
	client := umod.NewClient(umod.Driver(drv))

	return &UserResource{client}
}
