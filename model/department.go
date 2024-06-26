package model

import (
	"github.com/guregu/null"
	"gorm.io/gorm"
	"net-chat/pkg"
)

// Department 具体gorm tag见user表
type Department struct {
	ID int64 `json:"id,omitempty"`

	Title string `json:"title,omitempty"`

	ParentID null.Int `json:"parent_id,omitempty"`

	Level int64 `json:"level,omitempty"`

	Path string `json:"path,omitempty"`

	CreatedAt *pkg.LocalTimeX `json:"created_at,omitempty"` // 自定义时间JSON序列化
	UpdatedAt *pkg.LocalTimeX `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt  `json:"deleted_at,omitempty"`
}

func (d *Department) TableName() string {
	return "departments"
}
