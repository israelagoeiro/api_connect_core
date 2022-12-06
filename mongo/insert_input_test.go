package mongo_test

import (
	"fmt"
	"github.com/israelagoeiro/api_connect_core/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func IsValidData(a, b bson.D) bool {
	if len(a) != len(b) {
		return false
	}

	listA := getMap(a)
	listB := getMap(b)
	return compareMap(listA, listB)
}

func TestData(t *testing.T) {
	var input = mongo.NewInput()
	input.Data("qty", 20)

	expected := input.Values()
	obtained := bson.D{bson.E{Key: "qty", Value: 20}}
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if !IsValidData(expected, obtained) {
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}

func TestInsertDataMap(t *testing.T) {
	var input = mongo.NewInput()
	input.DataMap(map[string]any{
		"qty":   20,
		"total": 100,
		"color": "red",
	})

	expected := input.Values()
	obtained := bson.D{bson.E{Key: "qty", Value: 20}, bson.E{Key: "total", Value: 100}, bson.E{Key: "color", Value: "red"}}
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if !IsValidData(expected, obtained) {
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}

func TestInsertIsValid(t *testing.T) {
	var input = mongo.NewInput()
	input.Data("qty", 20)

	obtained := input.IsValid()
	fmt.Println("expected", true)
	fmt.Println("obtained", obtained)

	if obtained != true {
		t.Errorf("obtained %t, expected %t", obtained, true)
	}
}
