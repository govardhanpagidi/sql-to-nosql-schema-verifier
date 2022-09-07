package clients

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"verifier/src/models"
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

func GetCollectionCount(connString, dbName, collName, embeddedPath, embeddedType, relKey string) (int64, error) {
	var ctx = context.Background()
	count := int64(0)
	var err error
	coll := GetCollection(connString, dbName, collName, ctx)
	switch embeddedType {
	case models.NEW_DOCUMENT:
		count, err = coll.CountDocuments(ctx, bson.D{})
	case models.EMBEDDED_DOCUMENT:
		count, err = getEmbeddedDocumentCount(coll, embeddedPath, relKey)
	case models.EMBEDDED_DOCUMENT_ARRAY:
		count, err = getEmbeddedArrayCount(coll, embeddedPath)
	}

	if err != nil {
		fmt.Println(err)
	}
	return count, err
}

func getEmbeddedArrayCount(collection *mongo.Collection, embeddedPath string) (int64, error) {
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
		return count, err
	}

	var results []bson.Raw
	if err = cursor.All(context.TODO(), &results); err != nil {
		return count, err
	}

	for _, result := range results {
		count = int64(result.Lookup("total").Int32())
	}
	return count, err
}

func getEmbeddedDocumentCount(collection *mongo.Collection, embeddedPath, relKey string) (int64, error) {
	var ctx = context.Background()
	count := int64(0)
	var aggregateQry = bson.A{
		bson.D{
			{"$project",
				bson.D{
					{"elems", "$" + embeddedPath},
					{"fid", "1"},
				},
			},
		},
		bson.D{{"$unwind", bson.D{{"path", "$elems"}}}},
		bson.D{
			{"$group",
				bson.D{
					{"_id", "$fid"},
					{"elems", bson.D{{"$addToSet", "$elems." + relKey}}},
				},
			},
		},
		bson.D{
			{"$project",
				bson.D{
					{"item", 1},
					{"total", bson.D{{"$size", "$elems"}}},
				},
			},
		},
	}
	cursor, err := collection.Aggregate(ctx, aggregateQry)
	if err != nil && cursor == nil {
		fmt.Println(err)
		return count, err
	}

	var results []bson.Raw
	if err = cursor.All(context.TODO(), &results); err != nil {
		return count, err
	}
	for _, result := range results {
		count = int64(result.Lookup("total").Int32())
	}
	return count, err
}

type DocCount struct {
	Id    int64
	Total int64
}
