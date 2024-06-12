package model

import (
	"gorm.io/gorm"
	"net-chat/pkg"
)

type Message struct {
	ID int64 `gorm:"primaryKey" json:"id,omitempty"`

	SenderUserID int64 `json:"sender_user_id"`

	ReceiverUserID int64 `json:"receiver_user_id"`

	Content string `gorm:"column:content;type:text;" json:"content"`

	// 消息内容类型：1.文字 2.普通文件 3.图片 4.音频 5.视频 6.语音聊天 7.视频聊天
	ContentType int64 `gorm:"column:content_type;type:tinyint(4);" json:"content_type"`

	IsRead int64 `gorm:"column:is_read;type:tinyint(4);" json:"is_read"`

	CreatedAt *pkg.LocalTimeX `json:"created_at,omitempty"`
	UpdatedAt *pkg.LocalTimeX `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt  `gorm:"index" json:"deleted_at,omitempty"`
}

func (m *Message) TableName() string {
	return "messages"
}
