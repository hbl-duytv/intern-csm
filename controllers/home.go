package controllers

import (
	"net/http"

	"github.com/hbl-duytv/intern-csm/services"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("user")
	if usernameString, ok := username.(string); ok {
		user := services.GetUserByUsername(usernameString)
		c.HTML(http.StatusOK, "home.html", gin.H{"user": user})
	} else {
		c.Redirect(301, "/login")
	}
}
