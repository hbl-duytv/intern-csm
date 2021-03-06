package controllers

import (
	"net/http"
	"strconv"

	"github.com/hbl-duytv/intern-csm/models"

	"github.com/gin-gonic/gin"
	"github.com/hbl-duytv/intern-csm/constant"
	"github.com/hbl-duytv/intern-csm/helper"
	"github.com/hbl-duytv/intern-csm/services"
)

func ActiveEditorUser(c *gin.Context) {

	idUser, _ := strconv.Atoi(c.PostForm("id"))
	if services.UpdateStatusUser(idUser, constant.ActiveNumber) == nil {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Updated status user successfully!"})
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Updated status user fail!"})
}
func DeactiveEditorUser(c *gin.Context) {
	idUser, _ := strconv.Atoi(c.PostForm("id"))
	if services.UpdateStatusUser(idUser, constant.TypeEditor) == nil {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Updated status user successfully!"})
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Updated status user fail!"})
}
func DeleteUser(c *gin.Context) {
	idUser, _ := strconv.Atoi(c.PostForm("id"))
	if services.DeleteUser(idUser) == nil {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Delete user successfully!"})
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Delete user fail!"})
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
		newUser := models.User{
			Username:    username,
			Password:    helper.GetMD5Hash(constant.PasswordUserDefault),
			Email:       email,
			Name:        name,
			Gender:      gender,
			Birthday:    birthday,
			PhoneNumber: phoneNumber,
			Status:      status,
			Type:        constant.TypeEditor,
			Token:       token,
			Confirm:     constant.DeactiveNumber,
		}
		if services.CreateAccount(&newUser) == nil {
			c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "User create success!"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "User create success!"})

	}
}

func RenderEditorManagement(c *gin.Context) {

	editors, err := services.GetAllEditorUser()
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/home")
		return
	}
	username := services.GetCurrentUser(c)
	if usernameString, ok := username.(string); ok {
		if user, err := services.GetUserByUsername(usernameString); err == nil {

			c.HTML(http.StatusOK, "master.html", gin.H{"month": user.CreatedAt.Month(), "year": user.CreatedAt.Year(), "editors": editors, "user": user, "index": 1, "title": "Editor management"})
			return
		}

	}
	c.Redirect(http.StatusMovedPermanently, "/home")

}
