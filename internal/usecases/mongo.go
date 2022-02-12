package usecases

import (
	"github.com/Lenstack/clean-architecture/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo interface {
	Query(domain.Mongo) (*mongo.Cursor, error)
	FindBy(domain.Mongo) *mongo.SingleResult
	Insert(domain.Mongo) (*mongo.InsertOneResult, error)
	Update(domain.Mongo) (*mongo.UpdateResult, error)
	Delete(domain.Mongo) (*mongo.DeleteResult, error)
}
