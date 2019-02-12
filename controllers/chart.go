package controllers

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/hbl-duytv/intern-csm/constant"
	"github.com/hbl-duytv/intern-csm/services"
)

func RenderChart(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("user")
	if usernameString, ok := username.(string); ok {
		if user, err := services.GetUserByUsername(usernameString); err == nil {
			if arrCount, err := services.GetCountPostInYear(); err == nil {
				c.JSON(http.StatusOK, gin.H{
					"status":   http.StatusOK,
					"user":     user,
					"index":    4,
					"title":    "Chart",
					"month":    user.CreatedAt.Month(),
					"year":     user.CreatedAt.Year(),
					"arrCount": arrCount,
				})
			}

		}
	} else {
		c.Redirect(constant.DirectStatus, "/login")
	}
}
