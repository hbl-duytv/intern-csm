package services

import (
	"time"

	"github.com/hbl-duytv/intern-csm/constant"
	"github.com/hbl-duytv/intern-csm/helper"
	"github.com/hbl-duytv/intern-csm/models"
)

func GetAllEditorUser() ([]models.User, error) {
	var users []models.User
	if err := DB.Debug().Where("type=? AND confirm=?", constant.DEACTIVE_NUMBER, constant.ACTIVE_NUMBER).Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}
func GetUserByToken(token string) (models.User, error) {
	var user models.User
	if err := DB.Where("token=?", token).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
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
func ConfirmRegisterUser(user *models.User) error {
	if err := DB.Model(&user).Update("confirm", constant.ACTIVE_NUMBER).Error; err != nil {
		return err
	}
	return nil
}
func GetUserByID(id int) (models.User, error) {
	var user models.User
	if err := DB.Find(&user, "id = ?", id).Error; err != nil {
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
	return user, nil
}
func CreateAccount(username string, password string, email string, name string, gender string, birthday string, phoneNumber int, status int, typeUser int, token string, confirm int) error {
	passwordMD5 := helper.GetMD5Hash(password)

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
func GetTimeCreateUSer(id int) (int, int, error) {
	var user models.User
	if err := DB.Where("id=?", id).Find(&user).Error; err != nil {
		return int(user.CreatedAt.Month()), int(user.CreatedAt.Year()), err
	}
	return int(user.CreatedAt.Month()), int(user.CreatedAt.Year()), nil
}
func CheckTimeToConfirmUser(user models.User) bool {
	timeCheck := user.CreatedAt.Add(1 * time.Minute)
	now := time.Now()
	if now.Sub(timeCheck).Minutes() > 1 {
		return false
	}
	return true

}
