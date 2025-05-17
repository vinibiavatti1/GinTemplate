package message

type MessageCategory string

const (
	Success MessageCategory = "success"
	Info    MessageCategory = "info"
	Warning MessageCategory = "warning"
	Error   MessageCategory = "error"
)
