package database

import (
	"context"
	"fmt"
	"github.com/Lenstack/clean-architecture/internal/usecases"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	"time"
)

func MongoConnection() (*mongo.Client, context.Context, error) {
	datasourceName := fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_DATABASE"))
	ctx, _ := context.WithTimeout(context.Background(),
		30*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(datasourceName))
	return client, ctx, err
}

func MongoPing(logger usecases.Logger, client *mongo.Client, ctx context.Context) {
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		logger.LogError("%s", err)
	}
	logger.LogAccess("The %v Database Has Been Successfully Connected.", os.Getenv("DB_DATABASE"))
}

func MongoClose(logger usecases.Logger, client *mongo.Client, ctx context.Context) {
	if err := client.Disconnect(ctx); err != nil {
		logger.LogError("%s", err)
	}
}

func MongoListOfDatabases(logger usecases.Logger, client *mongo.Client, ctx context.Context) {
	list, err := client.ListDatabases(ctx, bson.M{})
	if err != nil {
		logger.LogError("%s", err)
	}
	logger.LogAccess("Count Of Databases Available: %v", len(list.Databases))
}
