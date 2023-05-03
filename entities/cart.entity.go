package entities

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID   int  `gorm:"type:int;not null;" json:"userID"`
	BookID   int  `gorm:"type:int;not null;" json:"bookID"`
	Quantity int  `gorm:"type:int;not null;" json:"quantity"`
	Book     Book `gorm:"foreignkey:BookID" json:"book"`
}
