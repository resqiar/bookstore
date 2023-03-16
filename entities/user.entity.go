package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"-"`
	IsAdmin  bool   `gorm:"default:false" json:"isAdmin"`
}
