package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/hbl-duytv/intern-csm/helper"
	"github.com/hbl-duytv/intern-csm/services"
)

func ActiveEditorUser(c *gin.Context) {
	idUser := c.PostForm("id")
	user, err := services.GetUserByID(idUser)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Updated status user fail!"})
		return
	}
	services.UpdateStatusUser(1, &user)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Updated status user successfully!"})
}
func DeactiveEditorUser(c *gin.Context) {
	idUser := c.PostForm("id")
	user, err := services.GetUserByID(idUser)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Updated status user fail!"})
		return
	}
	services.UpdateStatusUser(0, &user)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Updated status user successfully!"})
}
func DeleteUser(c *gin.Context) {
	idUser := c.PostForm("id")
	user, err := services.GetUserByID(idUser)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Delete user fail!"})
		return
	}
	services.DeleteUser(&user)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Delete user successfully!"})
}
func CreateUser(c *gin.Context) {
	username := c.PostForm("username")
	email := c.PostForm("email")
	name := c.PostForm("name")
	gender := c.PostForm("gender")
	birthday := c.PostForm("birthday")
	token := helper.GetToken(username)
	if phoneNumber, error := strconv.Atoi(c.PostForm("phone_number")); error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Phone number error!"})
	} else if status, error := strconv.Atoi(c.PostForm("status")); error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Status error!"})
	} else {
		services.CreateUser(username, "123456", email, name, gender, birthday, phoneNumber, status, 0, token, 1)
		c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "User create success!"})
	}
}
func EditorManagement(c *gin.Context) {
	editors, error := services.GetAllEditorUser()
	if error != nil {
		fmt.Print(error)
		return
	}
	session := sessions.Default(c)
	username := session.Get("user")
	if usernameString, ok := username.(string); ok {
		if currentUser, err := services.GetUserByUsername(usernameString); err == nil {
			c.HTML(http.StatusOK, "editor-management.html", gin.H{"editors": editors, "currentUser": currentUser})
			return
		}
	}
	c.Redirect(301, "/home")
}
