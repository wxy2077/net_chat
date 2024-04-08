package router

import (
	"github.com/gin-gonic/gin"
	"net-chat/controller"
	"net-chat/middleware"
)

func InitWsRouter(r *gin.RouterGroup) {

	ws := r.Group("/ws")

	ws.Use(middleware.JwtMiddleware())

	ws.GET("/connect", controller.WsConnect)

}
