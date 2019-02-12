package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/hbl-duytv/intern-csm/constant"
	"github.com/hbl-duytv/intern-csm/helper"
	"github.com/hbl-duytv/intern-csm/services"
)

func ActiveEditorUser(c *gin.Context) {
	idUser, _ := strconv.Atoi(c.PostForm("id"))
	if services.UpdateStatusUser(idUser, constant.ActiveNumber) == nil {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Updated status user successfully!"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Updated status user fail!"})
	}
}

func DeactiveEditorUser(c *gin.Context) {
	idUser, _ := strconv.Atoi(c.PostForm("id"))
	if services.UpdateStatusUser(idUser, constant.DeactiveNumber) == nil {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Updated status user successfully!"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Updated status user fail!"})
	}
}

func DeleteUser(c *gin.Context) {
	idUser, _ := strconv.Atoi(c.PostForm("id"))
	if services.DeleteUser(idUser) == nil {
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
	token := helper.GetToken(username)
	if phoneNumber, error := strconv.Atoi(c.PostForm("phone_number")); error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Phone number error!"})
	} else if status, error := strconv.Atoi(c.PostForm("status")); error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Status error!"})
	} else {
		if services.CreateAccount(username, password, email, name, gender, birthday, phoneNumber, status, constant.DeactiveNumber, token) == nil {
			c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "User create success!"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "User create success!"})
		}
	}
}

func RenderEditorManagement(c *gin.Context) {
	editors, _ := services.GetAllEditorUser()
	session := sessions.Default(c)
	session.Set("user", session_name)
	username := session.Get("user")

	if usernameString, ok := username.(string); ok {
		if user, err := services.GetUserByUsername(usernameString); err == nil {
			c.HTML(http.StatusOK, "master.html", gin.H{"month": user.CreatedAt.Month(), "year": user.CreatedAt.Year(), "editors": editors, "user": user, "index": 1, "title": "Editor management"})
		}
	} else {
		c.Redirect(constant.DirectStatus, "/home")
	}
}
