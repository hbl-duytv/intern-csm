package services

import (
	"fmt"

	"github.com/hbl-duytv/intern-csm/constant"
	"github.com/hbl-duytv/intern-csm/helper"
	"github.com/hbl-duytv/intern-csm/models"
)

func GetAllEditorUser() ([]models.User, error) {
	var users []models.User
	if err := DB.Find(&users, "type=?", constant.DEACTIVE_NUMBER).Error; err != nil {
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
func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	if err := DB.Find(&user, "email=?", email).Error; err != nil {
		return user, err
	}

	return user, nil
}
func UpdateStatusUser(id, status int) error {
	var user models.User
	if err := DB.First(&user, id).Error; err != nil {
		return err
	}
	if err := DB.Model(&user).Update("status", status).Error; err != nil {
		return err
	}
	return nil

}
func GetUserByID(id string) (models.User, error) {
	var user models.User
	if err := DB.Debug().Find(&user, "id = ?", id).Error; err != nil {
		fmt.Println(err)
		return user, err
	}

	return user, nil
}
func CreateUser(username string, password string, email string, status int, typeUser int) error {
	passwordMD5 := helper.GetMD5Hash(password)
	newUser := models.User{
		Username: username,
		Password: passwordMD5,
		Email:    email,
		Type:     typeUser,
		Status:   status,
	}
	if err := DB.Save(&newUser).Error; err != nil {
		return err
	}
	return nil

}
func DeleteUser(id int) error {
	var user models.User
	if err := DB.First(&user, id).Error; err != nil {
		return err
	}
	if err := DB.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
func RequireLogin(username, password string) (models.User, error) {
	var user models.User
	passwordMD5 := helper.GetMD5Hash(password)
	if err := DB.Where("username = ? AND password = ?", username, passwordMD5).Find(&user).Error; err != nil {
		return user, err
	}
	// DB.Where("username = ? AND password = ?", username, passwordMD5).Find(&user)
	return user, nil

}
