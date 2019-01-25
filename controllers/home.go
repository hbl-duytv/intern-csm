package controllers

import (
	"net/http"

	"github.com/hbl-duytv/intern-csm/helper"
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
func GetToken(c *gin.Context) {
	token := c.Param("token")
	result := helper.GetToken(token)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": result})
}
