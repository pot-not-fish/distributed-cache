package mw

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var (
	secretKey = []byte("ZhaZrDBQb7MYdJWaPf5gJmGbYyVjLYgz")
)

func GenerateJWT() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"iss": "cache",
		"sub": "admin",
	})
	return token.SignedString(secretKey)
}

func ParseToken(token string) error {
	if token == "" {
		return fmt.Errorf("nil token")
	}

	jwt, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return err
	}

	if !jwt.Valid {
		return fmt.Errorf("invalid token")
	}
	return nil
}

func JWTMiddlewareFunc(c *gin.Context) {
	token := c.PostForm("token")

	if err := ParseToken(token); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": err.Error()})
		return
	}
	c.Next()
}
