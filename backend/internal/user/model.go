package user

import "time"

type Role int

const (
	Administrator Role = iota
	HeadOfDepatment
	Worker
)

type User struct {
	ID          string `json:"id"`
	Login       string `json:"login"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Email       string `json:"email"`
	Role        Role
	Created     time.Time
	IsConfirmed bool
}
