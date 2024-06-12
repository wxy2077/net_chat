package model

import (
	"gorm.io/gorm"
	"net-chat/pkg"
)

type GroupMember struct {
	ID int64 `gorm:"primaryKey" json:"id,omitempty"`

	UserID int64 `json:"user_id"`

	// 群管理员为1，普通成员为0
	Role int64 `gorm:"type:tinyint(4);default:0;" json:"role"`

	CreatedAt *pkg.LocalTimeX `json:"created_at,omitempty"`
	UpdatedAt *pkg.LocalTimeX `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt  `gorm:"index" json:"deleted_at,omitempty"`
}

func (g *GroupMember) TableName() string {
	return "group_members"
}
