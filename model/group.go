package model

import (
	"github.com/guregu/null"
	"gorm.io/gorm"
	"net-chat/pkg"
)

type Group struct {
	ID int64 `gorm:"primaryKey" json:"id,omitempty"`

	// 群名称
	Name string `json:"name"`
	// 群头像
	Avatar null.String `json:"avatar"`
	// 群主
	UserID int64 `json:"user_id"`

	CreatedAt *pkg.LocalTimeX `json:"created_at,omitempty"`
	UpdatedAt *pkg.LocalTimeX `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt  `gorm:"index" json:"deleted_at,omitempty"`
}

func (g *Group) TableName() string {
	return "groups"
}
