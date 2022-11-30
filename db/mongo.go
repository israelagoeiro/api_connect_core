package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

type DataLog struct {
	Action       string
	SaveChange   bool
	SaveHistory  bool
	SaveInfo     bool
	SaveAnalytic bool
	Info         InfoModel
}

type UpdateOptions struct {
	ReturnOriginal bool
	Upsert         bool
}

type MongoUpdateParams struct {
	Collection    string
	Connection    string
	Database      string
	DataLog       DataLog
	Fields        []string
	Filter        MongoFilter
	FindParams    MongoFindParams
	Info          any
	Input         MongoInputUpdate
	UpdateOptions UpdateOptions
}

func (param MongoUpdateParams) findOneAndUpdate() DataResult {
	document := NewMongoDocumentUpdate(param)
	return document.UpdateOne()
}

type MongoInsertParams struct {
	Collection string
	Connection string
	Database   string
	DataLog    DataLog
	Input      MongoInputInsert
	FindParams MongoFindParams
}

func (param MongoInsertParams) insertOne() DataResult {
	document := NewMongoDocumentInsert(param)
	return document.InsertOne()
}
func (param MongoInsertParams) insertMany() DataResult {
	document := NewMongoDocumentInsert(param)
	return document.InsertOne()
}

type FindOptions struct {
	Sort bson.D
}
type MongoFindParams struct {
	Collection string
	Connection string
	Database   string
	Filter     MongoFilter
	Fields     []string
	Options    FindOptions
}

func (param MongoFindParams) findOne() DataResult {
	document := NewMongoDocumentFind(param)
	return document.FindOne()
}

func (param MongoFindParams) find() DataResult {
	document := NewMongoDocumentFind(param)
	return document.Find()
}

type MongoDeleteParams struct {
	Collection string
	Connection string
	Database   string
	DataLog    DataLog
	FindParams MongoFindParams
}

/*func (param MongoDeleteParams) deleteOne() DataResult {
	document := NewMongoDocumentDelete(param)
	return document.DeleteOne()
}*/

var DbMong *mongo.Client

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
}

func GetCollection(mongoClient *mongo.Client, database string, collectionName string) *mongo.Collection {
	return mongoClient.Database(database).Collection(collectionName)
}
