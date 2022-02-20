package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const UserCollectionName = "users"

type Users []User

type User struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"Name"`
	Account   Account            `bson:"Account"`
	CreatedAt time.Time          `bson:"Created_At"`
	UpdatedAt time.Time          `bson:"Updated_At,omitempty"`
}

type Account struct {
	Email    string `bson:"Email"`
	Password string `bson:"Password,omitempty"`
	Phone    uint   `bson:"Phone"`
	Token    string `bson:"Token"`
	Role     Role   `bson:"Role"`
	Verified bool   `bson:"Verified"`
}

type Role string

const (
	ADMIN Role = "admin"
	USER  Role = "user"
)
