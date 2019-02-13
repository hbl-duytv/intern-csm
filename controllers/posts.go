package controllers

import (
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
	username := services.GetCurrentUser(c)
	if usernameString, ok := username.(string); ok {
		if user, err := services.GetUserByUsername(usernameString); err == nil {
			if posts, err := services.GetPostWithAdminPermission(); err == nil {
				if month, year, err := services.GetTimeCreateUSer(user.ID); err == nil {
					c.HTML(http.StatusOK, "master.html", gin.H{"month": month, "year": year, "user": user, "transformPost": posts, "index": 2, "title": "Post - Management"})
				}
			}
			return
		}
	}
	c.Redirect(http.StatusMovedPermanently, "/home")
}

func RenderPostManagementEditor(c *gin.Context) {
	var posts []models.TransformPost
	username := services.GetCurrentUser(c)
	if user, err := services.GetUserByUsername(username.(string)); err == nil {
		if posts, err = services.GetPostWithEditorPermission(user.ID); err == nil {
			if month, year, err := services.GetTimeCreateUSer(user.ID); err == nil {
				c.HTML(http.StatusOK, "master.html", gin.H{"month": month, "year": year, "user": user, "transformPost": posts, "index": 2, "title": "Post - Management"})
			}
		}
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/home")
}
func RenderCreatePost(c *gin.Context) {
	username := services.GetCurrentUser(c)
	if user, err := services.GetUserByUsername(username.(string)); err == nil {
		if month, year, err := services.GetTimeCreateUSer(user.ID); err == nil {
			c.HTML(http.StatusOK, "master.html", gin.H{"month": month, "year": year, "user": user, "index": 0, "title": "Create Post"})
		}
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/home")
}

func RenderUpdatePost(c *gin.Context) {
	if idPost, err := strconv.Atoi(c.Param("id")); err == nil {
		if post, err := services.GetPostById(idPost); err == nil {
			if month, year, err := services.GetTimeCreateUSer(post.Creator); err == nil {
				c.HTML(http.StatusOK, "master.html", gin.H{"month": month, "year": year, "post": post, "index": 3, "title": "Update Post"})
			}
			return
		}
		c.Redirect(http.StatusMovedPermanently, "/home")
	}
}

func RenderDetailPost(c *gin.Context) {
	idPost, err1 := strconv.Atoi(c.Param("id"))
	post, err2 := services.GetPostById(idPost)
	comments, err3 := services.GetCommentsById(idPost)
	usernames, err4 := services.GetNameInCommentByPostID(idPost)
	if err1 == nil && err2 == nil && err3 == nil && err4 == nil {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "post": post, "comment": comments, "username": usernames})
	}

}

func CreatePost(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("user")
	user, err1 := services.GetUserByUsername(username.(string))
	creator, err2 := strconv.Atoi(c.PostForm("creator"))
	title := c.PostForm("title")
	topic := c.PostForm("topic")
	description := c.PostForm("description")
	content := c.PostForm("content")
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": "Has error"})
		return
	}
	if strings.Trim(title, " ") == "" || strings.Trim(topic, " ") == "" || description == "" || strings.Trim(content, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": "Nhập đầy đủ thông tin"})
		return
	}
	post := models.Post{
		Creator:     creator,
		Title:       title,
		Topic:       topic,
		Description: description,
		Content:     content,
	}
	if services.CreatePost(&post) == nil {
		c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "user": user, "message": "create post successfully!"})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "user": user, "message": "create post fail!"})

}
func ActiveStatusPost(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("user")
	var user models.User
	var err1 error
	if usernameString, ok := username.(string); ok {
		user, err1 = services.GetUserByUsername(usernameString)
	}
	idPost, err2 := strconv.Atoi(c.PostForm("id"))
	messageComment := c.PostForm("comment")
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Has error!"})
		return
	}
	if services.ChangeStatusPostWithComment(idPost, user.ID, constant.ActiveNumber, messageComment) == nil {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Active status post successfully!"})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Active status post fail!"})
}
func DeActiveStatusPost(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("user")
	var user models.User
	var err1 error
	if usernameString, ok := username.(string); ok {
		user, err1 = services.GetUserByUsername(usernameString)
	}

	idPost, err2 := strconv.Atoi(c.PostForm("id"))
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Has error!"})
		return
	}
	messageComment := c.PostForm("comment")
	if services.ChangeStatusPostWithComment(idPost, user.ID, constant.DeactiveNumber, messageComment) == nil {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "DeActive status post successfully!"})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Deactive status post fail!"})

}
func DeletePost(c *gin.Context) {

	idPost, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Delete this post fail"})
		return
	}
	if services.DeletePost(idPost) == nil {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Delete this post successfully"})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Delete this post fail"})
}
func UpdateContentPost(c *gin.Context) {

	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "update post fail!"})
		return
	}
	title := c.PostForm("title")
	topic := c.PostForm("topic")
	description := c.PostForm("description")
	content := c.PostForm("content")

	if strings.Trim(title, " ") == "" || strings.Trim(topic, " ") == "" || description == "" || strings.Trim(content, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": "Nhập đầy đủ thông tin"})
		return
	}
	if services.UpdateContentPost(id, title, topic, description, content) == nil {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "update post successfully!"})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "update post fail!"})

}
