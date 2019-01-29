package controllers

import (
	"net/http"
	"strconv"
	"strings"

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
		err := session.Save()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusUnauthorized, "error": "Failed to generate session token"})
			return
		}
		if user.Status == constant.ACTIVE_NUMBER {
			c.Redirect(constant.DIRECT_STATUS, "/home")
		} else {
			message := []byte("Tài khoản chưa được kích hoạt, vui lòng đợi kích hoạt từ người quản trị!")
			c.Data(http.StatusOK, "text/html; charset=utf-8", message)
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusNotFound, "error": "Authentication failed"})
	}

}
func RegisterSuccess(c *gin.Context) {
	var messageSuccess []byte
	token := c.Param("token")
	if user, err := services.GetUserByToken(token); err == nil {
		// fmt.Println("inf: ", user.Username, user.Confirm)
		if user.Confirm == constant.DEACTIVE_NUMBER {
			if services.CheckTimeToConfirmUser(user) {
				messageSuccess = []byte("Xác nhận tài khoản thành công, vui lòng đợi kích hoạt từ người quản trị!")
				// fmt.Println("confirm first:", user.Confirm)
				services.ConfirmRegisterUser(&user)
			} else {
				messageSuccess = []byte("Hết thời gian kích hoạt tài khoản!")
			}

		} else {
			// fmt.Println("confirm then", user.Confirm)
			messageSuccess = []byte("Tài khoản đã được xác nhận, vui lòng đợi kích hoạt từ người quản trị!")

		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", messageSuccess)
	}
}
func SendConfirmRegister(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")
	token := helper.GetToken(username)
	if services.CreateAccount(username, password, email, "", "", "", constant.DEACTIVE_NUMBER, constant.DEACTIVE_NUMBER, constant.DEACTIVE_NUMBER, token, constant.DEACTIVE_NUMBER) == nil {
		urlConfirm := "http://localhost:8000/confirm-register/" + token
		massageEmailConfirm := "<div>Bạn đã đăng ký tài khoản biên tập viên, vui lòng xác nhận :</div><a href= '" + urlConfirm + "'><button>Xác nhận đăng ký</button></a>"
		services.SendMail(email, massageEmailConfirm)
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("Gửi mail xác nhận thành công, vui lòng check mail để xác nhận đăng ký tài khoản!"))
	}

}
func ConfirmUserAfterRegister(c *gin.Context) {
	idUser, _ := strconv.Atoi(c.Param("id"))

	if services.UpdateStatusUser(idUser, constant.ACTIVE_NUMBER) == nil {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Updated status user successfully!"})
	} else {

	}

}
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		c.Redirect(constant.DIRECT_STATUS, "login")
		return
	}
	session.Delete("user")
	session.Save()
	c.Redirect(constant.DIRECT_STATUS, "login")
}
func CheckUserExist(c *gin.Context) {
	username := c.PostForm("username")
	if user, err := services.GetUserByUsername(username); err == nil {
		if user.ID == constant.DEACTIVE_NUMBER {
			c.JSON(http.StatusOK, gin.H{"check": true, "message": "Successfully check user"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"check": false, "message": "User exist!"})
	return
}
func CheckEmailExist(c *gin.Context) {
	email := c.PostForm("email")
	if user, err := services.GetUserByEmail(email); err == nil {
		if user.ID == constant.DEACTIVE_NUMBER {
			c.JSON(http.StatusOK, gin.H{"check": true, "message": "Successfully check email"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"check": false, "message": "Email exist!"})
	return
}
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "Login"})
}
