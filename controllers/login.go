package controllers

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/hbl-duytv/intern-csm/helper"
	"github.com/hbl-duytv/intern-csm/models"
	"github.com/hbl-duytv/intern-csm/services"
)

const activeNumber int = 1
const deactiveNumber int = 0

func Login(c *gin.Context) {

	session := sessions.Default(c)
	username := c.PostForm("username")
	password := c.PostForm("password")

	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Parameters can't be empty"})
		return
	}
	// hash password to md5
	passwordMD5 := helper.GetMD5Hash(password)
	currentUser := models.User{}
	services.DB.Where("username = ? AND password = ?", username, passwordMD5).Find(&currentUser)
	if currentUser.ID != deactiveNumber {
		session.Set("user", username)
		err := session.Save()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusUnauthorized, "error": "Failed to generate session token"})
			c.Redirect(301, "/login")
		} else {
			if currentUser.Status == activeNumber {
				c.Redirect(301, "/home")
				// RenderHome(c, user)
			} else {
				message := []byte("Tài khoản chưa được kích hoạt, vui lòng đợi kích hoạt từ người quản trị!")
				c.Data(http.StatusOK, "text/html; charset=utf-8", message)
			}
		}
	} else {
		// c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusNotFound, "error": "Authentication failed"})
		c.Redirect(301, "/login")
		// c.HTML(http.StatusUnauthorized, "index.html", gin.H{"status": http.StatusUnauthorized})
		// c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized})
	}
}
func RegisterSuccess(c *gin.Context) {
	username := c.Param("username")
	password := c.Param("password")
	email := c.Param("email")
	passwordMD5 := helper.GetMD5Hash(password)
	if username != "" && password != "" && email != "" {
		newUser := models.User{}
		services.DB.Where("username = ?", username).Find(&newUser)
		if newUser.ID == deactiveNumber {
			newUser := models.User{
				Username: username,
				Password: passwordMD5,
				Name:     "",
				Email:    email,
				Type:     deactiveNumber,
				Status:   deactiveNumber,
			}
			services.DB.Create(&newUser)
			messageSuccess := []byte("Xác nhận tài khoản thành công, vui lòng đợi kích hoạt từ người quản trị!")
			c.Data(http.StatusOK, "text/html; charset=utf-8", messageSuccess)
		} else {
			messageSuccess := []byte("Tài khoản đã được đăng ký!")
			c.Data(http.StatusOK, "text/html; charset=utf-8", messageSuccess)
		}
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
	var user []models.User

	services.DB.Find(&user, "status=?", deactiveNumber)
	if len(user) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user need active!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": user})
}
func ConfirmUserAfterRegister(c *gin.Context) {
	var user models.User
	idUser := c.Param("id")
	services.DB.First(&user, idUser)
	status := 1
	services.DB.Model(&user).Update("status", status)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Updated status user successfully!"})
}
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		c.Redirect(301, "/login")
		// c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusNotFound, "error": "Invalid session token"})
	} else {
		log.Println(user)
		session.Delete("user")
		session.Save()
		c.Redirect(301, "/login")
		// c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Successfully logged out"})
	}
}

func GetUserNotActive() []models.User {

	var user []models.User

	services.DB.Find(&user, "status=?", deactiveNumber)
	if len(user) < 0 {
		return nil
	}

	return user
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

	var user models.User
	services.DB.Where("email = ?", email).Find(&user)
	if user.ID == 0 {
		c.JSON(http.StatusOK, gin.H{"check": true, "message": "Successfully check email"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"check": false, "message": "Email exist!"})
	return
}
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "Login"})
}
