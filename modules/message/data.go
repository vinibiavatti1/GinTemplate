package message

var Messages = map[string]*Message{
	"loginSuccess":  NewMessage(Success, "Login Success"),
	"loginFailed":   NewMessage(Error, "Incorrect Username or Password"),
	"logoutSuccess": NewMessage(Success, "Logout Success"),
}
