package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/hbl-duytv/intern-csm/helper"
	"github.com/hbl-duytv/intern-csm/models"
	"github.com/hbl-duytv/intern-csm/services"
)

const UserActived = 1
const UserDeactived = 0
const UserConfirmed = 1
const UserNotConfirmed = 0

func Login(c *gin.Context) {

	session := sessions.Default(c)
	username := c.PostForm("username")
	password := c.PostForm("password")
	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Parameters can't be empty"})
		return
	}
	passwordMD5 := helper.GetMD5Hash(password)
	user := models.User{}
	services.DB.Where("username = ? AND password = ?", username, passwordMD5).Find(&user)
	if user.ID != 0 {
		session.Set("user", username)
		err := session.Save()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusUnauthorized, "error": "Failed to generate session token"})
			return
		}
		if user.Status == UserActived && user.Confirm == 1 {
			c.Redirect(301, "/home")
			return
		}
		message := []byte("Tài khoản chưa được kích hoạt, vui lòng đợi kích hoạt từ người quản trị!")
		c.Data(http.StatusOK, "text/html; charset=utf-8", message)
	} else {
		c.Redirect(301, "/login")
	}
}
func RegisterSuccess(c *gin.Context) {
	token := c.Param("token")
	if token != "" {
		user := services.GetUserByToken(token)
		if user.ID != 0 {
			if user.Confirm == UserConfirmed {
				services.ConfirmRegisterUser(&user)
				messageSuccess := []byte("Xác nhận tài khoản thành công, vui lòng đợi kích hoạt từ người quản trị!")
				c.Data(http.StatusOK, "text/html; charset=utf-8", messageSuccess)
				return
			}
			messageSuccess := []byte("Tài khoản đã được xác nhận, vui lòng đợi kích hoạt từ người quản trị!")
			c.Data(http.StatusOK, "text/html; charset=utf-8", messageSuccess)
			return
		}
	}
	messageFail := []byte("Xác nhận tài khoản thất bại, vui lòng thử lại!")
	c.Data(http.StatusOK, "text/html; charset=utf-8", messageFail)
}
func SendConfirmRegister(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")
	token := helper.GetToken(username)
	urlConfirm := "http://localhost:8000/confirm-register/" + token
	services.CreateUser(username, password, email, "", "", "", 0, 0, 0, token, 0)
	massageEmailConfirm := "<div>Bạn đã đăng ký tài khoản biên tập viên, vui lòng xác nhận :</div><button><a href=\"" + urlConfirm + "\">Xác nhận đăng ký</a></button>"
	SendMail(email, massageEmailConfirm)
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("Gửi mail xác nhận thành công, vui lòng check mail để xác nhận đăng ký tài khoản!"))
}
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		c.Redirect(301, "/login")
		return
	}
	session.Delete("user")
	session.Save()
	c.Redirect(301, "/login")
}
func CheckUserExist(c *gin.Context) {
	username := c.PostForm("username")
	user := services.GetUserByUsername(username)
	if user.ID == 0 {
		c.JSON(http.StatusOK, gin.H{"check": true, "message": "Successfully check user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"check": false, "message": "User exist!"})
}
func CheckEmailExist(c *gin.Context) {
	email := c.PostForm("email")
	user := services.GetUserByEmail(email)
	if user.ID == 0 {
		c.JSON(http.StatusOK, gin.H{"check": true, "message": "Successfully check email"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"check": false, "message": "Email exist!"})
}
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "Login"})
}
