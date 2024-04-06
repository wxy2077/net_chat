package logic

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	model "net-chat/model"
	"net-chat/pkg"
)

type UserLogic interface {

	// Login 登录
	Login(db *gorm.DB, account, password string) (token string, err error)

	UserInfo(db *gorm.DB, id int64) (user *model.User, err error)

	// UserDepList 连接查询用户所在部门
	UserDepList(db *gorm.DB, page int64) (list []*UserList, pagination *pkg.Pagination)

	// PreloadUserDep 预加载查询出所有的部门
	PreloadUserDep(db *gorm.DB, page int64) (list []*model.User, pagination *pkg.Pagination)

	// 同上
	PreloadUserDeps(db *gorm.DB, page int64) (list []*model.User, pagination *pkg.Pagination)

	// 原生SQL查询

	// 其他业务逻辑方法....
}

type userLogic struct {
}

// NewUserLogic 接口controller层直接调用
// 完成各种业务操作
func NewUserLogic() UserLogic {
	return &userLogic{}
}

func (u *userLogic) Login(db *gorm.DB, account, password string) (token string, err error) {

	user := new(model.User)
	if err := user.First(db, &model.FilterUser{
		Account: account,
	}, "id,account,password"); err != nil {
		// 错误返回可以再次进行封装
		return "", errors.New(err.Error())
	}

	// TODO
	// 处理密码散列校验
	// 生成token

	return "", err
}

func (u *userLogic) UserInfo(db *gorm.DB, id int64) (user *model.User, err error) {
	user = new(model.User)
	if err = user.First(db, &model.FilterUser{
		ID: id,
	}, "id,account,password"); err != nil {
		return nil, errors.New(err.Error())
	}
	return user, nil
}

type UserList struct {
	*model.User
	Title string `json:"title"`
}

func (u *userLogic) UserDepList(db *gorm.DB, page int64) (list []*UserList, pagination *pkg.Pagination) {

	db = db.Model(&model.User{}).Select("d.title,users.*").
		Joins("LEFT JOIN department_users du ON du.user_id = users.id").
		Joins("LEFT JOIN departments d ON d.dep_id = du.dep_id")

	pagination = pkg.Paginate(&pkg.Param{
		DB:      db,
		Page:    page,
		Limit:   15,
		OrderBy: []string{"id desc"},
	}, &list)

	return list, pagination
}

func (u *userLogic) PreloadUserDep(db *gorm.DB, page int64) (list []*model.User, pagination *pkg.Pagination) {
	db = db.Model(&model.User{}).
		Preload("DepartmentUser", func(tx *gorm.DB) *gorm.DB {
			return tx.Preload("Department", func(tx *gorm.DB) *gorm.DB {
				return tx.Select("id,title")
			})
		})

	pagination = pkg.Paginate(&pkg.Param{
		DB:      db,
		Page:    page,
		Limit:   15,
		OrderBy: []string{"id desc"},
	}, &list)

	return list, pagination
}

func (u *userLogic) PreloadUserDeps(db *gorm.DB, page int64) (list []*model.User, pagination *pkg.Pagination) {

	list = make([]*model.User, 0, 1)

	db = db.Model(&model.User{}).Preload("Department")

	pagination = pkg.Paginate(&pkg.Param{
		DB:      db,
		Page:    page,
		Limit:   15,
		OrderBy: []string{"id desc"},
	}, &list)

	return list, pagination
}
