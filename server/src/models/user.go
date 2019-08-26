package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

const (
	DELETE_FLAG_FALSE = iota
	DELETE_FLAG_TURE
)

// 用户表
type User struct {
	GormModel
	UserName   string `gorm:"not null;unique" json:"username"`
	PasswdSha1 string `gorm:"not null" json:"password"`
	Mobile     string `json:"mobile"  json:"mobile"`
	DeleteFlag uint8  `json:"deleteFlag"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

var UserRepository = newUserRepository()

func newUserRepository() *userRepository {
	return &userRepository{}
}

type userRepository struct{}

// 通过opt获取用户
func (this *userRepository) GetUserByOpt(db *gorm.DB, opt map[string]interface{}) *User {
	ret := &User{}
	if err := db.First(ret, opt).Error; err != nil {
		return nil
	}
	return ret
}

// 添加用户
func (this *userRepository) Post(db *gorm.DB, user *User) int64 {

	if err := db.Create(&user).Error; err != nil {
		return 0
	}
	return user.Id
}

// 获取用户信息
func (this *userRepository) Get(db *gorm.DB, id int64) *User {
	ret := &User{}
	if err := db.First(ret, "id = ? and delete_flag = ? ", id, DELETE_FLAG_FALSE).Error; err != nil {
		return nil
	}
	return ret
}

// 全局更新用户信息
func (this *userRepository) Update(db *gorm.DB, user *User) bool {
	if err := db.Save(user).Error; err != nil {
		return false
	}
	return true
}

// 删除用户
func (this *userRepository) Delete(db *gorm.DB, id int64) bool {
	user := &User{}
	user.Id = id

	if err := db.Model(&user).Update("delete_flag", DELETE_FLAG_TURE).Error; err != nil {
		return false
	}
	return true
}

// 获取所有正常用户
func (this *userRepository) GetUsers(db *gorm.DB) *[]User {
	users := &[]User{}
	if err := db.Find(users, "delete_flag = ?", DELETE_FLAG_FALSE).Error; err != nil {
		return nil
	}
	return users
}

// 获取正常用户,分页
func (this *userRepository) GetUsersPaged(db *gorm.DB, offset, perPage int64) *[]User {
	users := &[]User{}
	if err := db.Offset(offset).Limit(perPage).Find(users, "delete_flag = ?", DELETE_FLAG_FALSE).Error; err != nil {
		return nil
	}
	return users
}
