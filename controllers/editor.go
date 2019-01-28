package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/hbl-duytv/intern-csm/constant"
	"github.com/hbl-duytv/intern-csm/helper"
	"github.com/hbl-duytv/intern-csm/models"
	"github.com/hbl-duytv/intern-csm/services"
)

func ActiveEditorUser(c *gin.Context) {
	// var user models.User
	idUser, _ := strconv.Atoi(c.PostForm("id"))
	if services.UpdateStatusUser(idUser, constant.ACTIVE_NUMBER) == nil {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Updated status user successfully!"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Updated status user fail!"})
	}

	// services.DB.First(&user, idUser)
	// if user.ID != 0 {
	// 	services.DB.Model(&user).Update("status", 1)
	// 	services.DB.Save(&user)
	// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Updated status user successfully!"})
	// } else {
	// 	c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Updated status user fail!"})
	// }
}
func DeactiveEditorUser(c *gin.Context) {
	// var user models.User
	idUser, _ := strconv.Atoi(c.PostForm("id"))
	if services.UpdateStatusUser(idUser, constant.DEACTIVE_NUMBER) == nil {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Updated status user successfully!"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Updated status user fail!"})
	}
	// services.DB.First(&user, idUser)
	// if user.ID != 0 {
	// 	services.DB.Model(&user).Update("status", 0)
	// 	services.DB.Save(&user)
	// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Updated status user successfully!"})
	// } else {
	// 	c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Updated status user fail!"})
	// }
}
func DeleteUser(c *gin.Context) {
	// var user models.User
	idUser, _ := strconv.Atoi(c.PostForm("id"))
	if services.DeletePost(idUser) == nil {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Delete user successfully!"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Delete user fail!"})
	}
	// services.DB.First(&user, idUser)
	// if user.ID != 0 {
	// 	services.DB.Delete(&user)
	// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Delete user successfully!"})
	// } else {
	// 	c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Delete user fail!"})
	// }

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
		newUser := models.User{
			Username:    username,
			Password:    helper.GetMD5Hash(password),
			Email:       email,
			Name:        name,
			Gender:      gender,
			BirthDay:    birthday,
			PhoneNumber: phoneNumber,
			Type:        0,
			Status:      status,
		}
		services.DB.Create(&newUser)
		c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "User create success!"})
	}
}

func RenderEditorManagement(c *gin.Context) {
	fmt.Println("home0")
	editors, _ := services.GetAllEditorUser()
	session := sessions.Default(c)
	username := session.Get("user")
	if usernameString, ok := username.(string); ok {
		currentUser, _ := services.GetUserByUsername(usernameString)
		c.HTML(http.StatusOK, "master.html", gin.H{"editors": editors, "currentUser": currentUser, "index": 1, "title": "Editor management"})
	} else {
		c.Redirect(301, "/home")
	}
	c.Redirect(301, "/home")
}
