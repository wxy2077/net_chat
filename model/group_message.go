package model

import (
	"gorm.io/gorm"
	"net-chat/pkg"
)

type GroupMessage struct {
	ID int64 `gorm:"primaryKey" json:"id,omitempty"`

	SenderUserID int64 `json:"sender_user_id"`

	GroupID int64 `json:"group_id"`

	Content string `gorm:"column:content;type:text;" json:"content"`
	// 消息类型
	ContentType int64 `gorm:"column:content_type;type:tinyint(4);" json:"content_type"`

	CreatedAt *pkg.LocalTimeX `json:"created_at,omitempty"`
	UpdatedAt *pkg.LocalTimeX `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt  `gorm:"index" json:"deleted_at,omitempty"`
}

func (g *GroupMessage) TableName() string {
	return "group_messages"
}

func (g *GroupMessage) BatchCreate(db *gorm.DB, list []*GroupMessage) error {
	return db.Model(g).Omit("id").CreateInBatches(&list, 1000).Error
}
