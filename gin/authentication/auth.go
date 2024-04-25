package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var tokens []string

func main() {
	r := gin.Default()
	r.GET("/basic-auth", gin.BasicAuth(gin.Accounts{
		"admin": "secret",
	}), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "basic auth success",
		})
	})

	r.POST("login", gin.BasicAuth(gin.Accounts{
		"admin": "secret",
	}), func(c *gin.Context) {
		token, _ := randomHex(20)
		tokens = append(tokens, token)
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	})

	r.GET("bearer-auth", func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("Authorization")
		reqToken := strings.Split(bearerToken, " ")[1]
		fmt.Println("tokens => ", tokens)
		fmt.Println("reqToken => ", reqToken)
		for _, token := range tokens {
			if token == reqToken {
				c.JSON(http.StatusOK, gin.H{
					"data": "bearer auth token success",
				})
				return
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
	})
	log.Fatal(r.Run()) // default port :8080
}

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
