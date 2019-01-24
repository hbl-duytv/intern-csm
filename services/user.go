package services

import (
	"github.com/hbl-duytv/intern-csm/helper"
	"github.com/hbl-duytv/intern-csm/models"
)

func GetAllEditorUser() []models.User {
	var users []models.User
	DB.Find(&users, "type=?", 0)
	if len(users) == 0 {
		return nil
	} else {
		return users
	}
}
func GetUserByUsername(username string) models.User {
	var user models.User
	DB.Find(&user, "username=?", username)
	return user
}
func GetUserByEmail(email string) models.User {
	var user models.User
	DB.Find(&user, "email=?", email)
	return user
}
func UpdateStatusUser(status int, user *models.User) {
	DB.Model(&user).Update("status", status)
}
func GetUserByID(id string) models.User {
	var user models.User
	DB.First(&user, id)
	return user
}
func CreateUser(username string, password string, email string, name string, gender string, birthday string, phoneNumber int, status int, typeUser int) {
	passwordMD5 := helper.GetMD5Hash(password)
	newUser := models.User{
		Username:    username,
		Password:    passwordMD5,
		Name:        name,
		Gender:      gender,
		PhoneNumber: phoneNumber,
		Email:       email,
		Type:        typeUser,
		Status:      status,
	}
	DB.Save(&newUser)
}
func DeleteUser(user *models.User) {
	DB.Delete(&user)
}
