package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MongoInsertParams struct {
	Collection string
	Connection string
	Database   string
	DataLog    DataLog
	Input      MongoInputInsert
	FindParams MongoFindParams
}

func (param MongoInsertParams) _mongoInsertOne() DataResult {
	document := NewMongoDocumentInsert(param)
	return document.InsertOne()
}

func InsertOne(param InsertInterface) DataResult {
	return param._mongoInsertOne()
}

func (param MongoInsertParams) _mongoInsertMany() DataResult {
	document := NewMongoDocumentInsert(param)
	return document.InsertOne()
}

func InsertMany(param InsertInterface) DataResult {
	return param._mongoInsertMany()
}

type MongoDocumentInsert struct {
	InsertMany func() DataResult
	InsertOne  func() DataResult
}

func NewMongoDocumentInsert(param MongoInsertParams) MongoDocumentInsert {
	//apiFields := NewMongoFields(param)
	mongoDataLog := NewMongoDataLog(param.DataLog)

	apiDocumentUpdate := MongoDocumentInsert{
		InsertMany: func() DataResult {
			mongoDataLog.PrepareInsert(param.Input)

			if !param.Input.IsValid() {
				fmt.Println("InsertOne::MongoInsertParams::Input", param.Input.Values())
				panic("InsertOne::MongoInsertParams::Input - Documento de inserção deve ter pelo menos um elemento 'input.Data(?) ou input.Map(?)'")
			}

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			collection := GetCollection(DbMong, param.Database, param.Collection)
			//param.Input.Values()
			docs := []interface{}{
				bson.D{{"type", "English Breakfast"}, {"rating", 6}},
			}

			data, err := collection.InsertMany(ctx, docs)

			if err != nil {
				fmt.Println("Error MongoDocumentInsert:InsertOne:", err.Error())
			}

			return DataResult{
				_id: data.InsertedIDs,
			}
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
