package response

type User struct {
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Tel       string `json:"tel"`
	Bio       string `json:"bio"`
}
