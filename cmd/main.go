package main

import (
	"fmt"
	"strings"

	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"github.com/shoppehub/suser/session"
)

const ok = "ok"

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	session.NewSessionStore(r)

	r.GET("/get", func(c *gin.Context) {

		sid, _ := c.Cookie("sid")

		c.String(200, sid)
	})

	r.GET("/save", func(c *gin.Context) {
		url := location.Get(c)

		user := "123"
		session.SaveLoginSession(c, user)

		s := strings.Split(url.Hostname(), ".")
		c.String(200, s[0]+"  "+fmt.Sprint(len(s)))
	})
	r.GET("/logout", func(c *gin.Context) {

		session.LogoutSession(c)

		c.String(200, "ok")

	})
	r.Run(":5000")
}
