package initialize

import (
	"github.com/gin-gonic/gin"
	"net-chat/global"
	"net-chat/middleware"
	"net-chat/router"
)

func Engin() *gin.Engine {

	gin.SetMode(global.Config.Runtime.Mode)

	engine := gin.New()

	engine.Use(gin.Recovery())
	if gin.Mode() != gin.ReleaseMode {
		engine.Use(gin.Logger())
		engine.Use(middleware.CORSMiddleware())
	}

	apiV1 := engine.Group(global.Config.System.PrefixUrl)

	apiV1.Use(middleware.TraceMiddleware())

	apiV1.Use(middleware.PanicHandler())

	apiV1.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
		return
	})

	router.InitUserRouter(apiV1)
	router.InitWsRouter(apiV1)

	return engine
}
