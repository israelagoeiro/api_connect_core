package mongo

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Eq(value any) bson.D {
	return bson.D{bson.E{Key: "$eq", Value: value}}
}
func EqField(field string, value any) bson.D {
	return bson.D{bson.E{Key: field, Value: Eq(value)}}
}

func Gt(value any) bson.D {
	return bson.D{bson.E{Key: "$gt", Value: value}}
}

func GtField(field string, value any) bson.D {
	return bson.D{bson.E{Key: field, Value: Gt(value)}}
}

func Gte(value any) bson.D {
	return bson.D{bson.E{Key: "$gte", Value: value}}
}
func GteField(field string, value any) bson.D {
	return bson.D{bson.E{Key: field, Value: Gte(value)}}
}

func Lt(value any) bson.D {
	return bson.D{bson.E{Key: "$lt", Value: value}}
}

func LtField(field string, value any) bson.D {
	return bson.D{bson.E{Key: field, Value: Lt(value)}}
}

func Lte(value any) bson.D {
	return bson.D{bson.E{Key: "$lte", Value: value}}
}
func LteField(field string, value any) bson.D {
	return bson.D{bson.E{Key: field, Value: Lte(value)}}
}

func Ne(value any) bson.D {
	return bson.D{bson.E{Key: "$ne", Value: value}}
}
func NeField(field string, value any) bson.D {
	return bson.D{bson.E{Key: field, Value: Ne(value)}}
}

func In(values []any) bson.D {
	return bson.D{bson.E{Key: "$in", Value: values}}
}
func InField(field string, values []any) bson.D {
	return bson.D{bson.E{Key: field, Value: In(values)}}
}

func Nin(values []any) bson.D {
	return bson.D{bson.E{Key: "$nin", Value: values}}
}
func NinField(field string, values []any) bson.D {
	return bson.D{bson.E{Key: field, Value: Nin(values)}}
}

func Exists(value bool) bson.D {
	return bson.D{bson.E{Key: "$exists", Value: value}}
}
func ExistsField(field string, value bool) bson.D {
	return bson.D{bson.E{Key: field, Value: Exists(value)}}
}

func Type(value string) bson.D {
	return bson.D{bson.E{Key: "$type", Value: value}}
}
func TypeField(field string, value string) bson.D {
	return bson.D{bson.E{Key: field, Value: Type(value)}}
}

func Divide(values []any) bson.D {
	return bson.D{bson.E{Key: "$divide", Value: values}}
}
func DivideField(field string, values []any) bson.D {
	return bson.D{bson.E{Key: field, Value: Divide(values)}}
}

func Expr(values []any) bson.D {
	return bson.D{bson.E{Key: "$expr", Value: values}}
}
func ExprField(operator string, values []any) bson.D {
	return bson.D{bson.E{Key: "$expr", Value: bson.E{Key: operator, Value: values}}}
}

func Not(value any) bson.D {
	return bson.D{bson.E{Key: "$not", Value: value}}
}
func NotField(field string, value bson.D) bson.D {
	return bson.D{bson.E{Key: field, Value: Not(value)}}
}
func Regex(value any) bson.D {
	return bson.D{bson.E{Key: "$regex", Value: value}}
}

func RegexField(field string, regex any, options string) bson.D {
	return bson.D{bson.E{Key: field, Value: bson.D{bson.E{"$regex", regex}, bson.E{"$options", options}}}}
}
func RegexOption(value any, options string) bson.D {
	return bson.D{bson.E{"$regex", value}, bson.E{"$options", options}}
}

func Sort(field string, asc bool) bson.D {
	_direction := -1
	if asc {
		_direction = 1
	}
	return bson.D{{field, _direction}}
}

type MongoFilter struct {
	Id       func(value string)
	ObjectId func(value primitive.ObjectID)
	Add      func(field string, operator bson.D)
	Expr     func(operator bson.D)
	Regex    func(field string, operator bson.D)
	And      func(operator []any)
	Nor      func(operator []any)
	Or       func(operator []any)
	Values   func() bson.D
	Debug    func()
}

func NewFilter() MongoFilter {
	result := bson.D{}

	return MongoFilter{
		Id: func(value string) {
			objectId, _ := primitive.ObjectIDFromHex(value)
			result = append(result, bson.E{Key: "_id", Value: objectId})
		},
		ObjectId: func(value primitive.ObjectID) {
			result = append(result, bson.E{Key: "_id", Value: value})
		},
		Add: func(field string, operator bson.D) {
			result = append(result, bson.E{Key: field, Value: operator})
		},
		Nor: func(values []any) {
			listNor := bson.A{}
			for _, item := range values {
				listNor = append(listNor, item)
			}
			result = append(result, bson.E{Key: "$nor", Value: listNor})
		},
		Or: func(values []any) {
			listOr := bson.A{}
			for _, item := range values {
				listOr = append(listOr, item)
			}
			result = append(result, bson.E{Key: "$or", Value: listOr})
		},
		Expr: func(operator bson.D) {
			result = append(result, bson.E{Key: "$expr", Value: operator})
		},
		And: func(values []any) {
			listAnd := bson.A{}
			for _, item := range values {
				listAnd = append(listAnd, item)
			}
			result = append(result, bson.E{Key: "$and", Value: listAnd})
		},
		Regex: func(field string, values bson.D) {
			result = append(result, bson.E{Key: field, Value: values})
		},
		Values: func() bson.D {
			return result
		},
		Debug: func() {
			for i, item := range result {
				fmt.Println(i, item)
			}
		},
	}
}
