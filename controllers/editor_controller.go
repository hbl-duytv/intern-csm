package controllers

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/hbl-duytv/intern-csm/models"
	"github.com/hbl-duytv/intern-csm/services"
)

func GetAllEditorUser() []models.TransformUser {
	users := []models.User{}
	transformUSers := []models.TransformUser{}
	services.DB.Find(&users, "type=?", 0)
	if len(users) == 0 {
		return nil
	}
	for _, v := range users {
		transformUSers = append(transformUSers,
			models.TransformUser{v.ID, v.Username, v.Email, v.Name, v.Type, v.Gender, v.BirthDay, v.PhoneNumber, v.Status})
	}
	return transformUSers
}
func ActiveEditorUser(c *gin.Context) {
	var user models.User
	idUser := c.PostForm("id")
	services.DB.First(&user, idUser)
	if user.ID != 0 {
		services.DB.Model(&user).Update("status", 1)
		services.DB.Save(&user)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Updated status user successfully!"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Updated status user fail!"})
	}
}
func DeactiveEditorUser(c *gin.Context) {
	var user models.User
	idUser := c.PostForm("id")
	services.DB.First(&user, idUser)
	if user.ID != 0 {
		services.DB.Model(&user).Update("status", 0)
		services.DB.Save(&user)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Updated status user successfully!"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Updated status user fail!"})
	}
}
func DeleteUser(c *gin.Context) {
	var user models.User
	idUser := c.PostForm("id")
	services.DB.First(&user, idUser)
	if user.ID != 0 {
		services.DB.Delete(&user)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Delete user successfully!"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Delete user fail!"})
	}

}
func RenderEditorManagement(c *gin.Context) {
	editors := GetAllEditorUser()
	session := sessions.Default(c)
	username := session.Get("user")
	if usernameString, ok := username.(string); ok {
		currentUser := GetCurrentUser(usernameString)
		c.HTML(http.StatusOK, "editor-management.html", gin.H{"editors": editors, "currentUser": currentUser})
	} else {
		c.Redirect(301, "/home")
	}
	c.Redirect(301, "/home")
}
