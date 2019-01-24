package controllers

import (
	"net/http"

	"github.com/hbl-duytv/intern-csm/services"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{"transformListEditorUser": services.GetAllEditorUser()})
}
