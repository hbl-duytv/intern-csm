package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/hbl-duytv/intern-csm/models"
	"github.com/hbl-duytv/intern-csm/services"
)

func ActiveEditorUser(c *gin.Context) {
	var user models.User
	idUser := c.PostForm("id")
	services.GetUserByID(idUser)
	if user.ID != 0 {
		services.UpdateStatusUser(1, &user)
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
		services.UpdateStatusUser(0, &user)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Updated status user successfully!"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Updated status user fail!"})
	}
}
func DeleteUser(c *gin.Context) {
	var user models.User
	idUser := c.PostForm("id")
	services.GetUserByID(idUser)
	if user.ID != 0 {
		services.DeleteUser(&user)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Delete user successfully!"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Delete user fail!"})
	}

}
func CreateUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")
	name := c.PostForm("name")
	gender := c.PostForm("gender")
	birthday := c.PostForm("birthday")
	if phoneNumber, error := strconv.Atoi(c.PostForm("phone_number")); error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Phone number error!"})
	} else if status, error := strconv.Atoi(c.PostForm("status")); error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Status error!"})
	} else {
		services.CreateUser(username, password, email, name, gender, birthday, phoneNumber, status, 0)
		c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "User create success!"})
	}
}
func EditorManagement(c *gin.Context) {
	editors := services.GetAllEditorUser()
	session := sessions.Default(c)
	username := session.Get("user")
	if usernameString, ok := username.(string); ok {
		currentUser := services.GetUserByUsername(usernameString)
		c.HTML(http.StatusOK, "editor-management.html", gin.H{"editors": editors, "currentUser": currentUser})
	} else {
		c.Redirect(301, "/home")
	}
	c.Redirect(301, "/home")
}
