package db

import (
	"fmt"
	"github.com/israelagoeiro/api_connect_core/test"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestAddToSet(t *testing.T) {
	var input = NewMongoInputUpdate()
	input.AddToSet("qty", 21)
	input.AddToSet("total", 100)
	input.AddToSet("casa", "azul")

	got := input.Values()
	want := bson.D{bson.E{Key: "$addToSet", Value: bson.D{bson.E{Key: "qty", Value: 21}, bson.E{Key: "casa", Value: "azul"}, bson.E{Key: "total", Value: 100}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !test.IsValidAddToSet(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestAddToSetMap(t *testing.T) {
	var input = NewMongoInputUpdate()
	input.AddToSetMap(map[string]any{
		"qty":   20,
		"total": 100,
	})

	got := input.Values()
	want := bson.D{bson.E{Key: "$addToSet", Value: bson.D{bson.E{Key: "qty", Value: 20}, bson.E{Key: "total", Value: 100}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !test.IsValidAddToSet(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestInc(t *testing.T) {
	var input = NewMongoInputUpdate()
	input.Inc("total", 100)
	input.Inc("qty", 20)

	got := input.Values()
	want := bson.D{bson.E{Key: "$inc", Value: bson.D{bson.E{Key: "qty", Value: 20}, bson.E{Key: "total", Value: 100}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !test.IsValidInc(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestIncMap(t *testing.T) {
	var input = NewMongoInputUpdate()
	input.IncMap(map[string]any{
		"qty":   20,
		"total": 100,
	})

	got := input.Values()
	want := bson.D{bson.E{Key: "$inc", Value: bson.D{bson.E{Key: "qty", Value: 20}, bson.E{Key: "total", Value: 100}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !test.IsValidInc(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestSet(t *testing.T) {
	var input = NewMongoInputUpdate()
	input.Set("total", 100)
	input.Set("qty", 20)

	got := input.Values()
	want := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "qty", Value: 20}, bson.E{Key: "total", Value: 100}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !test.IsValidSet(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestSetMap(t *testing.T) {
	var input = NewMongoInputUpdate()
	input.SetMap(map[string]any{
		"qty":   20,
		"total": 100,
	})

	got := input.Values()
	want := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "qty", Value: 20}, bson.E{Key: "total", Value: 100}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !test.IsValidSet(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestUpdateIsValid(t *testing.T) {
	var input = NewMongoInputUpdate()
	input.AddToSet("qty", 20)

	got := input.IsValid()
	fmt.Println("got", got)
	fmt.Println("want", true)

	if got != true {
		t.Errorf("got %t, wanted %t", got, true)
	}
}
