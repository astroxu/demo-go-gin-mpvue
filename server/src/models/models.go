package models

//引入所有models
var Models = []interface{}{
	&User{},
}

type GormModel struct {
	Id int64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
}
