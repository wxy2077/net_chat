package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net-chat/global"
	model "net-chat/model"
	"net-chat/pkg"
	"time"
)

type UserLogic interface {

	// Login 登录
	Login(db *gorm.DB, account, password string) (token string, err error)

	UserInfo(db *gorm.DB, id int64) (user *model.User, err error)

	FriendList(db *gorm.DB, uid, page, pageSize int64) (list []*model.User, pagination *pkg.Pagination)

	Message(db *gorm.DB, uid, friendUid int64, msgId string, page, pageSize int64) (list []*model.Message, pagination *pkg.Pagination)

	UnreadMessage(db *gorm.DB, uid int64) (list []*model.Message)
}

type userLogic struct {
	ctx *gin.Context
}

// NewUserLogic 接口controller层直接调用
// 完成各种业务操作
func NewUserLogic(c *gin.Context) UserLogic {
	return &userLogic{ctx: c}
}

func (u *userLogic) Login(db *gorm.DB, account, password string) (token string, err error) {

	user := new(model.User)
	if err = user.First(db, &model.FilterUser{
		Account: account,
	}, "id,account,password"); err != nil {
		return "", errors.New("账号或者密码错误")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("账号或者密码错误")
	}

	token, _ = pkg.PwdJwt.GenerateJwtToken(jwt.MapClaims{
		"exp":         time.Now().Add(time.Hour * 24 * 365).Unix(),
		global.UserID: user.ID,
		"scope":       "[*]",
	})

	return token, nil
}

func (u *userLogic) UserInfo(db *gorm.DB, id int64) (user *model.User, err error) {
	user = new(model.User)
	if err = user.First(db, &model.FilterUser{
		ID: id,
	}); err != nil {
		return nil, errors.New(err.Error())
	}
	return user, nil
}

func (u *userLogic) FriendList(db *gorm.DB, uid, page, pageSize int64) (list []*model.User, pagination *pkg.Pagination) {
	if uid == 0 {
		return list, pagination
	}

	db = db.Model(&model.User{}).Joins("LEFT JOIN friends f ON (f.user_id = ? AND f.friend_user_id = users.id ) OR (f.user_id = users.id AND f.friend_user_id = ?)", uid, uid).
		Where("f.status = ?", model.FriendStatusAccepted)

	pagination = pkg.Paginate(&pkg.Param{
		DB:      db,
		Page:    page,
		Limit:   pageSize,
		OrderBy: []string{"id ASC"},
	}, &list)

	return list, pagination
}

func (u *userLogic) Message(db *gorm.DB, uid, friendUid int64, msgId string, page, pageSize int64) (list []*model.Message, pagination *pkg.Pagination) {
	if uid == 0 || friendUid == 0 || msgId == "" {
		return list, pagination
	}
	msg := new(model.Message)
	err := msg.Single(db, &model.MessageFilter{SenderUserID: uid, ReceiverUserID: friendUid, MsgID: msgId, Columns: "id"})
	if err != nil {
		return list, pagination
	}
	if msg.ID == 0 {
		return list, pagination
	}

	db = db.Model(&model.Message{}).
		Where("id < ?", msg.ID).
		Where("(sender_user_id = ? AND receiver_user_id = ?) OR (sender_user_id = ? AND receiver_user_id = ?)", uid, friendUid, friendUid, uid)

	pagination = pkg.Paginate(&pkg.Param{
		DB:      db,
		Page:    page,
		Limit:   pageSize,
		OrderBy: []string{"created_at DESC"},
	}, &list)

	return list, pagination
}

func (u *userLogic) UnreadMessage(db *gorm.DB, uid int64) (list []*model.Message) {
	if uid == 0 {
		return list
	}

	msg := new(model.Message)
	unreadList := msg.UnReadMsg(db, &model.MessageFilter{ReceiverUserID: uid})

	unreadIdList := make([]int64, 0)
	unreadCountMap := make(map[int64]int64)
	for _, v := range unreadList {
		unreadIdList = append(unreadIdList, v.LastMessageID)
		unreadCountMap[v.LastMessageID] = v.UnreadCount
	}

	list = msg.FindByIDs(db, &model.MessageFilter{IDs: unreadIdList})

	for _, v := range list {
		if _, ok := unreadCountMap[v.ID]; ok {
			v.UnreadCount = unreadCountMap[v.ID]
		}
	}

	// 更新为已读
	go func() {
		if err := msg.UpdateRead(global.DB, &model.MessageFilter{IDs: unreadIdList}); err != nil {
			global.Log.
				WithFields(logrus.Fields{"ids": unreadIdList}).
				Errorf("update error:%+v", err)
		}
	}()

	return list
}
