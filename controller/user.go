package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net-chat/global"
	"net-chat/logic"
	"net-chat/pkg"
	"net/http"
)

func Login(c *gin.Context) {
	var (
		r = pkg.NewResponse(c)
	)
	req := struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}{}

	if err := c.ShouldBindJSON(&req); err != nil {
		r.Fail(err)
		return
	}
	token, err := logic.NewUserLogic().Login(global.DB.WithContext(c.Request.Context()), req.Username, req.Password)
	if err != nil {
		r.Fail(err)
		return
	}

	r.OK(http.StatusOK, gin.H{"token": token})

}

func UserInfo(c *gin.Context) {
	var (
		r   = pkg.NewResponse(c)
		uid = c.GetInt64(global.UserID)
	)

	user, err := logic.NewUserLogic().UserInfo(global.DB.WithContext(c.Request.Context()), uid)
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
