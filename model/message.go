package model

import (
	"errors"
	"github.com/guregu/null"
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

	// 前端生成的uuid
	MsgID string `gorm:"column:msg_id;type:varchar(36);" json:"msg_id"`

	User *User `gorm:"foreignKey:SenderUserID;references:ID" json:"sender_user,omitempty"`

	CreatedAt *pkg.LocalTimeX `json:"created_at,omitempty"`
	UpdatedAt *pkg.LocalTimeX `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt  `gorm:"index" json:"deleted_at,omitempty"`

	UnreadCount int64 `gorm:"-" json:"unread_count,omitempty"`
}

func (m *Message) TableName() string {
	return "messages"
}

func (m *Message) Create(db *gorm.DB) error {
	return db.Omit("id").Create(m).Error
}

func (m *Message) BatchCreate(db *gorm.DB, list []*Message) error {
	return db.Model(m).Omit("id").CreateInBatches(&list, 1000).Error
}

type MessageFilter struct {
	IDs            []int64
	SenderUserID   int64
	ReceiverUserID int64
	MsgID          string
	IsRead         null.Int
	Columns        string
	LoadUser       bool
	OnlyCount      bool
	Page           int64
	PageSize       int64
}

func (m *Message) UpdateRead(db *gorm.DB, filter *MessageFilter) error {
	if len(filter.IDs) <= 0 {
		return nil
	}
	return db.Model(&Message{}).Where("id in (?)", filter.IDs).
		Update("is_read", CommonYes).Error
}

func (m *Message) FindByIDs(db *gorm.DB, filter *MessageFilter) (list []*Message) {
	db = db.Model(&Message{})
	if len(filter.IDs) == 0 {
		return nil
	}
	db.Where("id in (?)", filter.IDs).
		Preload("User", func(tx *gorm.DB) *gorm.DB {
			return tx.Select(CommonUserQueryField)
		}).
		Find(&list)

	return
}

type UnReadMsg struct {
	SenderUserID  int64
	UnreadCount   int64
	LastMessageID int64
}

func (m *Message) UnReadMsg(db *gorm.DB, filter *MessageFilter) (list []*UnReadMsg) {
	list = make([]*UnReadMsg, 0)

	db = db.Model(&Message{}).Where("is_read = ?", CommonNo)

	if filter.SenderUserID == 0 && filter.ReceiverUserID == 0 {
		return nil
	}
	if filter.ReceiverUserID != 0 {
		db = db.Where("receiver_user_id = ?", filter.ReceiverUserID)
	}

	db.Select("sender_user_id, COUNT(*) AS unread_count,MAX(id) AS last_message_id").
		Group("sender_user_id").Scan(&list)

	return
}

func (m *Message) Single(db *gorm.DB, filter *MessageFilter) error {

	db = db.Model(&Message{})
	if filter.SenderUserID == 0 && filter.ReceiverUserID == 0 {
		// 防止空条件查询
		return errors.New("filter is empty")
	}
	if filter.MsgID != "" {
		db = db.Where("msg_id = ?", filter.MsgID)
	}

	if filter.SenderUserID != 0 {
		db = db.Where("sender_user_id = ?", filter.SenderUserID)
	}
	if filter.ReceiverUserID != 0 {
		db = db.Where("receiver_user_id = ?", filter.ReceiverUserID)
	}
	if filter.Columns != "" {
		db = db.Select(filter.Columns)
	}
	return db.First(&m).Error
}

func (m *Message) List(db *gorm.DB, filter *MessageFilter) (list []*Message, count int64, pagination *pkg.Pagination) {
	db = db.Model(&Message{})
	if filter.SenderUserID == 0 && filter.ReceiverUserID == 0 {
		// 防止空条件查询
		return
	}
	if filter.SenderUserID != 0 {
		db = db.Where("sender_user_id = ?", filter.SenderUserID)
	}
	if filter.ReceiverUserID != 0 {
		db = db.Where("receiver_user_id = ?", filter.ReceiverUserID)
	}

	if filter.IsRead.Valid {
		db = db.Where("is_read = ?", filter.IsRead.Int64)
	}

	if filter.OnlyCount {
		db = db.Count(&count)
		return
	}

	if filter.LoadUser {
		db = db.Preload("User")
	}

	pagination = pkg.Paginate(&pkg.Param{
		DB:      db,
		Page:    filter.Page,
		Limit:   filter.PageSize,
		OrderBy: []string{"created_at DESC"},
	}, &list)
	return
}
