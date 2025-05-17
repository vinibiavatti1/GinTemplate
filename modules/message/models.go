package message

type Message struct {
	Category MessageCategory
	Message  string
}

func NewMessage(category MessageCategory, message string) *Message {
	return &Message{
		Category: category,
		Message:  message,
	}
}
