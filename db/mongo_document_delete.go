package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type MongoDeleteParams struct {
	Collection string
	Connection string
	Database   string
	DataLog    DataLog
	FindParams MongoFindParams
	Filter     MongoFilter
}

func DeleteOne(param DeleteInterface) DataResult {
	return param.deleteOne()
}
func DeleteMany(param DeleteInterface) DataResult {
	return param.deleteMany()
}

func (param MongoDeleteParams) deleteMany() DataResult {
	document := NewMongoDocumentDelete(param)
	return document.DeleteMany()
}

func (param MongoDeleteParams) deleteOne() DataResult {
	document := NewMongoDocumentDelete(param)
	return document.DeleteOne()
}

type MongoDocumentDelete struct {
	DeleteMany func() DataResult
	DeleteOne  func() DataResult
}

func NewMongoDocumentDelete(param MongoDeleteParams) MongoDocumentDelete {
	apiDocumentDelete := MongoDocumentDelete{
		DeleteMany: func() DataResult {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			collection := GetCollection(DbMong, param.Database, param.Collection)
			filter := bson.D{}
			if param.Filter.Values != nil {
				filter = param.Filter.Values()
			}
			result, _ := collection.DeleteMany(ctx, filter)

			return DataResult{
				_count:  result.DeletedCount,
				_ctx:    ctx,
				_cancel: cancel,
			}
		},
		DeleteOne: func() DataResult {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			collection := GetCollection(DbMong, param.Database, param.Collection)
			filter := bson.D{}
			if param.Filter.Values != nil {
				filter = param.Filter.Values()
			}
			result, _ := collection.DeleteOne(ctx, filter)

			return DataResult{
				_count:  result.DeletedCount,
				_ctx:    ctx,
				_cancel: cancel,
			}
		},
	}
	return apiDocumentDelete
}
