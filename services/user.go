package services

import (
	"time"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/hbl-duytv/intern-csm/constant"
	"github.com/hbl-duytv/intern-csm/helper"
	"github.com/hbl-duytv/intern-csm/models"
)

var SessionName interface{}

func GetAllEditorUser() ([]models.User, error) {
	var users []models.User
	if err := DB.Debug().Where("type=? AND confirm=?", constant.DeactiveNumber, constant.UserConfirmed).Find(&users).Error; err != nil {
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
	if err := DB.Model(&user).Update("confirm", constant.ActiveNumber).Error; err != nil {
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
func CreateAccount(newUser *models.User) error {
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

// User must click link in email to confirm account in 8 hours from creating account
func HasLimitTimeConfirm(user models.User) bool {
	timeConfirm := user.Confirm * 60 * 60
	timeCheck := user.CreatedAt.Add(time.Duration(timeConfirm) * time.Second)
	now := time.Now()
	if now.Sub(timeCheck).Seconds() > 0 {
		return false
	}
	return true

}

func GetCurrentUser(c *gin.Context) interface{} {
	session := sessions.Default(c)
	// session.Set("user", SessionName)
	return session.Get("user")
}

func GetNameInCommentByPostID(idPost int) ([]string, error) {
	var names []string
	rows, err := DB.Table("user").Select("user.name").Joins("join comment on comment.commentator_id = user.id where comment.post_id=?", idPost).Rows()
	if err != nil {
		return names, err
	}
	for rows.Next() {
		var name string
		rows.Scan(&name)
		names = append(names, name)
	}
	return names, nil

}
