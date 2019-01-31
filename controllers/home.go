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
		if user, err := services.GetUserByUsername(usernameString); err == nil {
			c.HTML(http.StatusOK, "home.html", gin.H{"user": user})
			return
		}
	}
	c.Redirect(301, "/login")
}
func NumberPostInYear(c *gin.Context) {
	var numberPosts [12]int
	var err error
	for index := 0; index < 12; index++ {
		numberPosts[index], err = services.GetNumberPostByMonth(index + 1)
	}
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "posts": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "numberPosts": numberPosts})
}
