package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/hbl-duytv/intern-csm/models"
	"github.com/hbl-duytv/intern-csm/services"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		if user == nil {
			// You'd normally redirect to login page
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		} else {
			// Continue down the chain to handler etc
			c.Next()
		}
	}
}
func Login(c *gin.Context) {

	session := sessions.Default(c)
	username := c.PostForm("username")
	password := c.PostForm("password")

	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Parameters can't be empty"})
		return
	}
	// hash password to md5
	passwordMD5 := GetMD5Hash(password)
	user := models.User{}
	services.DB.Where("username = ? AND password = ?", username, passwordMD5).Find(&user)
	if user.ID != 0 {
		session.Set("user", username)
		err := session.Save()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusUnauthorized, "error": "Failed to generate session token"})
		} else {
			// c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Successfully authenticated user"})
			if user.Status == 1 {
				RenderHome(c)
			} else {
				message := []byte("Tài khoản chưa được kích hoạt, vui lòng đợi kích hoạt từ người quản trị!")
				c.Data(http.StatusOK, "text/html; charset=utf-8", message)
			}
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusNotFound, "error": "Authentication failed"})
	}

}
func RegisterSuccess(c *gin.Context) {
	username := c.Param("username")
	password := c.Param("password")
	email := c.Param("email")
	passwordMD5 := GetMD5Hash(password)
	if username != "" && password != "" && email != "" {
		newUser := models.User{
			Username: username,
			Password: passwordMD5,
			Email:    email,
			Type:     0,
			Status:   0,
		}
		services.DB.Save(&newUser)
		messageSuccess := []byte("Xác nhận tài khoản thành công, vui lòng đợi kích hoạt từ người quản trị!")
		c.Data(http.StatusOK, "text/html; charset=utf-8", messageSuccess)
	} else {
		messageFail := []byte("Xác nhận tài khoản thất bại, vui lòng thử lại!")
		c.Data(http.StatusOK, "text/html; charset=utf-8", messageFail)
	}
}
func SendConfirmRegister(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")
	urlConfirm := "http://localhost:8000/confirm-register/" + username + "/" + password + "/" + email
	massageEmailConfirm := "<div>Bạn đã đăng ký tài khoản biên tập viên, vui lòng xác nhận :</div><button><a href=\"" + urlConfirm + "\">Xác nhận đăng ký</a></button>"
	SendMail(email, massageEmailConfirm)
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("Gửi mail xác nhận thành công, vui lòng check mail để xác nhận đăng ký tài khoản!"))
}
func GetAllUserNotActive(c *gin.Context) {
	user := []models.User{}
	transformUSer := []models.TransformUser{}
	services.DB.Find(&user, "status=?", 0)
	if len(user) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user need active!"})
		return
	}
	for _, v := range user {
		transformUSer = append(transformUSer,
			models.TransformUser{v.ID, v.Username, v.Email, v.Type, v.Gender, v.BirthDay, v.PhoneNumber, v.Status})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": transformUSer})
}
func ConfirmUserAfterRegister(c *gin.Context) {
	var user models.User
	idUser := c.Param("id")
	services.DB.First(&user, idUser)
	status := 1
	services.DB.Model(&user).Update("status", status)
	services.DB.Save(&user)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Updated status user successfully!"})
}
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusNotFound, "error": "Invalid session token"})
	} else {
		log.Println(user)
		session.Delete("user")
		session.Save()
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Successfully logged out"})
	}
}
func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
func GetUserNotACtive() []models.TransformUser {
	user := []models.User{}
	transformUSer := []models.TransformUser{}
	services.DB.Find(&user, "status=?", 0)
	if len(user) < 0 {
		return nil
	}
	for _, v := range user {
		transformUSer = append(transformUSer,
			models.TransformUser{v.ID, v.Username, v.Email, v.Type, v.Gender, v.BirthDay, v.PhoneNumber, v.Status})
	}
	return transformUSer
}
func CheckUserExist(c *gin.Context) {
	username := c.PostForm("username")
	user := models.User{}
	services.DB.Where("username = ?", username).Find(&user)
	if user.ID == 0 {
		c.JSON(http.StatusOK, gin.H{"check": true, "message": "Successfully check user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"check": false, "message": "User exist!"})
	return
}
func CheckEmailExist(c *gin.Context) {
	email := c.PostForm("email")
	user := models.User{}
	services.DB.Where("email = ?", email).Find(&user)
	if user.ID == 0 {
		c.JSON(http.StatusOK, gin.H{"check": true, "message": "Successfully check email"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"check": false, "message": "Email exist!"})
	return
}
func RenderIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "Login"})
}
