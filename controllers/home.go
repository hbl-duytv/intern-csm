package controllers

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/hbl-duytv/intern-csm/services"
)

func Home(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("user")
	if usernameString, ok := username.(string); ok {
		user, _ := services.GetUserByUsername(usernameString)
		c.HTML(http.StatusOK, "master.html", gin.H{"user": user, "index": -1, "title": "Home"})
	} else {
		c.Redirect(301, "/login")
	}
}
