package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username        string
	UserInformation UserInformation
}

type UserInformation struct {
	Firstname string
	Lastname  string
	Tel       string
}
