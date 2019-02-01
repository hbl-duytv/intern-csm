package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/hbl-duytv/intern-csm/constant"
	"github.com/hbl-duytv/intern-csm/helper"
	"github.com/hbl-duytv/intern-csm/services"
)

var session_name interface{}

func Login(c *gin.Context) {

	session := sessions.Default(c)
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println("username:", username)
	fmt.Println("password:", password)
	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Parameters can't be empty"})
		return
	}
	user, err := services.RequireLogin(username, password)
	if err == nil {

		session.Set("user", username)
		session_name = session.Get("user")
		fmt.Println("get:", session.Get("user"))
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
	fmt.Println("token: ", token)
	if user, err := services.GetUserByToken(token); err == nil {

		if user.Confirm == constant.DEACTIVE_NUMBER {
			//messageSuccess = []byte("Xác nhận tài khoản thành công, vui lòng đợi kích hoạt từ người quản trị!")

			//services.ConfirmRegisterUser(user.ID)
			if check, err := services.CheckTimeToConfirmUser(user.ID); err != nil {
				messageSuccess = []byte("Có lỗi xảy ra khi đăng kí!")
				c.Data(http.StatusBadRequest, "text/html; charset=utf-8", messageSuccess)
			} else {
				if !check {
					fmt.Println("log1")
					messageSuccess = []byte("Xác nhận tài khoản thành công, vui lòng đợi kích hoạt từ người quản trị!")
					c.Data(http.StatusOK, "text/html; charset=utf-8", messageSuccess)
					services.ConfirmRegisterUser(user.ID)
				} else {
					fmt.Println("log2")
					messageSuccess = []byte("Hết thời gian xác thực tài khoản")
					c.Data(http.StatusBadRequest, "text/html; charset=utf-8", messageSuccess)
				}
			}
		} else {
			fmt.Println("log3")
			messageSuccess = []byte("Tài khoản đã được xác nhận, vui lòng đợi kích hoạt từ người quản trị!")
		}

	} else {
		messageSuccess = []byte("Vui lòng đăng kí lại")
		c.Data(http.StatusBadRequest, "text/html; charset=utf-8", messageSuccess)
	}
}
func SendConfirmRegister(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")
	token := helper.GetToken(username)
	if services.CreateAccount(username, password, email, "", "", "", constant.DEACTIVE_NUMBER, constant.DEACTIVE_NUMBER, constant.DEACTIVE_NUMBER, token) == nil {
		urlConfirm := "http://localhost:8000/confirm-register/" + token
		massageEmailConfirm := "<div>Bạn đã đăng ký tài khoản biên tập viên, vui lòng xác nhận :</div><a href= '" + urlConfirm + "'><button>Xác nhận đăng ký</button></a>"
		if services.SendMail(email, massageEmailConfirm) == nil {
			c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("Gửi mail xác nhận thành công, vui lòng check mail để xác nhận đăng ký tài khoản!"))
		} else {
			c.Data(http.StatusBadRequest, "text/html; charset=utf-8", []byte("Gửi mail xác nhận thất bại, vui lòng kiểm tra lại!"))
		}

	}
}
func ConfirmUserAfterRegister(c *gin.Context) {
	idUser, _ := strconv.Atoi(c.Param("id"))
	if services.UpdateStatusUser(idUser, constant.ACTIVE_NUMBER) == nil {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Updated status user successfully!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Updated status user fail!"})
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
	fmt.Println("username: ", username)
	if _, err := services.GetUserByUsername(username); err == nil {
		// if user.ID == constant.DEACTIVE_NUMBER {
		c.JSON(http.StatusOK, gin.H{"check": true, "message": "Successfully check user"})
		// 	return
		// }
	}

	c.JSON(http.StatusBadRequest, gin.H{"check": false, "message": "User exist!"})
	return
}
func CheckEmailExist(c *gin.Context) {
	email := c.PostForm("email")
	fmt.Println("email: ", email)
	if _, err := services.GetUserByEmail(email); err == nil {
		// if user.ID == constant.DEACTIVE_NUMBER {
		c.JSON(http.StatusOK, gin.H{"check": true, "message": "Successfully check email"})
		// 	return
		// }
	}
	c.JSON(http.StatusBadRequest, gin.H{"check": false, "message": "Email exist!"})
	return
}
func Index(c *gin.Context) {
	fmt.Println("index")
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "Login"})
}
