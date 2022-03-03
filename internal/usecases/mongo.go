package usecases

import (
	"github.com/Lenstack/clean-architecture/internal/domain/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository interface {
	Query(entity.Mongo) (*mongo.Cursor, error)
	FindBy(entity.Mongo) *mongo.SingleResult
	Insert(entity.Mongo) (*mongo.InsertOneResult, error)
	Update(entity.Mongo) (*mongo.UpdateResult, error)
	Delete(entity.Mongo) (*mongo.DeleteResult, error)
}
