package db

import (
	"go.mongodb.org/mongo-driver/bson"
)

type MongoInputUpdate struct {
	AddToSet    func(field string, value any)
	AddToSetMap func(values map[string]any)
	Inc         func(field string, value any)
	IncMap      func(values map[string]any)
	Set         func(field string, value any)
	SetMap      func(values map[string]any)
	IsValid     func() bool
	Values      func() bson.D
}

func NewMongoInputUpdate() MongoInputUpdate {
	_isValid := false
	_listInc := map[string]any{}
	_listSet := map[string]any{}
	_listAddToSet := map[string]any{}

	return MongoInputUpdate{
		AddToSet: func(field string, value any) {
			_listAddToSet[field] = value
			_isValid = true
		},
		AddToSetMap: func(values map[string]any) {
			for field, value := range values {
				_listAddToSet[field] = value
			}
			_isValid = true
		},
		Inc: func(field string, value any) {
			_listInc[field] = value
			_isValid = true
		},
		IncMap: func(values map[string]any) {
			for field, value := range values {
				_listInc[field] = value
			}
			_isValid = true
		},
		Set: func(field string, value any) {
			_listSet[field] = value
			_isValid = true
		},
		SetMap: func(values map[string]any) {
			for field, value := range values {
				_listSet[field] = value
			}
			_isValid = true
		},
		IsValid: func() bool {
			return _isValid
		},
		Values: func() bson.D {
			result := bson.D{}

			if len(_listAddToSet) > 0 {
				listBson := bson.D{}
				for k, v := range _listAddToSet {
					listBson = append(listBson, bson.E{Key: k, Value: v})
				}
				result = append(result, bson.E{Key: "$addToSet", Value: listBson})
			}

			if len(_listInc) > 0 {
				listBson := bson.D{}
				for k, v := range _listInc {
					listBson = append(listBson, bson.E{Key: k, Value: v})
				}
				result = append(result, bson.E{Key: "$inc", Value: listBson})
			}

			if len(_listSet) > 0 {
				listBson := bson.D{}
				for k, v := range _listSet {
					listBson = append(listBson, bson.E{Key: k, Value: v})
				}
				result = append(result, bson.E{Key: "$set", Value: listBson})
			}
			return result
		},
	}
}
