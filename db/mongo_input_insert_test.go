package db

import (
	"fmt"
	"github.com/israelagoeiro/api_connect_core/test"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestData(t *testing.T) {
	var input = NewMongoInputInsert()
	input.Data("qty", 20)

	got := input.Values()
	want := bson.D{bson.E{Key: "qty", Value: 20}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !test.IsValidData(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestInsertDataMap(t *testing.T) {
	var input = NewMongoInputInsert()
	input.DataMap(map[string]any{
		"qty":   20,
		"total": 100,
		"color": "red",
	})

	got := input.Values()
	want := bson.D{bson.E{Key: "qty", Value: 20}, bson.E{Key: "total", Value: 100}, bson.E{Key: "color", Value: "red"}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !test.IsValidData(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestInsertIsValid(t *testing.T) {
	var input = NewMongoInputInsert()
	input.Data("qty", 20)
	got := input.IsValid()
	fmt.Println("got", got)
	fmt.Println("want", true)

	if got != true {
		t.Errorf("got %t, wanted %t", got, true)
	}
}
