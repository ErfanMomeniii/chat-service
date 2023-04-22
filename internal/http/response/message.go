package response

type Message struct {
	Receiver string `json:"receiver"`
	Sender   string `json:"sender"`
	Body     string `json:"body"`
}
