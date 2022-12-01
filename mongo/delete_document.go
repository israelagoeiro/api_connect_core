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

type DeleteDocument struct {
	DeleteMany func() DataResult
	DeleteOne  func() DataResult
}

func NewDeleteDocument(param DeleteParams) DeleteDocument {
	apiDocumentDelete := DeleteDocument{
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
