package controllers

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/hbl-duytv/intern-csm/constant"
	"github.com/hbl-duytv/intern-csm/services"
)

func Home(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("user")
	if usernameString, ok := username.(string); ok {
		if user, err := services.GetUserByUsername(usernameString); err == nil {

			c.HTML(http.StatusOK, "master.html", gin.H{"user": user, "index": -1, "title": "Home", "month": user.CreatedAt.Month(), "year": user.CreatedAt.Year()})
		}
	} else {
		c.Redirect(constant.DirectStatus, "/login")
	}
}
