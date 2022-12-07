package model

type Message struct {
	From   User
	To     User
	Body   string
	IsSeen bool
}
