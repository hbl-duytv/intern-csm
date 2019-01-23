package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/hbl-duytv/intern-csm/models"
	"github.com/hbl-duytv/intern-csm/services"
)

func GetPostWithAdminPermission(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("user")
	// txtStatus := "Bài viết chưa được duyệt"
	var transformPost []models.TransformPost
	var posts []models.Posts
	// status := 0
	services.DB.Find(&posts)
	if len(posts) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No posts found!"})
		return
	}
	for _, v := range posts {
		var user models.User
		services.DB.Select("name").Where("id=?", v.Creator).Find(&user)
		fmt.Println(services.DB.Debug().Select("name").Where("id=?", v.Creator).Find(&user).Error)
		transformPost = append(transformPost, models.TransformPost{
			v.ID,
			user.Name,
			v.Title,
			v.Topic,
			v.Description,
			v.Content,
			v.Status,
			v.CreatedAt,
			v.UpdatedAt,
		})
	}

	if usernameString, ok := username.(string); ok {
		user := GetCurrentUser(usernameString)
		c.HTML(http.StatusOK, "post-management.html", gin.H{"user": user, "transformPost": transformPost})
	} else {
		c.Redirect(301, "/home")
	}
	c.Redirect(301, "/home")
	// c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": transformPost})
}

func GetPostWithEditorPermission(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("user")
	user := GetCurrentUser(username.(string))
	// idUser := c.Param("id")
	// txtStatus := ""
	var tranformPost []models.TransformPost
	var posts []models.Posts
	services.DB.Find(&posts, "creator=?", user.ID)
	if len(posts) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No posts found!"})
		return
	}
	for _, v := range posts {
		// if v.Status == 0 {
		// 	txtStatus = "Bài viết chưa được duyệt"
		// } else {
		// 	txtStatus = "Bài viết đã được duyệt"
		// }
		var user models.User
		services.DB.Select("name").Where("id=?", v.Creator).Find(&user)

		tranformPost = append(tranformPost, models.TransformPost{
			v.ID,
			user.Name,
			v.Title,
			v.Topic,
			v.Description,
			v.Content,
			v.Status,
			v.CreatedAt,
			v.UpdatedAt,
		})
	}
	c.HTML(http.StatusOK, "post-management.html", gin.H{"user": user, "transformPost": tranformPost})

}
func RenderCreatePost(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("user")
	user := GetCurrentUser(username.(string))

	c.HTML(http.StatusOK, "create-post.html", gin.H{"user": user})

}
func RenderUpdatePost(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("user")
	user := GetCurrentUser(username.(string))
	idPost := c.PostForm("id")
	var posts models.Posts
	services.DB.Where("id=?", idPost).Find(&posts)
	c.HTML(http.StatusOK, "update-post.html", gin.H{"user": user, "posts": posts})

}
func CreatePost(c *gin.Context) {

	creator, _ := strconv.Atoi(c.PostForm("creator"))

	title := c.PostForm("title")
	topic := c.PostForm("topic")
	description := c.PostForm("description")
	content := c.PostForm("content")
	fmt.Println(title, topic, description, content)
	if strings.Trim(title, " ") == "" || strings.Trim(topic, " ") == "" || description == "" || strings.Trim(content, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": "Nhập đầy đủ thông tin"})
	} else {
		newPost := models.Posts{

			Creator:     creator,
			Title:       title,
			Topic:       topic,
			Description: description,
			Content:     content,
			Status:      0,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		services.DB.Save(&newPost)

		c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "create post successfully!"})
	}

}
func ActiveStatusPost(c *gin.Context) {
	var post models.Posts
	idPost := c.PostForm("id")
	services.DB.First(&post, idPost)
	if post.ID != 0 {
		services.DB.Model(&post).Update("status", 1)

		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Active status post successfully!"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Updated status post fail!"})
	}

}
func DeActiveStatusPost(c *gin.Context) {

	var post models.Posts
	idPost := c.PostForm("id")
	services.DB.First(&post, idPost)
	if post.ID != 0 {
		services.DB.Model(&post).Update("status", 0)

		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "DeActive status post successfully!"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Updated status post fail!"})
	}
}
func DeletePost(c *gin.Context) {

	idPost := c.PostForm("id")
	services.DB.Where("id=?", idPost).Delete(models.Posts{})
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Delete this post successfully"})
}
func UpdateContentPost(c *gin.Context) {

	id := c.PostForm("id")
	var post models.Posts
	services.DB.Where("id=?", id).Find(&post)
	title := c.PostForm("title")
	topic := c.PostForm("topic")
	description := c.PostForm("description")
	content := c.PostForm("content")

	if strings.Trim(id, " ") == "" || strings.Trim(title, " ") == "" || strings.Trim(topic, " ") == "" || description == "" || strings.Trim(content, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": "Nhập đầy đủ thông tin"})
	} else {

		post.Title = title
		post.Topic = topic
		post.Description = description
		post.Content = content
		services.DB.Save(&post)

		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "update post successfully!", "post": post})
	}

}
