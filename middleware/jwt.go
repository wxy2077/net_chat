package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"net-chat/global"
	"net-chat/pkg"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.Request.Header.Get("Authorization")
		if token == "" {
			token = c.Query("token")
		}
		if token == "" {
			pkg.NewResponse(c).Fail(errors.New("请先登录"))
			c.Abort()
			return
		}

		claims, err := pkg.PwdJwt.ParseToken(token)
		if err != nil {
			pkg.NewResponse(c).Fail(err)
			c.Abort()
			return
		}

		uid, ok := claims[global.UserID]
		if ok {
			c.Set(global.UserID, cast.ToInt64(uid))
		}

		c.Next()
	}
}
