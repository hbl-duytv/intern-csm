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
		if user, err := services.GetUserByUsername(usernameString); err == nil {
			month, year, _ := services.GetTimeCreateUSer(user.ID)
			c.HTML(http.StatusOK, "master.html", gin.H{"user": user, "index": -1, "title": "Home", "month": month, "year": year})
			return
		}

	}

}
