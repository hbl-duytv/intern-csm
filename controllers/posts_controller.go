package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/hbl-duytv/intern-csm/models"
	"github.com/hbl-duytv/intern-csm/services"
)

func GetPostWithAdminPermission(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("user")
	txtStatus := "Bài viết chưa được duyệt"
	var transformPost []models.TransformPost
	var posts []models.Posts
	status := 0
	services.DB.Find(&posts, "status=?", status)
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
			txtStatus,
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
	txtStatus := ""
	var tranformPost []models.TransformPost
	var posts []models.Posts
	services.DB.Find(&posts, "creator=?", user.ID)
	if len(posts) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No posts found!"})
		return
	}
	for _, v := range posts {
		if v.Status == 0 {
			txtStatus = "Bài viết chưa được duyệt"
		} else {
			txtStatus = "Bài viết đã được duyệt"
		}
		var user models.User
		services.DB.Select("name").Where("id=?", v.Creator).Find(&user)

		tranformPost = append(tranformPost, models.TransformPost{
			v.ID,
			user.Name,
			v.Title,
			v.Topic,
			v.Description,
			v.Content,
			txtStatus,
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
func CreatePost(c *gin.Context) {

	creator, _ := strconv.Atoi(c.PostForm("creator"))
	// status, _ := strconv.Atoi(c.PostForm("status"))
	title := c.PostForm("title")
	topic := c.PostForm("topic")
	description := c.PostForm("description")
	content := c.PostForm("content")
	fmt.Println(title, topic, description, content)
	if title == "" || topic == "" || description == "" || content == "" {
		c.JSON(http.StatusNoContent, gin.H{"error": "Parameters can't be empty"})
		return
	}
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
	fmt.Println(services.DB.Debug().Save(&newPost).Error)
	c.JSON(http.StatusOK, gin.H{"message": "create post successfully"})
	return
}
func UpdateStatusPost(c *gin.Context) {
	var post models.Posts
	idPost := c.PostForm("id")
	services.DB.First(&post, idPost)
	if post.ID != 0 {
		services.DB.Model(&post).Update("status", 1)
		services.DB.Save(&post)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Updated status post successfully!"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Updated status post fail!"})
	}

}
