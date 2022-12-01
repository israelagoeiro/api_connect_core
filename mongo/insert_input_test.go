package mongo

import (
	"fmt"
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
	var input = NewInsertInput()
	input.Data("qty", 20)

	got := input.Values()
	want := bson.D{bson.E{Key: "qty", Value: 20}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !IsValidData(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestInsertDataMap(t *testing.T) {
	var input = NewInsertInput()
	input.DataMap(map[string]any{
		"qty":   20,
		"total": 100,
		"color": "red",
	})

	got := input.Values()
	want := bson.D{bson.E{Key: "qty", Value: 20}, bson.E{Key: "total", Value: 100}, bson.E{Key: "color", Value: "red"}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !IsValidData(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestInsertIsValid(t *testing.T) {
	var input = NewInsertInput()
	input.Data("qty", 20)
	got := input.IsValid()
	fmt.Println("got", got)
	fmt.Println("want", true)

	if got != true {
		t.Errorf("got %t, wanted %t", got, true)
	}
}
