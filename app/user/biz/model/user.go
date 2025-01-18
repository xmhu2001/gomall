package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email          string `gorm:"uniqueIndex;type:varchar(128) not null"`
	PasswordHashed string `gorm:"type:varchar(255) not null"`
}

func (u *User) TableName() string {
	return "user"
}
func Create(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}
