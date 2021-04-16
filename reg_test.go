package suser

import (
	"context"
	"log"
	"testing"
	"github.com/shoppehub/suser/umod"
	"github.com/shoppehub/suser/umod/enttest"
	"github.com/shoppehub/suser/umod/migrate"
	"github.com/shoppehub/suser/umod/user"
	_ "github.com/mattn/go-sqlite3"
)

func TestReg(t *testing.T) {

	opts := []enttest.Option{
		enttest.WithOptions(umod.Log(t.Log)),
		enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)
	defer client.Close()
	userResource := UserResource{client}

	u := umod.User{
		Name: "123",
	}

	userResource.reg(&u)

	u2, _ := client.User.Query().Where(user.NameEQ("123")).Only(context.Background())

	log.Println("user returned: ", u2)
}
