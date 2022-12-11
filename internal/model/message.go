package model

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	FromRefer uint
	From      User `gorm:"foreignKey:FromRefer"`
	ToRefer   uint
	To        User `gorm:"foreignKey:FromRefer"`
	Body      string
	IsSeen    bool
}
