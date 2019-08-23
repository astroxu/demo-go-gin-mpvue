package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

// 用户表
type User struct {
	GormModel
	UserName   string `gorm:"not null;unique" json:"userName"`
	PasswdSha1 string `gorm:"not null" json:"PasswdSha1"`
	Mobile     string `json:"mobile"`
	DeleteFlag uint8  `json:"deleteFlag"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

var UserRepository = newUserRepository()

func newUserRepository() *userRepository {
	return &userRepository{}
}

type userRepository struct{}

// 获取用户信息
func (this *userRepository) Get(db *gorm.DB, id int64) *User {
	ret := &User{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

/*func (this *userRepository) Add(db *gorm.DB,user User) uint {

}
*/
