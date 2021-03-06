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
	router.GET("/login", controllers.Index)
	router.GET("/home", controllers.Home)
	router.GET("/editor-management2", controllers.RenderEditorManagement)
	router.GET("/confirm-register/:token", controllers.RegisterSuccess)
	router.GET("/logout", controllers.Logout)
	router.GET("/get-post-admin-permission3", controllers.RenderPostManagementAdmin)
	router.GET("/get-post-editor-permission", controllers.RenderPostManagementEditor)
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

	router.GET("/blog", controllers.Blog)
	router.GET("/blog/:postID", controllers.BlogDetailPost)
	router.GET("/total-page", controllers.GetTotalNumberAllPost)
	privateRouter := router.Group("/")
	{
		privateRouter.POST("/active-editor", controllers.ActiveEditorUser)
		privateRouter.POST("/deactive-editor", controllers.DeactiveEditorUser)
		privateRouter.POST("/delete-user", controllers.DeleteUser)
		privateRouter.POST("/delete-post", controllers.DeletePost)
		privateRouter.GET("/render-create-post", controllers.RenderCreatePost)
		privateRouter.GET("/render-update-post/:id", controllers.RenderUpdatePost)
		privateRouter.GET("/render-detail-post/:id", controllers.RenderDetailPost)
		privateRouter.POST("/create-user", controllers.CreateUser)
	}
	privateRouter.Use(middleware.AuthAdminRequired())
}
