package router

import (
	"github.com/gin-gonic/gin"
	"net-chat/controller"
)

func InitWsRouter(r *gin.RouterGroup) {

	ws := r.Group("/ws")

	ws.GET("/connect", controller.WsConnect)

}
