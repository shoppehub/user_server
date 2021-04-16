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

func initTest() *UserResource {
	urlConfig := "file:ent?mode=memory&cache=shared&_fk=1"

	drv, err := sql.Open("sqlite3", urlConfig)
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

	userResource := initTest()

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
