package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/basic-auth", gin.BasicAuth(gin.Accounts{
		"admin": "secret",
	}), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "basic auth success",
		})
	})
	r.Run() // default port :8080
}
