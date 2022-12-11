package request

type Message struct {
	Receiver uint   `json:"receiver_id"`
	Sender   uint   `json:"sender_id"`
	Body     string `json:"body"`
	IsSeen   bool   `json:"isSeen"`
}
