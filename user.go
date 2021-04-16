package suser

import (
	"context"
	"log"

	entsql "entgo.io/ent/dialect/sql"

	"github.com/shoppehub/suser/umod"
	"github.com/shoppehub/suser/umod/migrate"
)

type UserResource struct {
	Client *umod.Client
}

func New(drv *entsql.Driver) *UserResource {
	client := umod.NewClient(umod.Driver(drv))

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return &UserResource{client}
}
