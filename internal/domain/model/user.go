package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const UserCollectionName = "users"

type Users []User

type User struct {
	Id        primitive.ObjectID `json:"Id" bson:"_id,omitempty"`
	Name      string             `json:"Name" validate:"required"`
	Account   Account            `json:"Account" validate:"required"`
	CreatedAt time.Time          `json:"CreatedAt"`
	UpdatedAt time.Time          `json:"UpdatedAt"`
}

type Account struct {
	Email    string `json:"Email" validate:"email"`
	Password string `json:"Password" validate:"required"`
	Phone    uint   `json:"Phone" validate:"required"`
	Token    string `json:"Token"`
	Role     Role   `json:"Role"`
	Verified bool   `json:"Verified"`
}

type Role string

var (
	ADMIN Role = "admin"
	USER  Role = "user"
)
