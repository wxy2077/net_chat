package model

import (
	"github.com/guregu/null"
	"gorm.io/gorm"
	"net-chat/pkg"
)

const (
	CommonYes            = 1
	CommonNo             = 0
	CommonUserQueryField = "id,account,username,phone,avatar,email"
)

type User struct {
	ID       int64       `gorm:"primaryKey" json:"id,omitempty"`
	Account  string      `gorm:"column:account;type:varchar(191);default:'';comment:账号;" json:"account,omitempty"`
	Password string      `gorm:"column:password;type:varchar(191);comment:密码;" json:"-"`
	Username null.String `gorm:"column:username;type:varchar(191);comment:昵称;" json:"username,omitempty"`
	Phone    string      `gorm:"column:phone;type:varchar(16);comment:手机号;" json:"phone,omitempty"`
	Avatar   null.String `gorm:"column:avatar;type:varchar(191);comment:头像;" json:"avatar,omitempty"`
	Email    null.String `gorm:"column:email;type:varchar(191);comment:头像;" json:"email,omitempty"`

	CreatedAt *pkg.LocalTimeX `json:"created_at,omitempty"` // 自定义时间JSON序列化
	UpdatedAt *pkg.LocalTimeX `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt  `gorm:"index" json:"deleted_at,omitempty"`

	// 逻辑外键关系  references:为自身model的字段   foreignKey:为对应关联的model字段
	DepartmentUser []*DepartmentUser `gorm:"foreignKey:UserID;references:ID" json:"department_user,omitempty"`

	// 逻辑外间 多对多 同上
	Department []*Department `gorm:"many2many:department_users;foreignKey:ID;joinForeignKey:UserID;References:ID;JoinReferences:DepID;" json:"departments,omitempty"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) AfterFind(tx *gorm.DB) (err error) {

	// 没有设置头像时，设置默认头像
	if u.Avatar.IsZero() {
		u.Avatar.SetValid("http://default.avatar")
	}

	return nil
}

// FilterUser 筛选条件
type FilterUser struct {
	ID       int64
	IDs      []int64
	Account  string
	Username null.String
	Phone    string
	Page     int64
	PageSize int64
}

// First 获取一条数据
// 尽可能封装更多的条件场景
func (u *User) First(db *gorm.DB, filter *FilterUser, columns ...string) error {
	db = db.Model(&User{})
	if len(columns) > 0 {
		db = db.Select(columns)
	}
	if filter.ID > 0 {
		db = db.Where("id = ?", filter.ID)
	}
	if filter.Account != "" {
		db = db.Where("account = ?", filter.Account)
	}
	if filter.Username.Valid {
		db = db.Where("username = ?", filter.Username.String)
	}
	if filter.Phone != "" {
		db = db.Where("phone = ?", filter.Phone)
	}
	return db.First(u).Error
}

// Find 分页
// model 层只写不包含任何业务逻辑的纯SQL操作
// 然后在其他层组合基本操作逻辑成业务
func (u *User) Find(db *gorm.DB, filter *FilterUser, columns ...string) (list []*User, pagination *pkg.Pagination) {
	db = db.Model(&User{})

	if len(columns) > 0 {
		db = db.Select(columns)
	}
	if len(filter.IDs) > 0 {
		db = db.Where("id in (?)", filter.IDs)
	}
	if filter.Account != "" {
		db = db.Where("account", filter.Account)
	}
	if filter.Username.Valid {
		db = db.Where("username", filter.Username.String)
	}
	if filter.Phone != "" {
		db = db.Where("phone", filter.Phone)
	}
	pagination = pkg.Paginate(&pkg.Param{
		DB:      db,
		Page:    filter.Page,
		Limit:   filter.PageSize,
		OrderBy: []string{"created_at desc"},
	}, &list)
	return list, pagination
}
