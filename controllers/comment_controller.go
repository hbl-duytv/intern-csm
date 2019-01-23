package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hbl-duytv/intern-csm/models"
	"github.com/hbl-duytv/intern-csm/services"
)

func GetAllCommentByPostID(c *gin.Context) {
	postID := c.Param("postID")
	comments := []models.Comment{}
	services.DB.Where("post_id = ?", postID).Find(&comments)
	if len(comments) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No commentation!"})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "comments": comments})
	}
}
