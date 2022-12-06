package mongo_test

import (
	"fmt"
	"github.com/israelagoeiro/api_connect_core/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
	"testing"
)

func inspect(a any, b any) bool {
	fmt.Println("AB-->>", a, b)

	valA := reflect.ValueOf(a)
	valB := reflect.ValueOf(b)
	fmt.Println("valA-->>", valA, valA.Kind())
	fmt.Println("valB-->>", valB, valB.Kind())

	switch valA.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return a == b
	case reflect.String:
	case reflect.Slice:
		if valA.Len() != valB.Len() {
			return false
		}
		for i := 0; i < valA.Len(); i++ {
			inspect(reflect.ValueOf(valA.Index(i)), reflect.ValueOf(valA.Index(i)))
		}
		//valX := valA.NumField()
		fmt.Println("valX--->>>", valA.Len())
	case reflect.Struct:
		fmt.Println("STRUCT", reflect.ValueOf(valA).Bytes(), b)
	}

	return true
}

func getMap(dataA interface{}) map[string]string {
	result := map[string]string{}
	for i := 0; i < reflect.ValueOf(dataA).Len(); i++ {
		field := fmt.Sprintf("%v", reflect.ValueOf(dataA).Index(i).Field(0))
		value := fmt.Sprintf("%v", reflect.ValueOf(dataA).Index(i).Field(1))
		result[field] = value
	}
	return result
}

func compareMap(a map[string]string, b map[string]string) bool {
	keys := make([]string, 0, len(a))
	for k := range a {
		keys = append(keys, k)
	}
	status := true
	for _, k := range keys {
		if a[k] != b[k] {
			status = false
			break
		}
	}
	return status
}

