package session

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/location"

	"github.com/gin-gonic/gin"
)

var (
	jwtKey     = []byte("shoppe.xyz")
	ExpiresMin = 60 * 10
)

const (
	sidKey        = "sid"
	UserId        = "userId"
	BEARER_SCHEMA = "Bearer"
)

type Claims struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}

// 创建 session 中间件
func NewSessionStore(r *gin.Engine) {
	r.Use(location.Default())
	r.Use(cors.Default())
	r.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedPaths([]string{"/api/"})))

	r.Use(func(c *gin.Context) {
		isAuth := false
		authHeader := c.GetHeader("Authorization")
		var tokenString string
		if authHeader != "" {
			tokenString = authHeader[len(BEARER_SCHEMA):]
			isAuth = true
		}
		if tokenString == "" {
			tokenString, _ = c.Cookie(sidKey)
		}
		if tokenString == "" {
			return
		}
		claims := &Claims{}
		tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.Writer.WriteHeader(http.StatusUnauthorized)
				return
			}
			c.Writer.WriteHeader(http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			c.Writer.WriteHeader(http.StatusUnauthorized)
			return
		}
		c.Set(UserId, claims.UserId)
		// 如果小于10分钟，重新一次session
		if !isAuth && claims.ExpiresAt-time.Now().Unix() < int64(ExpiresMin) {
			// 重新刷新session
			SaveLoginSession(c, claims.UserId)
		}
	})

}

// 保存登录信息到session里
func SaveLoginSession(c *gin.Context, userId string) {
	// token 一年有效
	expirationTime := time.Now().Add(24 * time.Hour)
	sid, err := Sid(userId, expirationTime)
	if err != nil {
		log.Println(err)
		return
	}
	cookieConfig := &http.Cookie{
		Name:     sidKey,
		Path:     "/",
		HttpOnly: true,
		Value:    sid,
		Domain:   GetSubDomain(c),
		Expires:  expirationTime,
	}
	http.SetCookie(c.Writer, cookieConfig)
}

// 退出登录
func LogoutSession(c *gin.Context) {
	c.SetCookie(sidKey, "", -1, "/", GetSubDomain(c), false, true)
}

// 获取二级域名，比如 abc.x.com 得到的是 x.com
func GetSubDomain(c *gin.Context) string {
	url := location.Get(c)

	hostname := url.Hostname()

	if hostname == "localhost" || hostname == "127.0.0.1" {
		return hostname
	}

	s := strings.Split(url.Hostname(), ".")
	size := len(s)
	lastStr := s[size-1]
	lastStr2 := s[size-2]

	if isNum(lastStr) {
		return hostname
	}

	return lastStr2 + "." + lastStr
}

func isNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// 生成 sid
func Sid(userId string, expirationTime time.Time) (string, error) {
	claims := &Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}
	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}
