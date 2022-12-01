package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type InsertParams struct {
	Collection string
	Connection string
	Database   string
	DataLog    DataLog
	Input      InsertInput
	FindParams FindParams
}

type DocumentInsert struct {
	InsertMany func() DataResult
	InsertOne  func() DataResult
}

func NewDocumentInsert(param InsertParams) DocumentInsert {
	//apiFields := NewFields(param)
	mongoDataLog := NewMongoDataLog(param.DataLog)

	apiDocumentUpdate := DocumentInsert{
		InsertMany: func() DataResult {
			mongoDataLog.PrepareInsert(param.Input)

			if !param.Input.IsValid() {
				fmt.Println("InsertOne::InsertParams::Input", param.Input.Values())
				panic("InsertOne::InsertParams::Input - Documento de inserção deve ter pelo menos um elemento 'input.Data(?) ou input.Map(?)'")
			}

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			collection := GetCollection(param.Database, param.Collection)
			//param.Input.Values()
			docs := []interface{}{
				bson.D{{"type", "English Breakfast"}, {"rating", 6}},
			}

			data, err := collection.InsertMany(ctx, docs)

			if err != nil {
				fmt.Println("Error DocumentInsert:InsertOne:", err.Error())
			}

			return DataResult{
				_id: data.InsertedIDs,
			}
		},
		InsertOne: func() (result DataResult) {
			mongoDataLog.PrepareInsert(param.Input)

			if !param.Input.IsValid() {
				fmt.Println("InsertOne::InsertParams::Input", param.Input.Values())
				panic("InsertOne::InsertParams::Input - Documento de inserção deve ter pelo menos um elemento 'input.Data(?) ou input.Map(?)'")
			}

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			collection := GetCollection(param.Database, param.Collection)
			data, err := collection.InsertOne(ctx, param.Input.Values())

			if err != nil {
				fmt.Println("Error DocumentInsert:InsertOne:", err.Error())
			} else {
				result._id = data.InsertedID
				filter := NewFilter()
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
