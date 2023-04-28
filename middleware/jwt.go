package middleware

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const JWT_KEY = "JWT_PAYLOAD"

type User struct {
	Name string
	Role string
}

func unauthorized(c *gin.Context, code int, message string) {
	c.Header("WWW-Authenticate", "JWT realm=reward-shipping")
	c.Abort()
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

func CreateJWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := jwt.MapClaims{}

		parts := strings.Split(c.GetHeader("Authorization"), " ")
		tk := ""
		if len(parts) >= 2 {
			tk = parts[1]
		}
		if len(parts) == 1 {
			tk = parts[0]
		}

		if tk == "" {
			unauthorized(c, http.StatusUnauthorized, "authorization token required")
			return
		}

		jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("mock verification key"), nil
		})

		switch v := claims["exp"].(type) {
		case nil:
			unauthorized(c, http.StatusBadRequest, "missing exp field")
			return
		case float64:
			if int64(v) < time.Now().Unix() {
				unauthorized(c, http.StatusUnauthorized, "token is expired")
				return
			}
		case json.Number:
			n, err := v.Int64()
			if err != nil {
				unauthorized(c, http.StatusBadRequest, "exp must be float64 format")
				return
			}
			if n < time.Now().Unix() {
				unauthorized(c, http.StatusUnauthorized, "token is expired")
				return
			}
		default:
			unauthorized(c, http.StatusBadRequest, "exp must be float64 format")
			return
		}

		c.Set(JWT_KEY, claims)
		c.Next()
	}
}

func GetJWTClaims(c *gin.Context) (jwt.MapClaims, error) {
	data, exist := c.Get(JWT_KEY)
	if !exist {
		return nil, errors.New("jwt payload not exist")
	}
	claims, ok := data.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("wrong jwt payload format")
	}
	return claims, nil
}
