package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("my_secret_key")
var tokens []string

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func main() {
	r := gin.Default()
	r.GET("/basic-auth", gin.BasicAuth(gin.Accounts{
		"admin": "secret",
	}), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "basic auth success",
		})
	})

	r.POST("bearer-login", gin.BasicAuth(gin.Accounts{
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

	r.POST("jwt-login", gin.BasicAuth(gin.Accounts{
		"admin": "secret",
	}), func(c *gin.Context) {
		token, _ := generateJWT()
		tokens = append(tokens, token)

		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	})

	r.GET("/jwt-auth", func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("Authorization")
		reqToken := strings.Split(bearerToken, " ")[1]
		claims := &Claims{}
		tkn, err := jwt.ParseWithClaims(reqToken, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": "unauthorized",
				})
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "bad request",
			})
			return
		}
		if !tkn.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": "jwt auth successful.",
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

func generateJWT() (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: "username",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
