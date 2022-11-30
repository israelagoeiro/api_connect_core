package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

var DbMong *mongo.Client

type DataLog struct {
	Action       string
	SaveChange   bool
	SaveHistory  bool
	SaveInfo     bool
	SaveAnalytic bool
	Info         InfoModel
}

func MongoDBOpen() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_DB_CONNECT")))

	if err != nil {
		fmt.Println("MongoDBOpen--->", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 90*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}

/*var MongoClient *mongo.Client

func MongoDBClose() {
	if MongoClient == nil {
		return
	}

	err := MongoClient.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// TODO optional you can log your closed MongoDB client
	fmt.Println("Connection to MongoDB closed.")
}*/

func GetCollection(mongoClient *mongo.Client, database string, collectionName string) *mongo.Collection {
	return mongoClient.Database(database).Collection(collectionName)
}
