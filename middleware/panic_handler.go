package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net-chat/global"
	"net/http"
	"runtime"
	"strings"
)

func PanicHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {

				errResp := gin.H{
					"status":  500,
					"success": false,
					"error":   err,
				}
				if global.Config.Runtime.Mode != gin.ReleaseMode {
					e := fmt.Sprintf("%#v \n %s", err, PrintStack())
					errResp["stack"] = strings.Split(e, "\n")
				}
				c.JSON(http.StatusOK, errResp)
				return
			}
		}()
		c.Next()
	}
}

func PrintStack() string {
	buf := make([]byte, 1024*64)
	n := runtime.Stack(buf, false)
	buf = buf[:n]
	return fmt.Sprintln(string(buf[:n]))
}
