package user

type UserFilter struct {
	ID            *int64
	Name          *string
	Email         *string
	PasswordHash  *string
	Role          *string
	Active        *bool
	Hash          *string
	EmailVerified *bool
}
