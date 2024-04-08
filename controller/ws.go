package controller

import (
	"github.com/gin-gonic/gin"
	"net-chat/pkg/ws"
)

func WsConnect(c *gin.Context) {

	uid := c.GetInt64("uid")

	ws.ServeWs(c.Writer, c.Request, uid)
}
