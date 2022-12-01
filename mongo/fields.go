package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
)

type Fields struct {
	Values func() bson.D
}

func NewFields(listFields []string) Fields {

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

	return Fields{
		Values: func() bson.D {
			return _projection()
		},
	}
}
