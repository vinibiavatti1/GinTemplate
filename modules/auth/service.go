package auth

import (
	"app/modules/session"
	"app/modules/user"
	"app/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context, email string, password string) (bool, error) {
	passwordHash := utils.Hash(password)
	users, err := user.Select(&user.UserFilter{
		Email:        &email,
		PasswordHash: &passwordHash,
	})
	if err != nil {
		return false, fmt.Errorf("auth.Login: %v", err)
	}
	if len(users) == 0 {
		return false, nil
	}
	user := users[0]
	if !user.Active {
		return false, nil
	}
	session.StartSession(c, &session.Session{
		UserID:            user.ID,
		UserName:          user.Name,
		UserEmail:         user.Email,
		UserRole:          user.Role,
		UserEmailVerified: user.EmailVerified,
	})
	return true, nil
}

func Logout(c *gin.Context) {
	session.TerminateSession(c)
}
