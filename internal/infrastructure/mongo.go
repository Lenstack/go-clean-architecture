package infrastructure

import (
	"github.com/Lenstack/clean-architecture/internal/domain/entity"
	"github.com/Lenstack/clean-architecture/internal/infrastructure/database"
	"github.com/Lenstack/clean-architecture/internal/usecases"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

type Mongo struct {
	Database *mongo.Database
}

func NewMongo(logger usecases.LoggerRepository) (mongo usecases.MongoRepository) {
	client, ctx, err := database.MongoConnection()
	if err != nil {
		defer database.MongoClose(logger, client, ctx)
	}

	defer database.MongoPing(logger, client, ctx)
	defer database.MongoListOfDatabases(logger, client, ctx)

	return &Mongo{
		Database: client.Database(os.Getenv("DB_DATABASE")),
	}
}

func (m Mongo) Query(mongo entity.Mongo) (*mongo.Cursor, error) {
	return m.Database.Collection(mongo.CollectionName).Find(mongo.Context, mongo.Filter)
}

func (m Mongo) FindBy(mongo entity.Mongo) *mongo.SingleResult {
	return m.Database.Collection(mongo.CollectionName).FindOne(mongo.Context, mongo.Filter)
}

func (m Mongo) Insert(mongo entity.Mongo) (*mongo.InsertOneResult, error) {
	return m.Database.Collection(mongo.CollectionName).InsertOne(mongo.Context, mongo.Interface)
}

func (m Mongo) Update(mongo entity.Mongo) (*mongo.UpdateResult, error) {
	return m.Database.Collection(mongo.CollectionName).UpdateOne(mongo.Context, mongo.Filter, mongo.Interface)
}

func (m Mongo) Delete(mongo entity.Mongo) (*mongo.DeleteResult, error) {
	return m.Database.Collection(mongo.CollectionName).DeleteOne(mongo.Context, mongo.Filter)
}
