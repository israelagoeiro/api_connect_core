package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type DeleteParams struct {
	Collection string
	Connection string
	Database   string
	DataLog    DataLog
	FindParams FindParams
	Filter     MongoFilter
}

func DeleteOne(param DeleteInterface) DataResult {
	return param.deleteOne()
}
func DeleteMany(param DeleteInterface) DataResult {
	return param.deleteMany()
}

func (param DeleteParams) deleteMany() DataResult {
	document := NewDocumentDelete(param)
	return document.DeleteMany()
}

func (param DeleteParams) deleteOne() DataResult {
	document := NewDocumentDelete(param)
	return document.DeleteOne()
}

type DocumentDelete struct {
	DeleteMany func() DataResult
	DeleteOne  func() DataResult
}

func NewDocumentDelete(param DeleteParams) DocumentDelete {
	apiDocumentDelete := DocumentDelete{
		DeleteMany: func() DataResult {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			collection := GetCollection(param.Database, param.Collection)
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
			collection := GetCollection(param.Database, param.Collection)
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
