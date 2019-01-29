package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/hbl-duytv/intern-csm/constant"
	"github.com/hbl-duytv/intern-csm/models"
	"github.com/hbl-duytv/intern-csm/services"
)

func RenderPostManagementAdmin(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("user")
	if usernameString, ok := username.(string); ok {
		if user, err := services.GetUserByUsername(usernameString); err == nil {
			if posts, err := services.GetPostWithAdminPermission(); err == nil {
				month, year, _ := services.GetTimeCreateUSer(user.ID)
				c.HTML(http.StatusOK, "master.html", gin.H{"month": month, "year": year, "user": user, "transformPost": posts, "index": 2, "title": "Post - Management"})
			}
		}

	} else {
		c.Redirect(constant.DIRECT_STATUS, "/home")
	}

}

func RenderPostManagementEditor(c *gin.Context) {
	var posts []models.TransformPost

	session := sessions.Default(c)
	username := session.Get("user")
	if user, err := services.GetUserByUsername(username.(string)); err == nil {
		if posts, err = services.GetPostWithEditorPermission(user.ID); err == nil {
			month, year, _ := services.GetTimeCreateUSer(user.ID)
			c.HTML(http.StatusOK, "master.html", gin.H{"month": month, "year": year, "user": user, "transformPost": posts, "index": 2, "title": "Post - Management"})
		}
	} else {
		c.Redirect(constant.DIRECT_STATUS, "/home")
	}
}
func RenderCreatePost(c *gin.Context) {

	session := sessions.Default(c)
	username := session.Get("user")
	if user, err := services.GetUserByUsername(username.(string)); err == nil {
		month, year, _ := services.GetTimeCreateUSer(user.ID)
		c.HTML(http.StatusOK, "master.html", gin.H{"month": month, "year": year, "user": user, "index": 0, "title": "Create Post"})
	} else {
		c.Redirect(constant.DIRECT_STATUS, "/home")
	}
}

func RenderUpdatePost(c *gin.Context) {
	idPost, _ := strconv.Atoi(c.Param("id"))

	if post, err := services.GetPostById(idPost); err == nil {
		month, year, _ := services.GetTimeCreateUSer(post.Creator)
		c.HTML(http.StatusOK, "master.html", gin.H{"month": month, "year": year, "post": post, "index": 3, "title": "Update Post"})
	} else {
		c.Redirect(constant.DIRECT_STATUS, "/home")
	}

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
		if services.CreatePost(creator, title, topic, description, content) == nil {
			c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "user": user, "message": "create post successfully!"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "user": user, "message": "create post fail!"})
		}

	}

}
func ActiveStatusPost(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("user")
	var user models.User
	if usernameString, ok := username.(string); ok {
		user, _ = services.GetUserByUsername(usernameString)
	} else {
		c.Redirect(constant.DIRECT_STATUS, "/login")
	}
	// var post models.Post
	idPost, _ := strconv.Atoi(c.PostForm("id"))
	messageComment := c.PostForm("comment")

	if services.ChangeStatusPostWithComment(idPost, user.ID, constant.ACTIVE_NUMBER, messageComment) == nil {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Active status post successfully!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Active status post fail!"})
	}

}
func DeActiveStatusPost(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("user")
	var user models.User
	if usernameString, ok := username.(string); ok {
		user, _ = services.GetUserByUsername(usernameString)
	} else {
		c.Redirect(constant.DIRECT_STATUS, "/login")
	}

	idPost, _ := strconv.Atoi(c.PostForm("id"))
	messageComment := c.PostForm("comment")
	if services.ChangeStatusPostWithComment(idPost, user.ID, constant.DEACTIVE_NUMBER, messageComment) == nil {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "DeActive status post successfully!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Deactive status post fail!"})
	}

}
func DeletePost(c *gin.Context) {

	idPost, _ := strconv.Atoi(c.PostForm("id"))

	if services.DeletePost(idPost) == nil {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Delete this post successfully"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Delete this post fail"})
	}

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
		if services.UpdateContentPost(id, title, topic, description, content) == nil {
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "update post successfully!"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "update post fail!"})
		}

	}

}
