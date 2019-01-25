package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/hbl-duytv/intern-csm/models"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/hbl-duytv/intern-csm/constant"
	"github.com/hbl-duytv/intern-csm/helper"
	"github.com/hbl-duytv/intern-csm/services"
)

func Login(c *gin.Context) {
	session := sessions.Default(c)
	username := c.PostForm("username")
	password := c.PostForm("password")
	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Parameters can't be empty"})
		return
	}
	user, err := services.RequireLogin(username, password)
	if err == nil {
		session.Set("user", username)
		// services.SessionName = session.Get("user")
		err := session.Save()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusUnauthorized, "error": "Failed to generate session token"})
			return
		}
		if user.Status == constant.ActiveNumber && user.Confirm == constant.UserConfirmed {
			c.Redirect(http.StatusMovedPermanently, "/home")
			return
		}
		message := []byte("Tài khoản chưa được kích hoạt, vui lòng đợi kích hoạt từ người quản trị!")
		c.Data(http.StatusOK, "text/html; charset=utf-8", message)
		return

	}
	c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusNotFound, "error": "Authentication failed"})

}
func RegisterSuccess(c *gin.Context) {
	var messageSuccess []byte
	token := c.Param("token")
	if user, err := services.GetUserByToken(token); err == nil {
		if user.Confirm == constant.UserNotConfirmed {
			if services.HasLimitTimeConfirm(user) {
				messageSuccess = []byte("Xác nhận tài khoản thành công, vui lòng đợi kích hoạt từ người quản trị!")
				services.ConfirmRegisterUser(&user)
			} else {
				messageSuccess = []byte("Hết thời gian kích hoạt tài khoản!")
			}
		} else {
			messageSuccess = []byte("Tài khoản đã được xác nhận, vui lòng đợi kích hoạt từ người quản trị!")
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", messageSuccess)
	}
	messageFail := []byte("Xác nhận tài khoản thất bại, vui lòng thử lại!")
	c.Data(http.StatusOK, "text/html; charset=utf-8", messageFail)
}

func SendConfirmRegister(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")
	token := helper.GetToken(username)
	newUser := models.User{
		Username:    username,
		Password:    password,
		Email:       email,
		Name:        "",
		Gender:      "",
		Birthday:    "",
		PhoneNumber: constant.DeactiveNumber,
		Status:      constant.DeactiveNumber,
		Type:        constant.TypeEditor,
		Token:       token,
		Confirm:     constant.DeactiveNumber,
	}
	if err := services.CreateAccount(&newUser); err == nil {
		urlConfirm := "http://localhost:8000/confirm-register/" + token
		massageEmailConfirm := "<div>Bạn đã đăng ký tài khoản biên tập viên, vui lòng xác nhận :</div><a href= '" + urlConfirm + "'><button>Xác nhận đăng ký</button></a>"
		if err := services.SendMail(email, massageEmailConfirm); err == nil {
			c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("Gửi mail xác nhận thành công, vui lòng check mail để xác nhận đăng ký tài khoản!"))
			return
		}
		c.Data(http.StatusBadRequest, "text/html; charset=utf-8", []byte("Gửi mail xác nhận thất bại!"))

	}

}
func ConfirmUserAfterRegister(c *gin.Context) {
	if idUser, err := strconv.Atoi(c.Param("id")); err == nil {
		if services.UpdateStatusUser(idUser, constant.ActiveNumber) == nil {
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Updated status user successfully!"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Updated status user failed!"})
	}
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		c.Redirect(http.StatusMovedPermanently, "login")
		return
	}
	session.Delete("user")
	session.Save()
	c.Redirect(http.StatusMovedPermanently, "login")
}
func CheckUserExist(c *gin.Context) {
	username := c.PostForm("username")
	if _, err := services.GetUserByUsername(username); err == nil {

		c.JSON(http.StatusOK, gin.H{"check": true, "message": "Successfully check user"})
		return

	}

	c.JSON(http.StatusOK, gin.H{"check": false, "message": "User exist!"})
}
func CheckEmailExist(c *gin.Context) {
	email := c.PostForm("email")
	if _, err := services.GetUserByEmail(email); err == nil {

		c.JSON(http.StatusOK, gin.H{"check": true, "message": "Successfully check email"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"check": false, "message": "Email exist!"})
}

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "Login"})
}
