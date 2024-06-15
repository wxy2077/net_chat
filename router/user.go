package router

import (
	"github.com/gin-gonic/gin"
	"net-chat/controller"
	"net-chat/middleware"
)

func InitUserRouter(r *gin.RouterGroup) {

	r.POST("/user/login", controller.Login)

	user := r.Group("/user")
	user.Use(middleware.JwtMiddleware())

	user.GET("/info", controller.UserInfo)

	user.GET("/friend-list", controller.FriendList)

	user.GET("/message", controller.Message)
}
