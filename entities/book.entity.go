package entities

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string `gorm:"type:varchar(100);not null;" json:"title"`
	Description string `gorm:"type:text;not null;" json:"description"`
	ImageURL    string `gorm:"type:text;" json:"imageURL"`
	ReleaseDate string `gorm:"type:varchar(45); not null;" json:"releaseDate"`
	Author      string `gorm:"type:varchar(45); not null;" json:"author"`
	Price       int    `gorm:"type:int; not null;" json:"price"`
}