func IsValidEq(a, b bson.D) bool {
	if len(a) != len(b) {
		return false
	}

	mapA1 := map[string]string{}
	mapA2 := map[string]string{}
	operatorA := true
	for i := range a {
		for j := 0; j < reflect.ValueOf(a[i].Value).Len(); j++ {
			mapA1[a[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(a[i].Value).Index(j).Field(0))
			mapA2[a[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(a[i].Value).Index(j).Field(1))
			if mapA1[a[i].Key] != "$eq" {
				operatorA = false
				break
			}
		}
	}

	mapB1 := map[string]string{}
	mapB2 := map[string]string{}
	operatorB := true
	for i := range b {
		for j := 0; j < reflect.ValueOf(b[i].Value).Len(); j++ {
			mapB1[b[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(b[i].Value).Index(j).Field(0))
			mapB2[b[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(b[i].Value).Index(j).Field(1))
		}
		if mapB1[a[i].Key] != "$eq" {
			operatorB = false
			break
		}
	}

	if operatorA != operatorB {
		fmt.Println("Error IsValidEq: Operador de comparação do tipo Eq(?) não é válido")
		return false
	}

	if !compareMap(mapA1, mapB1) || !compareMap(mapA2, mapB2) {
		fmt.Println("Error IsValidEq: O identificador de campo do operador de comparação do tipo Eq(?) não é válido")
		return false
	}

	return true
}

func IsValidGt(a, b bson.D) bool {
	if len(a) != len(b) {
		return false
	}

	mapA1 := map[string]string{}
	mapA2 := map[string]string{}
	operatorA := true
	for i := range a {
		for j := 0; j < reflect.ValueOf(a[i].Value).Len(); j++ {
			mapA1[a[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(a[i].Value).Index(j).Field(0))
			mapA2[a[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(a[i].Value).Index(j).Field(1))
			if mapA1[a[i].Key] != "$gt" {
				operatorA = false
				break
			}
		}
	}

	mapB1 := map[string]string{}
	mapB2 := map[string]string{}
	operatorB := true
	for i := range b {
		for j := 0; j < reflect.ValueOf(b[i].Value).Len(); j++ {
			mapB1[b[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(b[i].Value).Index(j).Field(0))
			mapB2[b[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(b[i].Value).Index(j).Field(1))
		}
		if mapB1[a[i].Key] != "$gt" {
			operatorB = false
			break
		}
	}

	if operatorA != operatorB {
		fmt.Println("Error IsValidGt: Operador de comparação do tipo Gt(?) não é válido")
		return false
	}

	if !compareMap(mapA1, mapB1) || !compareMap(mapA2, mapB2) {
		fmt.Println("Error IsValidGt: O identificador de campo do operador de comparação do tipo Gt(?) não é válido")
		return false
	}

	return true
}

func IsValidGte(a, b bson.D) bool {
	if len(a) != len(b) {
		return false
	}

	mapA1 := map[string]string{}
	mapA2 := map[string]string{}
	operatorA := true
	for i := range a {
		for j := 0; j < reflect.ValueOf(a[i].Value).Len(); j++ {
			mapA1[a[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(a[i].Value).Index(j).Field(0))
			mapA2[a[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(a[i].Value).Index(j).Field(1))
			if mapA1[a[i].Key] != "$gte" {
				operatorA = false
				break
			}
		}
	}

	mapB1 := map[string]string{}
	mapB2 := map[string]string{}
	operatorB := true
	for i := range b {
		for j := 0; j < reflect.ValueOf(b[i].Value).Len(); j++ {
			mapB1[b[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(b[i].Value).Index(j).Field(0))
			mapB2[b[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(b[i].Value).Index(j).Field(1))
		}
		if mapB1[a[i].Key] != "$gte" {
			operatorB = false
			break
		}
	}

	if operatorA != operatorB {
		fmt.Println("Error IsValidGte: Operador de comparação do tipo Gte(?) não é válido")
		return false
	}

	if !compareMap(mapA1, mapB1) || !compareMap(mapA2, mapB2) {
		fmt.Println("Error IsValidGte: O identificador de campo do operador de comparação do tipo Gte(?) não é válido")
		return false
	}

	return true
}

func IsValidLt(a, b bson.D) bool {
	if len(a) != len(b) {
		return false
	}

	mapA1 := map[string]string{}
	mapA2 := map[string]string{}
	operatorA := true
	for i := range a {
		for j := 0; j < reflect.ValueOf(a[i].Value).Len(); j++ {
			mapA1[a[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(a[i].Value).Index(j).Field(0))
			mapA2[a[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(a[i].Value).Index(j).Field(1))
			if mapA1[a[i].Key] != "$lt" {
				operatorA = false
				break
			}
		}
	}

	mapB1 := map[string]string{}
	mapB2 := map[string]string{}
	operatorB := true
	for i := range b {
		for j := 0; j < reflect.ValueOf(b[i].Value).Len(); j++ {
			mapB1[b[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(b[i].Value).Index(j).Field(0))
			mapB2[b[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(b[i].Value).Index(j).Field(1))
		}
		if mapB1[a[i].Key] != "$lt" {
			operatorB = false
			break
		}
	}

	if operatorA != operatorB {
		fmt.Println("Error IsValidLt: Operador de comparação do tipo Lt(?) não é válido")
		return false
	}

	if !compareMap(mapA1, mapB1) || !compareMap(mapA2, mapB2) {
		fmt.Println("Error IsValidLt: O identificador de campo do operador de comparação do tipo Lt(?) não é válido")
		return false
	}

	return true
}

func IsValidLte(a, b bson.D) bool {
	if len(a) != len(b) {
		return false
	}

	mapA1 := map[string]string{}
	mapA2 := map[string]string{}
	operatorA := true
	for i := range a {
		for j := 0; j < reflect.ValueOf(a[i].Value).Len(); j++ {
			mapA1[a[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(a[i].Value).Index(j).Field(0))
			mapA2[a[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(a[i].Value).Index(j).Field(1))
			if mapA1[a[i].Key] != "$lte" {
				operatorA = false
				break
			}
		}
	}

	mapB1 := map[string]string{}
	mapB2 := map[string]string{}
	operatorB := true
	for i := range b {
		for j := 0; j < reflect.ValueOf(b[i].Value).Len(); j++ {
			mapB1[b[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(b[i].Value).Index(j).Field(0))
			mapB2[b[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(b[i].Value).Index(j).Field(1))
		}
		if mapB1[a[i].Key] != "$lte" {
			operatorB = false
			break
		}
	}

	if operatorA != operatorB {
		fmt.Println("Error IsValidLte: Operador de comparação do tipo Lte(?) não é válido")
		return false
	}

	if !compareMap(mapA1, mapB1) || !compareMap(mapA2, mapB2) {
		fmt.Println("Error IsValidLte: O identificador de campo do operador de comparação do tipo Lte(?) não é válido")
		return false
	}

	return true
}

func IsValidNe(a, b bson.D) bool {
	if len(a) != len(b) {
		return false
	}

	mapA1 := map[string]string{}
	mapA2 := map[string]string{}
	operatorA := true
	for i := range a {
		for j := 0; j < reflect.ValueOf(a[i].Value).Len(); j++ {
			mapA1[a[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(a[i].Value).Index(j).Field(0))
			mapA2[a[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(a[i].Value).Index(j).Field(1))
			if mapA1[a[i].Key] != "$ne" {
				operatorA = false
				break
			}
		}
	}

	mapB1 := map[string]string{}
	mapB2 := map[string]string{}
	operatorB := true
	for i := range b {
		for j := 0; j < reflect.ValueOf(b[i].Value).Len(); j++ {
			mapB1[b[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(b[i].Value).Index(j).Field(0))
			mapB2[b[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(b[i].Value).Index(j).Field(1))
		}
		if mapB1[a[i].Key] != "$ne" {
			operatorB = false
			break
		}
	}

	if operatorA != operatorB {
		fmt.Println("Error IsValidNe: Operador de comparação do tipo Ne(?) não é válido")
		return false
	}

	if !compareMap(mapA1, mapB1) || !compareMap(mapA2, mapB2) {
		fmt.Println("Error IsValidNe: O identificador de campo do operador de comparação do tipo Ne(?) não é válido")
		return false
	}

	return true
}

func IsValidIn(a, b bson.D) bool {
	if len(a) != len(b) {
		return false
	}

	mapA1 := map[string]string{}
	mapA2 := map[string]string{}
	operatorA := true
	for i := range a {
		for j := 0; j < reflect.ValueOf(a[i].Value).Len(); j++ {
			mapA1[a[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(a[i].Value).Index(j).Field(0))
			mapA2[a[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(a[i].Value).Index(j).Field(1))
			if mapA1[a[i].Key] != "$in" {
				operatorA = false
				break
			}
		}
	}

	mapB1 := map[string]string{}
	mapB2 := map[string]string{}
	operatorB := true
	for i := range b {
		for j := 0; j < reflect.ValueOf(b[i].Value).Len(); j++ {
			mapB1[b[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(b[i].Value).Index(j).Field(0))
			mapB2[b[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(b[i].Value).Index(j).Field(1))
		}
		if mapB1[a[i].Key] != "$in" {
			operatorB = false
			break
		}
	}

	if operatorA != operatorB {
		fmt.Println("Error IsValidIn: Operador de comparação do tipo In(?) não é válido")
		return false
	}

	if !compareMap(mapA1, mapB1) || !compareMap(mapA2, mapB2) {
		fmt.Println("Error IsValidIn: O identificador de campo do operador de comparação do tipo In(?) não é válido")
		return false
	}

	return true
}

func IsValidNin(a, b bson.D) bool {
	if len(a) != len(b) {
		return false
	}

	mapA1 := map[string]string{}
	mapA2 := map[string]string{}
	operatorA := true
	for i := range a {
		for j := 0; j < reflect.ValueOf(a[i].Value).Len(); j++ {
			mapA1[a[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(a[i].Value).Index(j).Field(0))
			mapA2[a[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(a[i].Value).Index(j).Field(1))
			if mapA1[a[i].Key] != "$nin" {
				operatorA = false
				break
			}
		}
	}

	mapB1 := map[string]string{}
	mapB2 := map[string]string{}
	operatorB := true
	for i := range b {
		for j := 0; j < reflect.ValueOf(b[i].Value).Len(); j++ {
			mapB1[b[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(b[i].Value).Index(j).Field(0))
			mapB2[b[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(b[i].Value).Index(j).Field(1))
		}
		if mapB1[a[i].Key] != "$nin" {
			operatorB = false
			break
		}
	}

	if operatorA != operatorB {
		fmt.Println("Error IsValidNin: Operador de comparação do tipo Nin(?) não é válido")
		return false
	}

	if !compareMap(mapA1, mapB1) || !compareMap(mapA2, mapB2) {
		fmt.Println("Error IsValidNin: O identificador de campo do operador de comparação do tipo Nin(?) não é válido")
		return false
	}

	return true
}

func IsValidExists(a, b bson.D) bool {
	if len(a) != len(b) {
		return false
	}

	mapA1 := map[string]string{}
	mapA2 := map[string]string{}
	operatorA := true
	for i := range a {
		for j := 0; j < reflect.ValueOf(a[i].Value).Len(); j++ {
			mapA1[a[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(a[i].Value).Index(j).Field(0))
			mapA2[a[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(a[i].Value).Index(j).Field(1))
			if mapA1[a[i].Key] != "$exists" {
				operatorA = false
				break
			}
		}
	}

	mapB1 := map[string]string{}
	mapB2 := map[string]string{}
	operatorB := true
	for i := range b {
		for j := 0; j < reflect.ValueOf(b[i].Value).Len(); j++ {
			mapB1[b[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(b[i].Value).Index(j).Field(0))
			mapB2[b[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(b[i].Value).Index(j).Field(1))
		}
		if mapB1[a[i].Key] != "$exists" {
			operatorB = false
			break
		}
	}

	if operatorA != operatorB {
		fmt.Println("Error IsValidExists: Operador de comparação do tipo Exists(?) não é válido")
		return false
	}

	if !compareMap(mapA1, mapB1) || !compareMap(mapA2, mapB2) {
		fmt.Println("Error IsValidExists: O identificador de campo do operador de comparação do tipo Exists(?) não é válido")
		return false
	}

	return true
}

func IsValidType(a, b bson.D) bool {
	if len(a) != len(b) {
		return false
	}

	mapA1 := map[string]string{}
	mapA2 := map[string]string{}
	operatorA := true
	for i := range a {
		for j := 0; j < reflect.ValueOf(a[i].Value).Len(); j++ {
			mapA1[a[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(a[i].Value).Index(j).Field(0))
			mapA2[a[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(a[i].Value).Index(j).Field(1))
			if mapA1[a[i].Key] != "$type" {
				operatorA = false
				break
			}
		}
	}

	mapB1 := map[string]string{}
	mapB2 := map[string]string{}
	operatorB := true
	for i := range b {
		for j := 0; j < reflect.ValueOf(b[i].Value).Len(); j++ {
			mapB1[b[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(b[i].Value).Index(j).Field(0))
			mapB2[b[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(b[i].Value).Index(j).Field(1))
		}
		if mapB1[a[i].Key] != "$type" {
			operatorB = false
			break
		}
	}

	if operatorA != operatorB {
		fmt.Println("Error IsValidType: Operador de comparação do tipo Type(?) não é válido")
		return false
	}

	if !compareMap(mapA1, mapB1) || !compareMap(mapA2, mapB2) {
		fmt.Println("Error IsValidType: O identificador de campo do operador de comparação do tipo Type(?) não é válido")
		return false
	}

	return true
}

func IsValidDivide(a, b bson.D) bool {
	if len(a) != len(b) {
		return false
	}

	mapA1 := map[string]string{}
	mapA2 := map[string]string{}
	operatorA := true
	for i := range a {
		for j := 0; j < reflect.ValueOf(a[i].Value).Len(); j++ {
			mapA1[a[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(a[i].Value).Index(j).Field(0))
			mapA2[a[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(a[i].Value).Index(j).Field(1))
			if mapA1[a[i].Key] != "$divide" {
				operatorA = false
				break
			}
		}
	}

	mapB1 := map[string]string{}
	mapB2 := map[string]string{}
	operatorB := true
	for i := range b {
		for j := 0; j < reflect.ValueOf(b[i].Value).Len(); j++ {
			mapB1[b[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(b[i].Value).Index(j).Field(0))
			mapB2[b[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(b[i].Value).Index(j).Field(1))
		}
		if mapB1[a[i].Key] != "$divide" {
			operatorB = false
			break
		}
	}

	if operatorA != operatorB {
		fmt.Println("Error IsValidDivide: Operador de comparação do tipo Divide(?) não é válido")
		return false
	}

	if !compareMap(mapA1, mapB1) || !compareMap(mapA2, mapB2) {
		fmt.Println("Error IsValidDivide: O identificador de campo do operador de comparação do tipo Divide(?) não é válido")
		return false
	}

	return true
}

func IsValidExpr(a, b bson.D) bool {
	if len(a) != len(b) {
		return false
	}

	mapA1 := map[string]string{}
	mapA2 := map[string]string{}
	operatorA := true
	for i := range a {
		for j := 0; j < reflect.ValueOf(a[i].Value).Len(); j++ {
			mapA1[a[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(a[i].Value).Index(j).Field(0))
			mapA2[a[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(a[i].Value).Index(j).Field(1))
			if mapA1[a[i].Key] != "$expr" {
				operatorA = false
				break
			}
		}
	}

	mapB1 := map[string]string{}
	mapB2 := map[string]string{}
	operatorB := true
	for i := range b {
		for j := 0; j < reflect.ValueOf(b[i].Value).Len(); j++ {
			mapB1[b[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(b[i].Value).Index(j).Field(0))
			mapB2[b[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(b[i].Value).Index(j).Field(1))
		}
		if mapB1[a[i].Key] != "$expr" {
			operatorB = false
			break
		}
	}

	if operatorA != operatorB {
		fmt.Println("Error IsValidExpr: Operador de comparação do tipo Expr(?) não é válido")
		return false
	}

	if !compareMap(mapA1, mapB1) || !compareMap(mapA2, mapB2) {
		fmt.Println("Error IsValidExpr: O identificador de campo do operador de comparação do tipo Expr(?) não é válido")
		return false
	}

	return true
}

func IsValidNot(a, b bson.D) bool {
	if len(a) != len(b) {
		return false
	}

	mapA1 := map[string]string{}
	mapA2 := map[string]string{}
	operatorA := true
	for i := range a {
		for j := 0; j < reflect.ValueOf(a[i].Value).Len(); j++ {
			mapA1[a[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(a[i].Value).Index(j).Field(0))
			mapA2[a[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(a[i].Value).Index(j).Field(1))
			if mapA1[a[i].Key] != "$not" {
				operatorA = false
				break
			}
		}
	}

	mapB1 := map[string]string{}
	mapB2 := map[string]string{}
	operatorB := true
	for i := range b {
		for j := 0; j < reflect.ValueOf(b[i].Value).Len(); j++ {
			mapB1[b[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(b[i].Value).Index(j).Field(0))
			mapB2[b[i].Key] = fmt.Sprintf("%v", reflect.ValueOf(b[i].Value).Index(j).Field(1))
		}
		if mapB1[a[i].Key] != "$not" {
			operatorB = false
			break
		}
	}

	if operatorA != operatorB {
		fmt.Println("Error IsValidNot: Operador de comparação do tipo Not(?) não é válido")
		return false
	}

	if !compareMap(mapA1, mapB1) || !compareMap(mapA2, mapB2) {
		fmt.Println("Error IsValidNot: O identificador de campo do operador de comparação do tipo Not(?) não é válido")
		return false
	}

	return true
}

func TestEq(t *testing.T) {
	var filter = mongo.NewFilter()
	filter.Add("gato", mongo.Eq("verde"))
	filter.Add("bola", mongo.Eq("azul"))

	expected := filter.Values()
	obtained := bson.D{bson.E{Key: "gato", Value: bson.D{bson.E{Key: "$eq", Value: "verde"}}}, bson.E{Key: "bola", Value: bson.D{bson.E{Key: "$eq", Value: "azul"}}}}
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if !IsValidEq(expected, obtained) {
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}
func TestEqField(t *testing.T) {
	expected := mongo.EqField("tags", "B")
	obtained := bson.D{bson.E{Key: "tags", Value: bson.D{bson.E{Key: "$eq", Value: "B"}}}}
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if !IsValidEq(expected, obtained) {
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}

func TestGt(t *testing.T) {
	var filter = mongo.NewFilter()
	filter.Add("gato", mongo.Gt("verde"))
	filter.Add("bola", mongo.Gt("azul"))

	expected := filter.Values()
	obtained := bson.D{bson.E{Key: "gato", Value: bson.D{bson.E{Key: "$gt", Value: "verde"}}}, bson.E{Key: "bola", Value: bson.D{bson.E{Key: "$gt", Value: "azul"}}}}
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if !IsValidGt(expected, obtained) {
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}
func TestGtField(t *testing.T) {
	expected := mongo.GtField("tags", "B")
	obtained := bson.D{bson.E{Key: "tags", Value: bson.D{bson.E{Key: "$gt", Value: "B"}}}}
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if !IsValidGt(expected, obtained) {
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}

func TestGte(t *testing.T) {
	var filter = mongo.NewFilter()
	filter.Add("gato", mongo.Gte("verde"))
	filter.Add("bola", mongo.Gte("azul"))

	expected := filter.Values()
	obtained := bson.D{bson.E{Key: "gato", Value: bson.D{bson.E{Key: "$gte", Value: "verde"}}}, bson.E{Key: "bola", Value: bson.D{bson.E{Key: "$gte", Value: "azul"}}}}
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if !IsValidGte(expected, obtained) {
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}
func TestGteField(t *testing.T) {
	expected := mongo.GteField("tags", "B")
	obtained := bson.D{bson.E{Key: "tags", Value: bson.D{bson.E{Key: "$gte", Value: "B"}}}}
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if !IsValidGte(expected, obtained) {
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}

func TestLt(t *testing.T) {
	var filter = mongo.NewFilter()
	filter.Add("gato", mongo.Lt("verde"))
	filter.Add("bola", mongo.Lt("azul"))

	expected := filter.Values()
	obtained := bson.D{bson.E{Key: "gato", Value: bson.D{bson.E{Key: "$lt", Value: "verde"}}}, bson.E{Key: "bola", Value: bson.D{bson.E{Key: "$lt", Value: "azul"}}}}
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if !IsValidLt(expected, obtained) {
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}
func TestLtField(t *testing.T) {
	expected := mongo.LtField("tags", "B")
	obtained := bson.D{bson.E{Key: "tags", Value: bson.D{bson.E{Key: "$lt", Value: "B"}}}}
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if !IsValidLt(expected, obtained) {
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}

func TestLte(t *testing.T) {
	var filter = mongo.NewFilter()
	filter.Add("gato", mongo.Lte("verde"))
	filter.Add("bola", mongo.Lte("azul"))

	expected := filter.Values()
	obtained := bson.D{bson.E{Key: "gato", Value: bson.D{bson.E{Key: "$lte", Value: "verde"}}}, bson.E{Key: "bola", Value: bson.D{bson.E{Key: "$lte", Value: "azul"}}}}
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if !IsValidLte(expected, obtained) {
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}
func TestLteField(t *testing.T) {
	expected := mongo.LteField("tags", "B")
	obtained := bson.D{bson.E{Key: "tags", Value: bson.D{bson.E{Key: "$lte", Value: "B"}}}}
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if !IsValidLte(expected, obtained) {
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}

func TestNe(t *testing.T) {
	var filter = mongo.NewFilter()
	filter.Add("gato", mongo.Ne("verde"))
	filter.Add("bola", mongo.Ne("azul"))

	expected := filter.Values()
	obtained := bson.D{bson.E{Key: "gato", Value: bson.D{bson.E{Key: "$ne", Value: "verde"}}}, bson.E{Key: "bola", Value: bson.D{bson.E{Key: "$ne", Value: "azul"}}}}
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if !IsValidNe(expected, obtained) {
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}
func TestNeField(t *testing.T) {
	expected := mongo.NeField("tags", "B")
	obtained := bson.D{bson.E{Key: "tags", Value: bson.D{bson.E{Key: "$ne", Value: "B"}}}}
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if !IsValidNe(expected, obtained) {
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}

func TestIn(t *testing.T) {
	var filter = mongo.NewFilter()
	filter.Add("qty", mongo.In([]any{20, 30, 41}))
	filter.Add("type", mongo.In([]any{"casa", "azul", "verde"}))

	expected := filter.Values()
	obtained := bson.D{bson.E{Key: "qty", Value: bson.D{bson.E{Key: "$in", Value: []any{20, 30, 41}}}}, bson.E{Key: "type", Value: bson.D{bson.E{Key: "$in", Value: []any{"casa", "azul", "verde"}}}}}
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if !IsValidIn(expected, obtained) {
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}
func TestInField(t *testing.T) {
	expected := mongo.InField("tags", []any{20, 30, 41})
	obtained := bson.D{bson.E{Key: "tags", Value: bson.D{bson.E{Key: "$in", Value: []any{20, 30, 41}}}}}
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if !IsValidIn(expected, obtained) {
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}

func TestNin(t *testing.T) {
	var filter = mongo.NewFilter()
	filter.Add("qty", mongo.Nin([]any{20, 30, 41}))
	filter.Add("type", mongo.Nin([]any{"casa", "azul", "verde"}))

	expected := filter.Values()
	obtained := bson.D{bson.E{Key: "qty", Value: bson.D{bson.E{Key: "$nin", Value: []any{20, 30, 41}}}}, bson.E{Key: "type", Value: bson.D{bson.E{Key: "$nin", Value: []any{"casa", "azul", "verde"}}}}}
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if !IsValidNin(expected, obtained) {
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}
func TestNinField(t *testing.T) {
	expected := mongo.NinField("tags", []any{20, 30, 41})
	obtained := bson.D{bson.E{Key: "tags", Value: bson.D{bson.E{Key: "$nin", Value: []any{20, 30, 41}}}}}
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if !IsValidIn(expected, obtained) {
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}

func TestExists(t *testing.T) {
	var filter = mongo.NewFilter()
	filter.Add("gato", mongo.Exists(false))
	filter.Add("bola", mongo.Exists(true))

	expected := filter.Values()
	obtained := bson.D{bson.E{Key: "gato", Value: bson.D{bson.E{Key: "$exists", Value: false}}}, bson.E{Key: "bola", Value: bson.D{bson.E{Key: "$exists", Value: true}}}}
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if !IsValidExists(expected, obtained) {
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}
func TestExistsField(t *testing.T) {
	expected := mongo.ExistsField("tags", false)
	obtained := bson.D{bson.E{Key: "tags", Value: bson.D{bson.E{Key: "$exists", Value: false}}}}
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if !IsValidExists(expected, obtained) {
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}

func TestType(t *testing.T) {
	var filter = mongo.NewFilter()
	filter.Add("gato", mongo.Type("verde"))
	filter.Add("bola", mongo.Type("azul"))

	expected := filter.Values()
	obtained := bson.D{bson.E{Key: "gato", Value: bson.D{bson.E{Key: "$type", Value: "verde"}}}, bson.E{Key: "bola", Value: bson.D{bson.E{Key: "$type", Value: "azul"}}}}
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if !IsValidType(expected, obtained) {
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}
func TestTypeField(t *testing.T) {
	expected := mongo.TypeField("type", "car")
	obtained := bson.D{bson.E{Key: "type", Value: bson.D{bson.E{Key: "$type", Value: "car"}}}}
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if !IsValidType(expected, obtained) {
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}

func TestDivide(t *testing.T) {
	var filter = mongo.NewFilter()
	filter.Add("qty", mongo.Divide([]any{20, 30, 41}))
	filter.Add("type", mongo.Divide([]any{"casa", "azul", "verde"}))

	expected := filter.Values()
	obtained := bson.D{bson.E{Key: "qty", Value: bson.D{bson.E{Key: "$divide", Value: []any{20, 30, 41}}}}, bson.E{Key: "type", Value: bson.D{bson.E{Key: "$divide", Value: []any{"casa", "azul", "verde"}}}}}
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if !IsValidDivide(expected, obtained) {
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}
func TestDivideField(t *testing.T) {
	expected := mongo.DivideField("tags", []any{20, 30, 41})
	obtained := bson.D{bson.E{Key: "tags", Value: bson.D{bson.E{Key: "$divide", Value: []any{20, 30, 41}}}}}
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if !IsValidDivide(expected, obtained) {
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}

func TestExpr(t *testing.T) {
	var filter = mongo.NewFilter()
	filter.Add("qty", mongo.Expr([]any{20, 30, 41}))
	filter.Add("type", mongo.Expr([]any{"casa", "azul", "verde"}))

	expected := filter.Values()
	obtained := bson.D{bson.E{Key: "qty", Value: bson.D{bson.E{Key: "$expr", Value: []any{20, 30, 41}}}}, bson.E{Key: "type", Value: bson.D{bson.E{Key: "$expr", Value: []any{"casa", "azul", "verde"}}}}}
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if !IsValidExpr(expected, obtained) {
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}

/*func TestExprField(t *testing.T) {
	expected := ExprField("tags", []any{20, 30, 41})
	obtained := bson.D{bson.E{Key: "$expr", Value: bson.E{Key: "tags", Value: []any{20, 30, 41}}}}
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if !IsValidExpr(expected, obtained) {
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}*/

func TestNot(t *testing.T) {
	var filter = mongo.NewFilter()
	filter.Add("item", mongo.Not("/^p.*/"))
	filter.Add("bola", mongo.Not(mongo.Regex("/^p.*/")))

	expected := filter.Values()
	obtained := bson.D{bson.E{Key: "item", Value: bson.D{bson.E{Key: "$not", Value: "/^p.*/"}}}, bson.E{Key: "bola", Value: bson.D{bson.E{Key: "$not", Value: mongo.Regex("/^p.*/")}}}}
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if !IsValidNot(expected, obtained) {
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}
func TestNotField(t *testing.T) {
	expected := mongo.NotField("item", mongo.Gt(1.99))
	obtained := bson.D{bson.E{Key: "item", Value: bson.D{bson.E{Key: "$not", Value: mongo.Gt(1.99)}}}}
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if !IsValidNot(expected, obtained) {
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}
