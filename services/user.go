package services

import (
	"fmt"

	"github.com/hbl-duytv/intern-csm/helper"
	"github.com/hbl-duytv/intern-csm/models"
)

func GetAllEditorUser() []models.User {
	var users []models.User
	DB.Find(&users, "type=? AND confirm=?", 0, 1)
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
func GetUserByToken(token string) models.User {
	var user models.User
	DB.Find(&user, "token=?", token)
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
func CreateUser(username string, password string, email string, name string, gender string, birthday string, phoneNumber int, status int, typeUser int, token string, confirm int) {
	passwordMD5 := helper.GetMD5Hash(password)
	fmt.Println(birthday)
	newUser := models.User{
		Username:    username,
		Password:    passwordMD5,
		Name:        name,
		Gender:      gender,
		Birthday:    birthday,
		PhoneNumber: phoneNumber,
		Email:       email,
		Type:        typeUser,
		Status:      status,
		Token:       token,
		Confirm:     confirm,
	}
	DB.Save(&newUser)
}
func DeleteUser(user *models.User) {
	DB.Delete(&user)
}
func ConfirmRegisterUser(user *models.User) {
	DB.Model(&user).Update("confirm", 1)
}
