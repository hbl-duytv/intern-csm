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
		user := GetCurrentUser(usernameString)
		c.HTML(http.StatusOK, "post-management.html", gin.H{"user": user, "transformPost": services.GetPostWithAdminPermission()})
	} else {
		c.Redirect(301, "/home")
	}
	c.Redirect(301, "/home")

}

func RenderPostManagementEditor(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("user")
	user := GetCurrentUser(username.(string))
	c.HTML(http.StatusOK, "post-management.html", gin.H{"user": user, "transformPost": services.GetPostWithEditorPermission(user.ID)})

}
func RenderCreatePost(c *gin.Context) {

	session := sessions.Default(c)
	username := session.Get("user")
	user := GetCurrentUser(username.(string))

	c.HTML(http.StatusOK, "create-post.html", gin.H{"user": user})
}

func RenderUpdatePost(c *gin.Context) {
	idPost, _ := strconv.Atoi(c.Param("id"))
	// var post models.Post
	// services.DB.Find(&post, "id=?", idPost)
	c.HTML(http.StatusOK, "update-post.html", gin.H{"post": services.GetPostById(idPost)})

}
func RenderDetailPost(c *gin.Context) {
	idPost, _ := strconv.Atoi(c.Param("id"))
	var usernames []string

	// var post models.Post
	// var comment []models.Comment
	// services.DB.Find(&post, "id=?", idPost)
	// services.DB.Find(&comment, "post_id=?", idPost)
	post := services.GetPostById(idPost)
	comments := services.GetCommentsById(idPost)
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
	user := GetCurrentUser(username.(string))
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
		// newPost := models.Post{

		// 	Creator:     creator,
		// 	Title:       title,
		// 	Topic:       topic,
		// 	Description: description,
		// 	Content:     content,
		// 	Status:      0,
		// 	CreatedAt:   time.Now(),
		// 	UpdatedAt:   time.Now(),
		// }
		// services.DB.Save(&newPost)

		c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "user": user, "message": "create post successfully!"})
	}

}
func ActiveStatusPost(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("user")
	var user models.User
	if usernameString, ok := username.(string); ok {
		user = GetCurrentUser(usernameString)
	} else {
		c.Redirect(301, "/login")
	}
	// var post models.Post
	idPost, _ := strconv.Atoi(c.PostForm("id"))
	messageComment := c.PostForm("comment")

	services.ChangeStatusPostWithComment(idPost, user.ID, 1, messageComment)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Active status post successfully!"})
	// fmt.Println("message comment : ", messageComment)
	// idPostInt, _ := strconv.Atoi(idPost)
	// services.DB.First(&post, idPost)
	// if post.ID != 0 {
	// 	services.DB.Model(&post).Update("status", 1)
	// 	newComment := models.Comment{
	// 		PostID:        idPostInt,
	// 		CommentatorID: user.ID,
	// 		Message:       messageComment,
	// 	}
	// 	services.DB.Save(&newComment)
	// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Active status post successfully!"})
	// } else {
	// 	c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Updated status post fail!"})
	// }

}
func DeActiveStatusPost(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("user")
	var user models.User
	if usernameString, ok := username.(string); ok {
		user = GetCurrentUser(usernameString)
	} else {
		c.Redirect(301, "/login")
	}
	// var post models.Post
	idPost, _ := strconv.Atoi(c.PostForm("id"))
	messageComment := c.PostForm("comment")
	services.ChangeStatusPostWithComment(idPost, user.ID, 0, messageComment)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "DeActive status post successfully!"})
	// services.DB.First(&post, idPost)
	// idPostInt, _ := strconv.Atoi(idPost)
	// if post.ID != 0 {
	// 	services.DB.Model(&post).Update("status", 0)
	// 	newComment := models.Comment{
	// 		PostID:        idPostInt,
	// 		CommentatorID: user.ID,
	// 		Message:       messageComment,
	// 	}
	// 	services.DB.Save(&newComment)
	// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "DeActive status post successfully!"})
	// } else {
	// 	c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Updated status post fail!"})
	// }
}
func DeletePost(c *gin.Context) {

	idPost, _ := strconv.Atoi(c.PostForm("id"))
	// services.DB.Where("id=?", idPost).Delete(models.Post{})
	services.DeletePost(idPost)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Delete this post successfully"})
}
func UpdateContentPost(c *gin.Context) {

	id, _ := strconv.Atoi(c.PostForm("id"))
	// var post models.Post
	// services.DB.Where("id=?", id).Find(&post)
	title := c.PostForm("title")
	topic := c.PostForm("topic")
	description := c.PostForm("description")
	content := c.PostForm("content")

	if strings.Trim(title, " ") == "" || strings.Trim(topic, " ") == "" || description == "" || strings.Trim(content, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": "Nhập đầy đủ thông tin"})
	} else {
		services.UpdateContentPost(id, title, topic, description, content)
		// post.Title = title
		// post.Topic = topic
		// post.Description = description
		// post.Content = content
		// services.DB.Save(&post)

		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "update post successfully!"})
	}

}
