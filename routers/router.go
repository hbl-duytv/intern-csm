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
	router.GET("/confirm-register/:username/:password/:email", controllers.RegisterSuccess)
	router.GET("/logout", controllers.Logout)
	router.GET("/get-post-admin-permission", controllers.GetPostWithAdminPermission)
	router.GET("/get-post-editor-permission", controllers.GetPostWithEditorPermission)
	// post api router
	router.POST("/register", controllers.SendConfirmRegister)
	router.POST("/check-user-exist", controllers.CheckUserExist)
	router.POST("/check-email-exist", controllers.CheckEmailExist)

	router.POST("/confirm-user-after-register/:id", controllers.ConfirmUserAfterRegister)
	router.POST("/login", controllers.Login)
	router.POST("/active-status-post", controllers.ActiveStatusPost)
	router.POST("/deactive-status-post", controllers.DeActiveStatusPost)
	router.POST("/create-post", controllers.CreatePost)
	router.POST("/update-content-post", controllers.UpdateContentPost)
	privateRouter := router.Group("/")
	{
		privateRouter.POST("/active-editor", controllers.ActiveEditorUser)
		privateRouter.POST("/deactive-editor", controllers.DeactiveEditorUser)
		privateRouter.POST("/delete-user", controllers.DeleteUser)
		privateRouter.POST("/delete-post", controllers.DeletePost)
		privateRouter.GET("/editor-management", controllers.RenderEditorManagement)
		privateRouter.GET("/render-create-post", controllers.RenderCreatePost)
		privateRouter.GET("/render-update-post/:id", controllers.RenderUpdatePost)
		privateRouter.GET("/render-detail-post/:id", controllers.RenderDetailPost)
	}
	privateRouter.Use(controllers.AuthAdminRequired())
}
