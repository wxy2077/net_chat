package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net-chat/global"
	"net-chat/logic"
	"net-chat/pkg"
	"net/http"
)

func UserInfo(c *gin.Context) {
	var (
		r = pkg.NewResponse(c)
	)
	req := struct {
		ID int64 `json:"id" form:"id" binding:"required"`
	}{}

	if err := c.ShouldBindQuery(&req); err != nil {
		r.Fail(err)
		return
	}
	user, err := logic.NewUserLogic().UserInfo(global.DB.WithContext(c.Request.Context()), req.ID)
	if err != nil {
		r.Fail(err)
		return
	}

	r.OK(http.StatusOK, user)
}

func JoinFunc(c *gin.Context) {

	r := pkg.NewResponse(c)

	list, pagination := logic.NewUserLogic().UserDepList(global.DB.WithContext(c.Request.Context()), 1)

	global.Log.WithContext(c.Request.Context()).
		WithFields(logrus.Fields{"list": "dadsdasda"}).
		Error(33333)

	r.OkWithPage(http.StatusOK, list, pagination)
}

func PreloadFunc(c *gin.Context) {
	var (
		r = pkg.NewResponse(c)
	)

	list, pagination := logic.NewUserLogic().PreloadUserDep(global.DB.WithContext(c.Request.Context()), 1)

	r.OkWithPage(http.StatusOK, list, pagination)
}

func PreloadsFunc(c *gin.Context) {
	var (
		r = pkg.NewResponse(c)
	)
	list, pagination := logic.NewUserLogic().PreloadUserDeps(global.DB.WithContext(c.Request.Context()), 1)

	r.OkWithPage(http.StatusOK, list, pagination)
}
