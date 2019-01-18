package routers

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/hbl-duytv/intern-csm/controllers"
)

func InitRouter(router *gin.Engine) {
	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	router.POST("/login", controllers.Login)
	router.GET("/logout", controllers.Logout)
	router.POST("/register", controllers.SendConfirmRegister)
	router.POST("/check-user-exist", controllers.CheckUserExist)
	router.POST("/check-email-exist", controllers.CheckEmailExist)
	router.GET("/confirm-register/:username/:password/:email", controllers.RegisterSuccess)
	router.POST("/confirm-user-after-register/:id", controllers.ConfirmUserAfterRegister)
	// router.GET("/admin", controllers.GetAllUserNotActive)
	privateRouter := router.Group("/api")
	{
		privateRouter.GET("/home")
		privateRouter.GET("/posts")
	}
	privateRouter.Use(controllers.AuthRequired())
}
