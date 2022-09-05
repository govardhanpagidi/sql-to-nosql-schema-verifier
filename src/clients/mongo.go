package clients

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getMongoClient(connString string, ctx context.Context) (*mongo.Client, error) {
	var err error
	var client *mongo.Client

	opts := options.Client()
	opts.ApplyURI(connString)
	opts.SetMaxPoolSize(100)

	if client, err = mongo.Connect(ctx, opts); err != nil {
		return client, err
	}
	return client, err
}

func GetCollection(connString, dbName, collName string, ctx context.Context) *mongo.Collection {
	client, err := getMongoClient(connString, ctx)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var collection *mongo.Collection
	collection = client.Database(dbName).Collection(collName)
	return collection
}

func GetCollectionCount(connString, dbName, collName string) int64 {
	var ctx = context.Background()

	coll := GetCollection(connString, dbName, collName, ctx)
	count, err := coll.CountDocuments(ctx, bson.D{})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return count
}
