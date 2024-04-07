package router

import (
	"github.com/gin-gonic/gin"
	"net-chat/controller"
)

func InitUserRouter(r *gin.RouterGroup) {

	user := r.Group("/user")

	user.POST("/login", controller.Login)

	user.GET("/info", controller.UserInfo)
	user.GET("/join", controller.JoinFunc)
	user.GET("/preload", controller.PreloadFunc)
	user.GET("/preloads", controller.PreloadsFunc)
}
