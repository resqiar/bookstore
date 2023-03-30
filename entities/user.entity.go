package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(100);unique;not null;" json:"email"`
	Username string `gorm:"type:varchar(100);unique;not null;" json:"username"`
	Password string `json:"-"`
	IsAdmin  bool   `gorm:"default:false" json:"isAdmin"`
}
