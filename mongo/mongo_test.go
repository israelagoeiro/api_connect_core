package mongo_test

import (
	"fmt"
	"github.com/israelagoeiro/api_connect_core/test"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestNewClient(t *testing.T) {
	var filter = NewFilter()
	filter.Add("gato", Eq("verde"))
	filter.Add("bola", Eq("azul"))

	got := filter.Values()
	want := bson.D{bson.E{Key: "gato", Value: bson.D{bson.E{Key: "$eq", Value: "verde"}}}, bson.E{Key: "bola", Value: bson.D{bson.E{Key: "$eq", Value: "azul"}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !test.IsValidEq(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
