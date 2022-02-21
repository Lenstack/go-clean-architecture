package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const UserCollectionName = "users"

type Users []User

type User struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"Id" `
	Name      string             `bson:"Name" json:"Name" validate:"required"`
	Account   Account            `bson:"Account" json:"Account" validate:"required"`
	CreatedAt time.Time          `bson:"Created_At" json:"CreatedAt"`
	UpdatedAt time.Time          `bson:"Updated_At,omitempty" json:"UpdatedAt"`
}

type Account struct {
	Email    string `bson:"Email" json:"Email" validate:"required"`
	Password string `bson:"Password,omitempty" json:"Password" validate:"required"`
	Phone    uint   `bson:"Phone" json:"Phone" validate:"required"`
	Token    string `bson:"Token" json:"Token"`
	Role     Role   `bson:"Role" json:"Role"`
	Verified bool   `bson:"Verified" json:"Verified"`
}

type Role string

const (
	ADMIN Role = "admin"
	USER  Role = "user"
)
