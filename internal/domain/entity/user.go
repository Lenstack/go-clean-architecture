package entity

import (
	"time"
)

type User struct {
	Id        string
	Name      string
	Account   Account
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Account struct {
	Email    string
	Password string
	Phone    uint
	Token    string
	Role     Role
	Verified bool
}

type Role string
