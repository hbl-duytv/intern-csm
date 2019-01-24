package routers

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/hbl-duytv/intern-csm/controllers"
	"github.com/hbl-duytv/intern-csm/middleware"
)

func InitRouter(router *gin.Engine) {
	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	router.GET("/", controllers.Index)
	router.GET("/home", controllers.Home)
	router.POST("/login", controllers.Login)
	router.GET("/logout", controllers.Logout)
	router.POST("/register", controllers.SendConfirmRegister)
	router.POST("/check-user-exist", controllers.CheckUserExist)
	router.POST("/check-email-exist", controllers.CheckEmailExist)
	router.GET("/confirm-register/:username/:password/:email", controllers.RegisterSuccess)
	router.POST("/confirm-user-after-register/:id", controllers.ConfirmUserAfterRegister)
	privateRouter := router.Group("/api")
	{
		privateRouter.GET("/posts")
	}
	privateRouter.Use(middleware.AuthRequired())
}
