package interfaces

import (
	"context"
	"github.com/Lenstack/clean-architecture/internal/domain"
	"github.com/Lenstack/clean-architecture/internal/usecases"
	"github.com/Lenstack/clean-architecture/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type UserRepository struct {
	Mongo usecases.Mongo
}

func (ur *UserRepository) FindAll() (users domain.Users, err error) {
	ctx := context.Background()
	filter := bson.D{}

	cur, err := ur.Mongo.Query(domain.Mongo{CollectionName: domain.UserCollectionName, Context: ctx, Filter: filter})
	if err != nil {
		return
	}

	if err = cur.All(ctx, &users); err != nil {
		return
	}

	return users, nil
}

func (ur *UserRepository) FindById(userId string) (user domain.User, err error) {
	var ctx = context.TODO()
	objectID, _ := primitive.ObjectIDFromHex(userId)
	var filter = bson.M{"_id": objectID}

	if err = ur.Mongo.FindBy(domain.Mongo{Context: ctx, CollectionName: domain.UserCollectionName, Filter: filter}).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return
		}
	}
	return user, nil
}

func (ur *UserRepository) Create(userData domain.User) (result interface{}, err error) {
	var ctx = context.TODO()
	token, _ := utils.GenerateToken(userData.Account.Email)
	userData.Account.Password = utils.HashPassword(userData.Account.Password)
	userData.Account.Role = domain.USER
	userData.Account.Verified = false
	userData.Account.Token = token
	userData.CreatedAt = time.Now()
	userData.UpdatedAt = time.Now()

	_, err = ur.Mongo.Insert(domain.Mongo{Context: ctx, CollectionName: domain.UserCollectionName, Interface: userData})
	if err != nil {
		return
	}

	return token, nil
}

func (ur *UserRepository) Update(userId string, userData domain.User) (result interface{}, err error) {
	var ctx = context.TODO()
	objectID, _ := primitive.ObjectIDFromHex(userId)
	var filter = bson.D{{"_id", objectID}}
	var update = bson.D{{"$set", userData}}

	res, err := ur.Mongo.Update(domain.Mongo{Context: ctx, CollectionName: domain.UserCollectionName, Filter: filter, Interface: update})
	if err != nil {
		return
	}
	result = res.ModifiedCount
	return result, nil
}

func (ur *UserRepository) Delete(userId string) (result interface{}, err error) {
	var ctx = context.TODO()
	objectID, _ := primitive.ObjectIDFromHex(userId)
	var filter = bson.D{{"_id", objectID}}

	res, err := ur.Mongo.Delete(domain.Mongo{Context: ctx, CollectionName: domain.UserCollectionName, Filter: filter})
	if err != nil {
		return
	}
	result = res.DeletedCount

	return result, nil
}
