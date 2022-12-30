package request

type User struct {
	Username  string `json:"username" binding:"required"`
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Tel       string `json:"tel"`
}
