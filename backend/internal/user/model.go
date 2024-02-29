package user

import "time"

type Role string

const (
	Administrator   Role = "Administrator"
	HeadOfDepatment Role = "HeadOfDepatment"
	Worker          Role = "Worker"
)

type User struct {
	ID          string `json:"id"`
	Login       string `json:"login"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Email       string `json:"email"`
	Password    []byte
	Role        Role
	Created     time.Time
	IsConfirmed bool
}
