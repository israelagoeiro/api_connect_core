package mongo_test

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
	"testing"
)

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

func TestAddToSet(t *testing.T) {
	var input = NewUpdateInput()
	input.AddToSet("qty", 21)
	input.AddToSet("total", 100)
	input.AddToSet("casa", "azul")

	expected := input.Values()
	obtained := bson.D{bson.E{Key: "$addToSet", Value: bson.D{bson.E{Key: "qty", Value: 21}, bson.E{Key: "casa", Value: "azul"}, bson.E{Key: "total", Value: 100}}}}
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if !IsValidAddToSet(expected, obtained) {
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}

func TestAddToSetMap(t *testing.T) {
	var input = NewUpdateInput()
	input.AddToSetMap(map[string]any{
		"qty":   20,
		"total": 100,
	})

	expected := input.Values()
	obtained := bson.D{bson.E{Key: "$addToSet", Value: bson.D{bson.E{Key: "qty", Value: 20}, bson.E{Key: "total", Value: 100}}}}
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if !IsValidAddToSet(expected, obtained) {
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}

func TestInc(t *testing.T) {
	var input = NewUpdateInput()
	input.Inc("total", 100)
	input.Inc("qty", 20)

	expected := input.Values()
	obtained := bson.D{bson.E{Key: "$inc", Value: bson.D{bson.E{Key: "qty", Value: 20}, bson.E{Key: "total", Value: 100}}}}
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if !IsValidInc(expected, obtained) {
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}

func TestIncMap(t *testing.T) {
	var input = NewUpdateInput()
	input.IncMap(map[string]any{
		"qty":   20,
		"total": 100,
	})

	expected := input.Values()
	obtained := bson.D{bson.E{Key: "$inc", Value: bson.D{bson.E{Key: "qty", Value: 20}, bson.E{Key: "total", Value: 100}}}}
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if !IsValidInc(expected, obtained) {
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}

func TestSet(t *testing.T) {
	var input = NewUpdateInput()
	input.Set("total", 100)
	input.Set("qty", 20)

	expected := input.Values()
	obtained := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "qty", Value: 20}, bson.E{Key: "total", Value: 100}}}}
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if !IsValidSet(expected, obtained) {
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}

func TestSetMap(t *testing.T) {
	var input = NewUpdateInput()
	input.SetMap(map[string]any{
		"qty":   20,
		"total": 100,
	})

	expected := input.Values()
	obtained := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "qty", Value: 20}, bson.E{Key: "total", Value: 100}}}}
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if !IsValidSet(expected, obtained) {
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}

func TestUpdateIsValid(t *testing.T) {
	var input = NewUpdateInput()
	input.AddToSet("qty", 20)

	expected := input.IsValid()
	fmt.Println("expected", expected)
	fmt.Println("obtained", true)

	if got != true {
		t.Errorf("expected %t, obtained %t", expected, true)
	}
}
