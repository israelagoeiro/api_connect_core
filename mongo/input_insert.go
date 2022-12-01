package mongo

import (
	"fmt"
	"github.com/israelagoeiro/api_connect_core/util"
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
)

type InsertInput struct {
	Model   func(values any)
	Data    func(field string, values any)
	DataMap func(values map[string]any)
	IsValid func() bool
	Values  func() bson.D
}

func toDoc(v interface{}) (doc bson.D, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	return
}
func NewInsertInput() InsertInput {
	_data := bson.D{}
	_isValid := false
	_listData := map[string]any{}

	return InsertInput{
		Model: func(value interface{}) {
			_data, _ = toDoc(value)
			_isValid = true
		},
		Data: func(field string, value any) {
			_listData[field] = value
			_isValid = true
		},
		DataMap: func(values map[string]any) {
			for field, value := range values {
				_listData[field] = value
			}
			_isValid = true
		},
		IsValid: func() bool {
			return _isValid
		},
		Values: func() bson.D {
			result := bson.D{}
			if _data != nil {
				result = _data
			}

			keys := make([]string, 0, len(result))
			for _, v := range result {
				keys = append(keys, fmt.Sprintf("%v", reflect.ValueOf(v).Field(0)))
			}

			if len(_listData) > 0 {
				for k, v := range _listData {
					if util.ContainsStr(keys, k) {
						fmt.Println("InsertParams::Input", result, "Error:", "{", k, ":", v, "}")
						panic("Error InsertParams: Foi encontrado uma duplicidade de parâmetros em '" + k + "'.")
					}
					result = append(result, bson.E{Key: k, Value: v})
				}

			}
			return result
		},
	}
}
