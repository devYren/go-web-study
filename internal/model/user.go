package model

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Username     string `gorm:"uniqueIndex;size:32;not null"`
	Email        string `gorm:"uniqueIndex;size:128;not null"`
	PasswordHash string `gorm:"size:255;not null"`
}

func (UserModel) TableName() string {
	return "users"
}
