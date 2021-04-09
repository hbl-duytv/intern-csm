package controllers

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/hbl-duytv/intern-csm/models"
	"github.com/hbl-duytv/intern-csm/services"
)

func GetCurrentUser(username string) models.User {
	// user := models.User{}
	var user models.User
	services.DB.Find(&user, "username=?", username)
	// transformUSer := models.TransformUser{}
	// if user.ID != 0 {
	// 	transformUSer = models.TransformUser{user.ID, user.Username, user.Email, user.Name, user.Type, user.Gender, user.BirthDay, user.PhoneNumber, user.Status}
	// 	return transformUSer
	// }
	return user
}
func Home(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("user")
	if usernameString, ok := username.(string); ok {
		user := GetCurrentUser(usernameString)
		c.HTML(http.StatusOK, "master.html", gin.H{"user": user, "index": -1, "title": "Home"})
	} else {
		c.Redirect(301, "/login")
	}
	c.Redirect(301, "/login")
}

// func Home1(c *gin.Context) {
// 	fmt.Println("render home 1")
// 	c.HTML(http.StatusOK, "create-post-1.html", gin.H{"status": http.StatusOK, "title": "Test"})

// }
