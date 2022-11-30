package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoDocumentInsert struct {
	Collection string
	Options    *options.FindOneAndUpdateOptions
	InsertMany func() DataResult
	InsertOne  func() DataResult
}

func NewMongoDocumentInsert(param MongoInsertParams) MongoDocumentInsert {
	//apiFields := NewMongoFields(param)
	mongoDataLog := NewMongoDataLog(param.DataLog)

	apiDocumentUpdate := MongoDocumentInsert{
		InsertMany: func() DataResult {
			return DataResult{}
		},
		InsertOne: func() (result DataResult) {
			mongoDataLog.PrepareInsert(param.Input)

			if !param.Input.IsValid() {
				fmt.Println("InsertOne::MongoInsertParams::Input", param.Input.Values())
				panic("InsertOne::MongoInsertParams::Input - Documento de inserção deve ter pelo menos um elemento 'input.Data(?) ou input.Map(?)'")
			}

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			collection := GetCollection(DbMong, param.Database, param.Collection)
			data, err := collection.InsertOne(ctx, param.Input.Values())

			if err != nil {
				fmt.Println("Error MongoDocumentInsert:InsertOne:", err.Error())
			} else {
				result._id = data.InsertedID
				filter := NewMongoFilter()
				objectId := data.InsertedID.(primitive.ObjectID)
				filter.ObjectId(objectId)
				param.FindParams.Filter = filter
				if param.FindParams.Collection != "" {
					result = FindOne(param.FindParams)
				}
			}
			return result
		},
	}

	return apiDocumentUpdate
}
