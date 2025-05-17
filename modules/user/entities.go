package user

import "time"

type User struct {
	ID            int64     `db:"ID"`
	Name          string    `db:"Name"`
	Email         string    `db:"Email"`
	PasswordHash  string    `db:"PasswordHash"`
	Role          string    `db:"Role"`
	CreatedAt     time.Time `db:"CreatedAt"`
	UpdatedAt     time.Time `db:"UpdatedAt"`
	Active        bool      `db:"Active"`
	LastLogin     time.Time `db:"LastLogin"`
	Hash          string    `db:"Hash"`
	EmailVerified bool      `db:"EmailVerified"`
}
