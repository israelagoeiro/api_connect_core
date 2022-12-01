package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type DocumentUpdate struct {
	UpdateOne        func() DataResult
	UpdateMany       func() DataResult
	FindOneAndUpdate func() DataResult
}

type UpdateOptions struct {
	ReturnOriginal bool
	Upsert         bool
}

type UpdateParams struct {
	Collection    string
	Connection    string
	Database      string
	DataLog       DataLog
	Fields        []string
	Filter        MongoFilter
	FindParams    FindParams
	Info          any
	Input         UpdateInput
	UpdateOptions UpdateOptions
}

func (param UpdateParams) _mongoFindOneAndUpdate() DataResult {
	document := NewDocumentUpdate(param)
	return document.FindOneAndUpdate()
}

func FindOneAndUpdate(param UpdateInterface) DataResult {
	return param._mongoFindOneAndUpdate()
}

func (param UpdateParams) _mongoUpdateOne() DataResult {
	document := NewDocumentUpdate(param)
	return document.UpdateOne()
}
func UpdateOne(param UpdateInterface) DataResult {
	return param._mongoUpdateOne()
}

func (param UpdateParams) _mongoUpdateMany() DataResult {
	document := NewDocumentUpdate(param)
	return document.UpdateMany()
}
func UpdateMany(param UpdateInterface) DataResult {
	return param._mongoUpdateMany()
}

func NewDocumentUpdate(param UpdateParams) DocumentUpdate {
	mongoDataLog := NewMongoDataLog(param.DataLog)

	_options := func() *options.FindOneAndUpdateOptions {
		returnOriginal := 0
		if param.UpdateOptions.ReturnOriginal {
			returnOriginal = 1
		}

		apiFields := NewFields(param.Fields)
		fmt.Println("apiFields.Values()", apiFields.Values())

		return options.FindOneAndUpdate().
			SetProjection(apiFields.Values()).
			SetUpsert(param.UpdateOptions.Upsert).
			SetReturnDocument(options.ReturnDocument(returnOriginal))
	}

	apiDocumentUpdate := DocumentUpdate{
		FindOneAndUpdate: func() DataResult {
			return DataResult{}
		},
		UpdateMany: func() DataResult {
			return DataResult{}
		},
		UpdateOne: func() (result DataResult) {
			mongoDataLog.PrepareUpdate(param.Input)

			if !param.Input.IsValid() {
				fmt.Println("UpdateOne::UpdateParams::Input", param.Input.Values())
				panic("UpdateOne::UpdateParams::Input - Documento de atualização deve ter pelo menos um elemento 'input.Set(?) ou input.AddToSet(?)'")
			}

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			collection := GetCollection(param.Database, param.Collection)

			if param.FindParams.Collection != "" {
				data := collection.FindOneAndUpdate(ctx, param.Filter.Values(), param.Input.Values(), _options())
				result._decode = data.Decode
				result._debug = func() {
					var dataResult any
					err := data.Decode(&dataResult)
					if err != nil {
						fmt.Println("Error UpdateOne::UpdateParams:Debug()")
						return
					}
					fmt.Println("UpdateOne::UpdateParams:Debug()", "Collection:"+param.Collection, dataResult)
				}
				result._result = func() interface{} {
					var dataResult interface{}
					err := data.Decode(&dataResult)
					if err != nil {
						fmt.Println("Error UpdateOne::UpdateParams:Result()")
						return nil
					}
					return dataResult
				}
			} else {
				data, err := collection.UpdateOne(ctx, param.Filter.Values(), param.Input.Values())
				if err != nil {
					fmt.Println("Error UpdateOne::UpdateParams:", err.Error())
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
