package mongo_test

import (
	"fmt"
	"github.com/israelagoeiro/api_connect_core/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestNewClient(t *testing.T) {
	var filter = mongo.NewFilter()
	filter.Add("gato", mongo.Eq("verde"))
	filter.Add("bola", mongo.Eq("azul"))

	expected := filter.Values()
	obtained := bson.D{bson.E{Key: "gato", Value: bson.D{bson.E{Key: "$eq", Value: "verde"}}}, bson.E{Key: "bola", Value: bson.D{bson.E{Key: "$eq", Value: "azul"}}}}
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if !mongo.IsValidEq(expected, obtained) {
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}
