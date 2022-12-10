package model

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	FromRefer int
	From      User `gorm:"foreignKey:FromRefer"`
	ToRefer   int
	To        User `gorm:"foreignKey:FromRefer"`
	Body      string
	IsSeen    bool
}
