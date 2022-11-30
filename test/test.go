package test

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
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

func IsValidData(a, b bson.D) bool {
	if len(a) != len(b) {
		return false
	}

	listA := getMap(a)
	listB := getMap(b)
	return compareMap(listA, listB)
}

func IsValidAddToSet(a, b bson.D) bool {
	if len(a) != len(b) {
		return false
	}

	statusA := false
	var dataA interface{}
	for i := range a {
		if a[i].Key == "$addToSet" {
			statusA = true
			dataA = a[i].Value
			break
		}
	}

	statusB := false
	var dataB interface{}
	for i := range b {
		if b[i].Key == "$addToSet" {
			statusB = true
			dataB = b[i].Value
			break
		}
	}

	if !statusA || !statusB {
		fmt.Println("Error IsValidAddToSet: Não foi encontrato parâmetros do tipo AddToSet(?) ou AddToSetMap(?)")
		return false
	}

	if reflect.ValueOf(dataA).Len() != reflect.ValueOf(dataB).Len() {
		return false
	}
	if dataA != nil && dataB != nil {
		listA := getMap(dataA)
		listB := getMap(dataB)
		return compareMap(listA, listB)
	}
	return false
}

func IsValidInc(a, b bson.D) bool {
	if len(a) != len(b) {
		return false
	}

	statusA := false
	var dataA interface{}
	for i := range a {
		if a[i].Key == "$inc" {
			statusA = true
			dataA = a[i].Value
			break
		}
	}

	statusB := false
	var dataB interface{}
	for i := range b {
		if b[i].Key == "$inc" {
			statusB = true
			dataB = b[i].Value
			break
		}
	}

	if !statusA || !statusB {
		fmt.Println("Error IsValidInc: Não foi encontrato parâmetros do tipo Inc(?) ou IncMap(?)")
		return false
	}

	if reflect.ValueOf(dataA).Len() != reflect.ValueOf(dataB).Len() {
		return false
	}
	if dataA != nil && dataB != nil {
		listA := getMap(dataA)
		listB := getMap(dataB)
		return compareMap(listA, listB)
	}
	return false
}

func IsValidSet(a, b bson.D) bool {
	if len(a) != len(b) {
		return false
	}

	statusA := false
	var dataA interface{}
	for i := range a {
		if a[i].Key == "$set" {
			statusA = true
			dataA = a[i].Value
			break
		}
	}

	statusB := false
	var dataB interface{}
	for i := range b {
		if b[i].Key == "$set" {
			statusB = true
			dataB = b[i].Value
			break
		}
	}

	if !statusA || !statusB {
		fmt.Println("Error IsValidSet: Não foi encontrato parâmetros do tipo Set(?) ou SetMap(?)")
		return false
	}

	if reflect.ValueOf(dataA).Len() != reflect.ValueOf(dataB).Len() {
		return false
	}
	if dataA != nil && dataB != nil {
		listA := getMap(dataA)
		listB := getMap(dataB)
		return compareMap(listA, listB)
	}
	return false
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
