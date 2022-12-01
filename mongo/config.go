package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

var Client *mongo.Client

func NewClient() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_DB_CONNECT")))

	if err != nil {
		fmt.Println("Error NewClient", err)
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

	fmt.Println("NewClient: Connected to MongoDB")

	return client
}

func Disconnect() {
	if Client == nil {
		return
	}

	err := Client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// TODO optional you can log your closed MongoDB client
	fmt.Println("Connection to MongoDB closed.")
}

func GetCollection(database string, collectionName string) *mongo.Collection {
	return Client.Database(database).Collection(collectionName)
}
