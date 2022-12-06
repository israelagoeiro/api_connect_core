package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type FindOptions struct {
	Sort bson.D
}

type FindParams struct {
	Collection string
	Connection string
	Database   string
	Filter     MongoFilter
	Fields     []string
	Options    FindOptions
}

type FindDocument struct {
	Find    func() DataResult
	FindOne func() DataResult
}

func NewFindDocument(param FindParams) FindDocument {
	apiFields := NewFields(param.Fields)
	updateDocument := FindDocument{
		Find: func() DataResult {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			collection := GetCollection(param.Database, param.Collection)
			opts := options.Find()

			if param.Options.Sort != nil {
				opts.SetSort(param.Options.Sort)
			}

			if param.Fields != nil {
				opts.SetProjection(apiFields.Values())
			}

			filter := bson.D{}
			if param.Filter.Values != nil {
				filter = param.Filter.Values()
			}
			result, _ := collection.Find(ctx, filter, opts)

			return DataResult{
				_all:         result.All,
				_ctx:         ctx,
				_cancel:      cancel,
				_mongoResult: result,
				_debug: func() {
					var dataResult []any
					if err := result.All(ctx, &dataResult); err != nil {
						fmt.Println("Error Find:Debug()", err.Error())
					}
					err := result.Close(ctx)
					if err != nil {
						fmt.Println("Error Find:Debug()", err.Error())
						cancel()
					}
					fmt.Println("Find:Debug()", "Collection:"+param.Collection, dataResult)
				},
			}
		},
		FindOne: func() DataResult {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			collection := GetCollection(param.Database, param.Collection)
			opts := options.FindOne()

			if param.Options.Sort != nil {
				opts.SetSort(param.Options.Sort)
			}

			if param.Fields != nil {
				opts.SetProjection(apiFields.Values())
			}

			filter := bson.D{}
			if param.Filter.Values != nil {
				filter = param.Filter.Values()
			}
			result := collection.FindOne(ctx, filter, opts)

			return DataResult{
				_decode: result.Decode,
				_debug: func() {
					var dataResult any
					err := result.Decode(&dataResult)
					if err != nil {
						fmt.Println("Error FindOne:Debug()", err.Error())
						return
					}
					fmt.Println("FindOne:Debug()", "Collection:"+param.Collection, dataResult)
				},
			}
		},
	}
	return updateDocument
}
