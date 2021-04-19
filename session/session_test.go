package session

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSaveLoginSession(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	NewSessionStore(r)

	r.GET("/get", func(c *gin.Context) {

		fmt.Println(444, "get", c.GetString(UserId))

		c.String(200, c.GetString(UserId))

	})

	r.GET("/save", func(c *gin.Context) {

		userId := "1234567"

		SaveLoginSession(c, userId)

		c.String(200, "ok")

	})

	r.GET("/logout", func(c *gin.Context) {

		LogoutSession(c)

		c.String(200, "ok")

	})

	res1 := httptest.NewRecorder()
	req1, _ := http.NewRequest("GET", "/save", nil)
	req1.Host = "pp.youkeda.com"
	r.ServeHTTP(res1, req1)

	res2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("GET", "/get", nil)
	req2.Host = "pp.youkeda.com"

	req2.Header.Set("Cookie", res1.Header().Get("Set-Cookie"))
	r.ServeHTTP(res2, req2)

	res3 := httptest.NewRecorder()
	req3, _ := http.NewRequest("GET", "/logout", nil)
	req3.Host = "pp.youkeda.com"

	req3.Header.Set("Cookie", res1.Header().Get("Set-Cookie"))
	r.ServeHTTP(res3, req3)

	fmt.Println(res3.Header())

}
