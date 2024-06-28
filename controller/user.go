package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net-chat/global"
	"net-chat/logic"
	"net-chat/pkg"
	"net/http"
	"sort"
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
	token, err := logic.NewUserLogic(c).Login(global.DB.WithContext(c.Request.Context()), req.Username, req.Password)
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
		req = struct {
			FriendUserID int64 `form:"friend_user_id"`
		}{}
	)
	if err := c.ShouldBindQuery(&req); err != nil {
		r.Fail(err)
		return
	}
	// 优先查询好友信息
	if req.FriendUserID != 0 {
		uid = req.FriendUserID
	}

	user, err := logic.NewUserLogic(c).UserInfo(global.DB.WithContext(c.Request.Context()), uid)
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

	list, pagination := logic.NewUserLogic(c).FriendList(global.DB.WithContext(c.Request.Context()), uid, req.Page, req.PageSize)

	r.OkWithPage(http.StatusOK, list, pagination)
}

func Message(c *gin.Context) {
	var (
		r   = pkg.NewResponse(c)
		uid = c.GetInt64(global.UserID)
		req = struct {
			FriendUserID int64  `form:"friend_user_id"`
			MsgID        string `form:"msg_id"`
			Page         int64  `form:"page"`
			PageSize     int64  `form:"page_size"`
		}{
			Page:     pkg.DefaultPage,
			PageSize: 20,
		}
	)
	if err := c.ShouldBindQuery(&req); err != nil {
		r.Fail(err)
		return
	}

	list, pagination := logic.NewUserLogic(c).Message(global.DB.WithContext(c.Request.Context()), uid, req.FriendUserID, req.MsgID, req.Page, req.PageSize)

	sort.Slice(list, func(i, j int) bool {
		return list[i].ID < list[j].ID
	})

	r.OkWithPage(http.StatusOK, list, pagination)
}

func UnreadMessage(c *gin.Context) {
	var (
		r   = pkg.NewResponse(c)
		uid = c.GetInt64(global.UserID)
	)

	list := logic.NewUserLogic(c).UnreadMessage(global.DB.WithContext(c.Request.Context()), uid)

	r.OK(http.StatusOK, list)
}
