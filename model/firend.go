package model

import (
	"gorm.io/gorm"
	"net-chat/pkg"
)

const (
	FriendStatusPending  = 0 // 申请中
	FriendStatusAccepted = 1 // 同意
	FriendStatusBlocked  = 2 // 拒绝
)

type Friend struct {
	ID int64 `gorm:"primaryKey" json:"id,omitempty"`

	UserID int64 `json:"user_id"`

	FriendUserID int64 `json:"friend_user_id"`

	// 申请状态 1:申请中 2:同意 3:拒绝
	Status int64 `gorm:"type:tinyint(4);default:0;comment:申请状态 1:申请中 2:同意 3:拒绝;" json:"status"`

	CreatedAt *pkg.LocalTimeX `json:"created_at,omitempty"`
	UpdatedAt *pkg.LocalTimeX `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt  `gorm:"index" json:"deleted_at,omitempty"`
}

func (f *Friend) TableName() string {
	return "friends"
}
