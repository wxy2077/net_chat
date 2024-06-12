package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
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
		r.Fail(errors.New("账号密码必填"))
		return
	}
	token, err := logic.NewUserLogic().Login(global.DB.WithContext(c.Request.Context()), req.Username, req.Password)
	if err != nil {
		global.Log.WithContext(c.Request.Context()).
			WithFields(logrus.Fields{"username": req.Username}).
			Error("login error")

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

func FriendList(c *gin.Context) {
	var (
		r   = pkg.NewResponse(c)
		uid = c.GetInt64(global.UserID)
		req = struct {
			Page     int64 `form:"page"`
			PageSize int64 `form:"page_size"`
		}{
			Page:     pkg.DefaultPage,
			PageSize: pkg.DefaultPageSize,
		}
	)
	_ = c.ShouldBindQuery(&req)

	list, pagination := logic.NewUserLogic().FriendList(global.DB.WithContext(c.Request.Context()), uid, req.Page, req.PageSize)

	r.OkWithPage(http.StatusOK, list, pagination)
}

func Message(c *gin.Context) {
	var (
		r   = pkg.NewResponse(c)
		uid = c.GetInt64(global.UserID)
		req = struct {
			FriendUserID int64 `form:"friend_user_id"`
			Page         int64 `form:"page"`
			PageSize     int64 `form:"page_size"`
		}{
			Page:     pkg.DefaultPage,
			PageSize: pkg.DefaultPageSize,
		}
	)
	if err := c.ShouldBindQuery(&req); err != nil {
		r.Fail(err)
		return
	}

	list, pagination := logic.NewUserLogic().Message(global.DB.WithContext(c.Request.Context()), uid, req.FriendUserID, req.Page, req.PageSize)

	r.OkWithPage(http.StatusOK, list, pagination)
}

func JoinFunc(c *gin.Context) {

	r := pkg.NewResponse(c)

	list, pagination := logic.NewUserLogic().UserDepList(global.DB.WithContext(c.Request.Context()), 1)

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
