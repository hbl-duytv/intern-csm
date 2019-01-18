package controllers

import (
	"github.com/hbl-duytv/intern-csm/database"
	"github.com/hbl-duytv/intern-csm/models"
)

func GetAllEditorUser() []models.TransformUser {
	user := []models.User{}
	transformUSer := []models.TransformUser{}
	database.DB.Find(&user, "type=?", 0)
	if len(user) == 0 {
		return nil
	}
	for _, v := range user {
		transformUSer = append(transformUSer,
			models.TransformUser{v.ID, v.Username, v.Email, v.Type, v.Gender, v.BirthDay, v.PhoneNumber, v.Status})
	}
	return transformUSer
}
