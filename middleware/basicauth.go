package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BasicAuth(c *gin.Context) {
	uname, pwd, ok := c.Request.BasicAuth()
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Username or Password must be filled",
		})
		return
	}

	if uname == "admin" && pwd == "password" {
		return
	} else if uname == "editor" && pwd == "secret" {
		return
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "User is unauthorized",
		})
		return
	}
}