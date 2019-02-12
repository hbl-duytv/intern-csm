package middleware

import (
	"net/http"

	"github.com/hbl-duytv/intern-csm/constant"
	"github.com/hbl-duytv/intern-csm/services"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		if user == nil {
			// redirect to login page when user not login
			c.Redirect(301, "login")
		} else {
			// Continue down the chain to handler etc
			c.Next()
		}
	}
}

func AuthAdminRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		username := session.Get("user")
		if username == nil {
			c.Redirect(301, "login")
		} else if usernameString, ok := username.(string); ok {
			user, _ := services.GetUserByUsername(usernameString)
			if user.Type == constant.ActiveNumber {
				c.Next()
			} else {
				c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "User not allowed"})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		}
	}
}
