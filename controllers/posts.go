package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/hbl-duytv/intern-csm/models"
	"github.com/hbl-duytv/intern-csm/services"
)

func RenderPostManagementAdmin(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("user")
	if usernameString, ok := username.(string); ok {
		user, _ := services.GetUserByUsername(usernameString)
		posts, _ := services.GetPostWithAdminPermission()
		c.HTML(http.StatusOK, "master.html", gin.H{"user": user, "transformPost": posts, "index": 2, "title": "Post - Management"})
	} else {
		c.Redirect(301, "/home")
	}
	c.Redirect(301, "/home")

}

func RenderPostManagementEditor(c *gin.Context) {
	var posts []models.TransformPost
	var err error
	session := sessions.Default(c)
	username := session.Get("user")
	user, err1 := services.GetUserByUsername(username.(string))
	if err1 == nil {
		posts, err = services.GetPostWithEditorPermission(user.ID)
		if err == nil {
			c.HTML(http.StatusOK, "master.html", gin.H{"user": user, "transformPost": posts, "index": 2, "title": "Post - Management"})
		}
	}

}
func RenderCreatePost(c *gin.Context) {

	session := sessions.Default(c)
	username := session.Get("user")
	user, _ := services.GetUserByUsername(username.(string))

	c.HTML(http.StatusOK, "master.html", gin.H{"user": user, "index": 0, "title": "Create Post"})
}

func RenderUpdatePost(c *gin.Context) {
	idPost, _ := strconv.Atoi(c.Param("id"))

	post, _ := services.GetPostById(idPost)
	c.HTML(http.StatusOK, "master.html", gin.H{"post": post, "index": 3, "title": "Update Post"})

}
func RenderDetailPost(c *gin.Context) {
	idPost, _ := strconv.Atoi(c.Param("id"))
	var usernames []string

	post, _ := services.GetPostById(idPost)
	comments, _ := services.GetCommentsById(idPost)
	for _, v := range comments {
		var user models.User
		services.DB.Find(&user, "id=?", v.CommentatorID)
		usernames = append(usernames, user.Name)
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "post": post, "comment": comments, "username": usernames})

}
func CreatePost(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("user")
	user, _ := services.GetUserByUsername(username.(string))
	creator, _ := strconv.Atoi(c.PostForm("creator"))
	title := c.PostForm("title")
	topic := c.PostForm("topic")
	description := c.PostForm("description")
	content := c.PostForm("content")
	fmt.Println(title, topic, description, content)
	if strings.Trim(title, " ") == "" || strings.Trim(topic, " ") == "" || description == "" || strings.Trim(content, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": "Nhập đầy đủ thông tin"})
	} else {
		services.CreatePost(creator, title, topic, description, content)

		c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "user": user, "message": "create post successfully!"})
	}

}
func ActiveStatusPost(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("user")
	var user models.User
	if usernameString, ok := username.(string); ok {
		user, _ = services.GetUserByUsername(usernameString)
	} else {
		c.Redirect(301, "/login")
	}
	// var post models.Post
	idPost, _ := strconv.Atoi(c.PostForm("id"))
	messageComment := c.PostForm("comment")

	services.ChangeStatusPostWithComment(idPost, user.ID, 1, messageComment)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Active status post successfully!"})

}
func DeActiveStatusPost(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("user")
	var user models.User
	if usernameString, ok := username.(string); ok {
		user, _ = services.GetUserByUsername(usernameString)
	} else {
		c.Redirect(301, "/login")
	}

	idPost, _ := strconv.Atoi(c.PostForm("id"))
	messageComment := c.PostForm("comment")
	services.ChangeStatusPostWithComment(idPost, user.ID, 0, messageComment)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "DeActive status post successfully!"})

}
func DeletePost(c *gin.Context) {

	idPost, _ := strconv.Atoi(c.PostForm("id"))

	services.DeletePost(idPost)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Delete this post successfully"})
}
func UpdateContentPost(c *gin.Context) {

	id, _ := strconv.Atoi(c.PostForm("id"))

	title := c.PostForm("title")
	topic := c.PostForm("topic")
	description := c.PostForm("description")
	content := c.PostForm("content")

	if strings.Trim(title, " ") == "" || strings.Trim(topic, " ") == "" || description == "" || strings.Trim(content, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": "Nhập đầy đủ thông tin"})
	} else {
		services.UpdateContentPost(id, title, topic, description, content)

		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "update post successfully!"})
	}

}
