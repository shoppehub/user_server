package suser

import (
	"log"
	"testing"
	"time"

	"entgo.io/ent/dialect/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/shoppehub/suser/errors"
	"github.com/shoppehub/suser/umod"
)

func InitTest() *UserResource {
	drv, err := sql.Open("sqlite3", "file:ent?_fk=1")
	if err != nil {

		log.Fatal(err)
	}
	// Get the underlying sql.DB object of the driver.
	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	return New(drv)
}

func TestReg(t *testing.T) {

	// opts := []enttest.Option{
	// 	enttest.WithOptions(umod.Log(t.Log)),
	// 	enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
	// }
	// urlConfig := "file:ent?mode=memory&cache=shared&_fk=1"
	// urlConfig = "file:ent?_fk=1"
	// client := enttest.Open(t, "sqlite3", urlConfig, opts...)
	// defer client.Close()

	userResource := InitTest()

	u := &umod.User{
		Name:     "123",
		NickName: "123",
	}

	_, err := userResource.reg(u)

	if err != nil {
		log.Println(err)
		if errors.IsCodeError(err) {
			// log.Println(err)
			codeerr, _ := err.(errors.CodeError)
			if codeerr.Code == R6004 {
				return
			}

		}
		t.Fail()
	}

}
