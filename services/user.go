package services

import (
	"fmt"

	"github.com/hbl-duytv/intern-csm/constant"
	"github.com/hbl-duytv/intern-csm/helper"
	"github.com/hbl-duytv/intern-csm/models"
)

func GetAllEditorUser() ([]models.User, error) {
	var users []models.User
	if err := DB.Find(&users, "type=?", constant.Editor).Error; err != nil {
		return users, err
	}
	return users, nil
}
func GetUserByUsername(username string) (models.User, error) {
	var user models.User
	if err := DB.Find(&user, "username=?", username).Error; err != nil {
		return user, err
	}
	return user, nil
}
func GetUserByUsernamePassword(username string, password string) (models.User, error) {
	var user models.User
	passwordMD5 := helper.GetMD5Hash(password)
	if err := DB.Where("username = ? AND password = ?", username, passwordMD5).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	if err := DB.Find(&user, "email=?", email).Error; err != nil {
		return user, err
	}
	return user, nil
}
func GetUserByToken(token string) (models.User, error) {
	var user models.User
	if err := DB.Find(&user, "token=?", token).Error; err != nil {
		return user, err
	}
	return user, nil
}
func UpdateStatusUser(status int, user *models.User) error {
	if err := DB.Model(&user).Update("status", status).Error; err != nil {
		return err
	}
	return nil
}
func GetUserByID(id string) (models.User, error) {
	var user models.User
	if err := DB.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}
func CreateUser(username string, password string, email string, name string, gender string, birthday string, phoneNumber int, status int, typeUser int, token string, confirm int) error {
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
	if err := DB.Save(&newUser).Error; err != nil {
		return err
	}
	return nil
}
func DeleteUser(user *models.User) error {
	if err := DB.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
func ConfirmRegisterUser(user *models.User) error {
	if err := DB.Model(&user).Update("confirm", constant.UserConfirmed).Error; err != nil {
		return err
	}
	return nil
}
