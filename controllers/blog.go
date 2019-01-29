package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hbl-duytv/intern-csm/constant"
	"github.com/hbl-duytv/intern-csm/services"
)

func Blog(c *gin.Context) {
	page := c.Query("page")
	pageInt, _ := strconv.Atoi(page)
	posts, totalPost, errPost := services.GetPostActiveLimit(pageInt)
	if errPost != nil {
		c.HTML(http.StatusOK, "blog.html", gin.H{"posts": posts, "error": errPost})
		return
	}
	c.HTML(http.StatusOK, "blog.html", gin.H{"posts": posts, "totalPost": totalPost})
}
func BlogDetailPost(c *gin.Context) {
	postID := c.Param("postID")
	post, err := services.GetPostByID(postID)
	if err != nil {
		c.HTML(http.StatusOK, "blog-detail-post.html", gin.H{"post": nil})
		return
	}
	c.HTML(http.StatusOK, "blog-detail-post.html", gin.H{"post": post})
}
func GetTotalNumberAllPost(c *gin.Context) {
	totalAllPost, errPage := services.GetTotalNumberPost()
	var totalPage int
	totalPage = totalAllPost / constant.LimitPost
	remainder := totalAllPost % constant.LimitPost
	if remainder != 0 {
		totalPage++
	}
	if errPage != nil {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "error": errPage})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "totalPage": totalPage})
}
