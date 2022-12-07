package model

type User struct {
	Username        string
	UserInformation UserInformation
}

type UserInformation struct {
	Firstname string
	Lastname  string
	Tel       string
}
