package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hbl-duytv/intern-csm/models"
	"github.com/hbl-duytv/intern-csm/services"
)

func GetAllEditorUser() []models.TransformUser {
	user := []models.User{}
	transformUSer := []models.TransformUser{}
	services.DB.Find(&user, "type=?", 0)
	if len(user) == 0 {
		return nil
	}
	for _, v := range user {
		transformUSer = append(transformUSer,
			models.TransformUser{v.ID, v.Username, v.Email, v.Type, v.Gender, v.BirthDay, v.PhoneNumber, v.Status})
	}
	return transformUSer
}
func RenderHome(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{"transformListEditorUser": GetAllEditorUser()})
}
