package controller

import (
	"github.com/gin-gonic/gin"
	"net-chat/pkg/ws"
)

func WsConnect(c *gin.Context) {

	ws.ServeWs(c.Writer, c.Request, 1)
}
