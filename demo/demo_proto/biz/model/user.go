package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"uniqueIndex;type:varchar(128) not null"`
	Password string `gorm:"type:varchar(64) not null"`
}

// TableName : 指定生成的表名字（可选，默认是snake_case+s）
func (u *User) TableName() string {
	return "user"
}
