package routers

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/hbl-duytv/intern-csm/controllers"
)

func InitRouter(router *gin.Engine) {
	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	router.GET("/", controllers.RenderIndex)
	router.GET("/login", controllers.RenderIndex)
	router.GET("/home", controllers.RenderHome)
	router.POST("/login", controllers.Login)
	router.GET("/logout", controllers.Logout)
	// post api router
	router.POST("/register", controllers.SendConfirmRegister)
	router.POST("/check-user-exist", controllers.CheckUserExist)
	router.POST("/check-email-exist", controllers.CheckEmailExist)
	router.GET("/confirm-register/:username/:password/:email", controllers.RegisterSuccess)
	router.POST("/confirm-user-after-register/:id", controllers.ConfirmUserAfterRegister)
	privateRouter := router.Group("/")
	{
		privateRouter.POST("/active-editor", controllers.ActiveEditorUser)
		privateRouter.POST("/deactive-editor", controllers.DeactiveEditorUser)
		privateRouter.POST("/delete-user", controllers.DeleteUser)
		privateRouter.POST("/create-user", controllers.CreateUser)
		privateRouter.GET("/editor-management", controllers.RenderEditorManagement)
	}
	privateRouter.Use(controllers.AuthAdminRequired())
}
