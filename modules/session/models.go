package session

type Session struct {
	UserID            int64
	UserName          string
	UserEmail         string
	UserRole          string
	UserEmailVerified bool
}
