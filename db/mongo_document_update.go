package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoDocumentUpdate struct {
	UpdateOne  func() DataResult
	UpdateMany func() DataResult
}

func NewMongoDocumentUpdate(param MongoUpdateParams) MongoDocumentUpdate {
	mongoDataLog := NewMongoDataLog(param.DataLog)

	_options := func() *options.FindOneAndUpdateOptions {
		returnOriginal := 0
		if param.UpdateOptions.ReturnOriginal {
			returnOriginal = 1
		}

		apiFields := NewMongoFields(param.Fields)
		fmt.Println("apiFields.Values()", apiFields.Values())

		return options.FindOneAndUpdate().
			SetProjection(apiFields.Values()).
			SetUpsert(param.UpdateOptions.Upsert).
			SetReturnDocument(options.ReturnDocument(returnOriginal))
	}

	apiDocumentUpdate := MongoDocumentUpdate{
		UpdateMany: func() DataResult {
			return DataResult{}
		},
		UpdateOne: func() (result DataResult) {
			mongoDataLog.PrepareUpdate(param.Input)

			if !param.Input.IsValid() {
				fmt.Println("UpdateOne::MongoUpdateParams::Input", param.Input.Values())
				panic("UpdateOne::MongoUpdateParams::Input - Documento de atualização deve ter pelo menos um elemento 'input.Set(?) ou input.AddToSet(?)'")
			}

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			collection := GetCollection(DbMong, param.Database, param.Collection)

			if param.FindParams.Collection != "" {
				data := collection.FindOneAndUpdate(ctx, param.Filter.Values(), param.Input.Values(), _options())
				result._decode = data.Decode
				result._debug = func() {
					var dataResult any
					err := data.Decode(&dataResult)
					if err != nil {
						fmt.Println("Error UpdateOne::MongoUpdateParams:Debug()")
						return
					}
					fmt.Println("UpdateOne::MongoUpdateParams:Debug()", "Collection:"+param.Collection, dataResult)
				}
				result._result = func() interface{} {
					var dataResult interface{}
					err := data.Decode(&dataResult)
					if err != nil {
						fmt.Println("Error UpdateOne::MongoUpdateParams:Result()")
						return nil
					}
					return dataResult
				}
			} else {
				data, err := collection.UpdateOne(ctx, param.Filter.Values(), param.Input.Values())
				if err != nil {
					fmt.Println("Error UpdateOne::MongoUpdateParams:", err.Error())
				} else {
					result._modifiedCount = data.ModifiedCount
					result._matchedCount = data.MatchedCount
					result._upsertedCount = data.UpsertedCount
					result._upsertedID = data.UpsertedID
				}
			}
			return result
		},
	}

	return apiDocumentUpdate
}