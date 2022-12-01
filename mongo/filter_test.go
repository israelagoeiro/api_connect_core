package mongo

import (
	"fmt"
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
	var filter = NewFilter()
	filter.Add("gato", Eq("verde"))
	filter.Add("bola", Eq("azul"))

	got := filter.Values()
	want := bson.D{bson.E{Key: "gato", Value: bson.D{bson.E{Key: "$eq", Value: "verde"}}}, bson.E{Key: "bola", Value: bson.D{bson.E{Key: "$eq", Value: "azul"}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !IsValidEq(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestEqField(t *testing.T) {
	got := EqField("tags", "B")
	want := bson.D{bson.E{Key: "tags", Value: bson.D{bson.E{Key: "$eq", Value: "B"}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !IsValidEq(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestGt(t *testing.T) {
	var filter = NewFilter()
	filter.Add("gato", Gt("verde"))
	filter.Add("bola", Gt("azul"))

	got := filter.Values()
	want := bson.D{bson.E{Key: "gato", Value: bson.D{bson.E{Key: "$gt", Value: "verde"}}}, bson.E{Key: "bola", Value: bson.D{bson.E{Key: "$gt", Value: "azul"}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !IsValidGt(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestGtField(t *testing.T) {
	got := GtField("tags", "B")
	want := bson.D{bson.E{Key: "tags", Value: bson.D{bson.E{Key: "$gt", Value: "B"}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !IsValidGt(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestGte(t *testing.T) {
	var filter = NewFilter()
	filter.Add("gato", Gte("verde"))
	filter.Add("bola", Gte("azul"))

	got := filter.Values()
	want := bson.D{bson.E{Key: "gato", Value: bson.D{bson.E{Key: "$gte", Value: "verde"}}}, bson.E{Key: "bola", Value: bson.D{bson.E{Key: "$gte", Value: "azul"}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !IsValidGte(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestGteField(t *testing.T) {
	got := GteField("tags", "B")
	want := bson.D{bson.E{Key: "tags", Value: bson.D{bson.E{Key: "$gte", Value: "B"}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !IsValidGte(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestLt(t *testing.T) {
	var filter = NewFilter()
	filter.Add("gato", Lt("verde"))
	filter.Add("bola", Lt("azul"))

	got := filter.Values()
	want := bson.D{bson.E{Key: "gato", Value: bson.D{bson.E{Key: "$lt", Value: "verde"}}}, bson.E{Key: "bola", Value: bson.D{bson.E{Key: "$lt", Value: "azul"}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !IsValidLt(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestLtField(t *testing.T) {
	got := LtField("tags", "B")
	want := bson.D{bson.E{Key: "tags", Value: bson.D{bson.E{Key: "$lt", Value: "B"}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !IsValidLt(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestLte(t *testing.T) {
	var filter = NewFilter()
	filter.Add("gato", Lte("verde"))
	filter.Add("bola", Lte("azul"))

	got := filter.Values()
	want := bson.D{bson.E{Key: "gato", Value: bson.D{bson.E{Key: "$lte", Value: "verde"}}}, bson.E{Key: "bola", Value: bson.D{bson.E{Key: "$lte", Value: "azul"}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !IsValidLte(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestLteField(t *testing.T) {
	got := LteField("tags", "B")
	want := bson.D{bson.E{Key: "tags", Value: bson.D{bson.E{Key: "$lte", Value: "B"}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !IsValidLte(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNe(t *testing.T) {
	var filter = NewFilter()
	filter.Add("gato", Ne("verde"))
	filter.Add("bola", Ne("azul"))

	got := filter.Values()
	want := bson.D{bson.E{Key: "gato", Value: bson.D{bson.E{Key: "$ne", Value: "verde"}}}, bson.E{Key: "bola", Value: bson.D{bson.E{Key: "$ne", Value: "azul"}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !IsValidNe(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestNeField(t *testing.T) {
	got := NeField("tags", "B")
	want := bson.D{bson.E{Key: "tags", Value: bson.D{bson.E{Key: "$ne", Value: "B"}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !IsValidNe(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestIn(t *testing.T) {
	var filter = NewFilter()
	filter.Add("qty", In([]any{20, 30, 41}))
	filter.Add("type", In([]any{"casa", "azul", "verde"}))

	got := filter.Values()
	want := bson.D{bson.E{Key: "qty", Value: bson.D{bson.E{Key: "$in", Value: []any{20, 30, 41}}}}, bson.E{Key: "type", Value: bson.D{bson.E{Key: "$in", Value: []any{"casa", "azul", "verde"}}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !IsValidIn(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestInField(t *testing.T) {
	got := InField("tags", []any{20, 30, 41})
	want := bson.D{bson.E{Key: "tags", Value: bson.D{bson.E{Key: "$in", Value: []any{20, 30, 41}}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !IsValidIn(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNin(t *testing.T) {
	var filter = NewFilter()
	filter.Add("qty", Nin([]any{20, 30, 41}))
	filter.Add("type", Nin([]any{"casa", "azul", "verde"}))

	got := filter.Values()
	want := bson.D{bson.E{Key: "qty", Value: bson.D{bson.E{Key: "$nin", Value: []any{20, 30, 41}}}}, bson.E{Key: "type", Value: bson.D{bson.E{Key: "$nin", Value: []any{"casa", "azul", "verde"}}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !IsValidNin(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestNinField(t *testing.T) {
	got := NinField("tags", []any{20, 30, 41})
	want := bson.D{bson.E{Key: "tags", Value: bson.D{bson.E{Key: "$nin", Value: []any{20, 30, 41}}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !IsValidIn(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestExists(t *testing.T) {
	var filter = NewFilter()
	filter.Add("gato", Exists(false))
	filter.Add("bola", Exists(true))

	got := filter.Values()
	want := bson.D{bson.E{Key: "gato", Value: bson.D{bson.E{Key: "$exists", Value: false}}}, bson.E{Key: "bola", Value: bson.D{bson.E{Key: "$exists", Value: true}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !IsValidExists(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestExistsField(t *testing.T) {
	got := ExistsField("tags", false)
	want := bson.D{bson.E{Key: "tags", Value: bson.D{bson.E{Key: "$exists", Value: false}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !IsValidExists(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestType(t *testing.T) {
	var filter = NewFilter()
	filter.Add("gato", Type("verde"))
	filter.Add("bola", Type("azul"))

	got := filter.Values()
	want := bson.D{bson.E{Key: "gato", Value: bson.D{bson.E{Key: "$type", Value: "verde"}}}, bson.E{Key: "bola", Value: bson.D{bson.E{Key: "$type", Value: "azul"}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !IsValidType(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestTypeField(t *testing.T) {
	got := TypeField("type", "car")
	want := bson.D{bson.E{Key: "type", Value: bson.D{bson.E{Key: "$type", Value: "car"}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !IsValidType(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDivide(t *testing.T) {
	var filter = NewFilter()
	filter.Add("qty", Divide([]any{20, 30, 41}))
	filter.Add("type", Divide([]any{"casa", "azul", "verde"}))

	got := filter.Values()
	want := bson.D{bson.E{Key: "qty", Value: bson.D{bson.E{Key: "$divide", Value: []any{20, 30, 41}}}}, bson.E{Key: "type", Value: bson.D{bson.E{Key: "$divide", Value: []any{"casa", "azul", "verde"}}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !IsValidDivide(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestDivideField(t *testing.T) {
	got := DivideField("tags", []any{20, 30, 41})
	want := bson.D{bson.E{Key: "tags", Value: bson.D{bson.E{Key: "$divide", Value: []any{20, 30, 41}}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !IsValidDivide(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestExpr(t *testing.T) {
	var filter = NewFilter()
	filter.Add("qty", Expr([]any{20, 30, 41}))
	filter.Add("type", Expr([]any{"casa", "azul", "verde"}))

	got := filter.Values()
	want := bson.D{bson.E{Key: "qty", Value: bson.D{bson.E{Key: "$expr", Value: []any{20, 30, 41}}}}, bson.E{Key: "type", Value: bson.D{bson.E{Key: "$expr", Value: []any{"casa", "azul", "verde"}}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !IsValidExpr(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

/*func TestExprField(t *testing.T) {
	got := ExprField("tags", []any{20, 30, 41})
	want := bson.D{bson.E{Key: "$expr", Value: bson.E{Key: "tags", Value: []any{20, 30, 41}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !IsValidExpr(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}*/

func TestNot(t *testing.T) {
	var filter = NewFilter()
	filter.Add("item", Not("/^p.*/"))
	filter.Add("bola", Not(Regex("/^p.*/")))

	got := filter.Values()
	want := bson.D{bson.E{Key: "item", Value: bson.D{bson.E{Key: "$not", Value: "/^p.*/"}}}, bson.E{Key: "bola", Value: bson.D{bson.E{Key: "$not", Value: Regex("/^p.*/")}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !IsValidNot(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestNotField(t *testing.T) {
	got := NotField("item", Gt(1.99))
	want := bson.D{bson.E{Key: "item", Value: bson.D{bson.E{Key: "$not", Value: Gt(1.99)}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !IsValidNot(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
