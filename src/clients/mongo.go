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

func GetCollectionCount(connString, dbName, collName string, embeddedPath string) int64 {
	var ctx = context.Background()
	count := int64(0)
	coll := GetCollection(connString, dbName, collName, ctx)
	if embeddedPath == "" {
		count, err := coll.CountDocuments(ctx, bson.D{})
		if err != nil {
			fmt.Println(err)
		}
		return count
	}
	count = getEmbeddedCollectionCount(coll, embeddedPath)
	return count
}

func getEmbeddedCollectionCount(collection *mongo.Collection, embeddedPath string) int64 {
	var ctx = context.Background()
	count := int64(0)
	var aggregateQry = bson.A{
		bson.D{
			{"$project",
				bson.D{
					{"item", 1},
					{"doc_count",
						bson.D{
							{"$cond",
								bson.D{
									{"if", bson.D{{"$isArray", "$" + embeddedPath}}},
									{"then", bson.D{{"$size", "$" + embeddedPath}}},
									{"else", 0},
								},
							},
						},
					},
				},
			},
		},
		bson.D{
			{"$group",
				bson.D{
					{"_id", 1},
					{"total", bson.D{{"$sum", "$doc_count"}}},
				},
			},
		},
	}
	cursor, err := collection.Aggregate(ctx, aggregateQry)
	if err != nil && cursor == nil {
		fmt.Println(err)
		panic(err)
	}

	var results []bson.Raw
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	for _, result := range results {
		count = int64(result.Lookup("total").Int32())
	}
	return count
}

type DocCount struct {
	Id    int64
	Total int64
}
