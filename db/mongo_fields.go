package db

import (
	"go.mongodb.org/mongo-driver/bson"
)

type MongoFields struct {
	Values func() bson.D
}

func NewMongoFields(listFields []string) MongoFields {

	_projection := func() bson.D {
		fields := bson.D{{"_id", 1}}
		if len(listFields) > 0 {
			for _, item := range listFields {
				fields = append(fields, bson.E{item, 1})
			}
		}
		return fields
	}

	/*if param.Info != nil {
		fmt.Println("Not Implemented Info")
	}*/

	return MongoFields{
		Values: func() bson.D {
			return _projection()
		},
	}
}
