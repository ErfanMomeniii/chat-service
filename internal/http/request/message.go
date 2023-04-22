package request

type Message struct {
	Receiver uint   `json:"receiver_id" binding:"required"`
	Sender   uint   `json:"sender_id" binding:"required"`
	Body     string `json:"body"`
}
