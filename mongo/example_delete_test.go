package mongo

import (
	"fmt"
	"github.com/israelagoeiro/api_connect_core/examples"
	"time"
)

func ExampleDeleteMany() {
	start := time.Now()

	filter := NewFilter()
	filter.Add("quantity", Gte(20))

	//db.inventory.find( { quantity: { $gte: 20 } } )
	//db.inventory.find(bson.D{bson.E{Key: "quantity", Value: bson.E{Key: "$gte", Value: 20}}})

	//filter.Id("6384f0e452ed0e02aa02d688")
	//filter.Add("nserlum", db.Eq(654321))
	//filter.Add("status", db.Eq(true))
	//filter.Add("nserlum", db.Gt(500000))
	//filter.Add("nserlum", db.Gte(500000))
	//filter.Add("nserlum", db.Lt(500000))
	//filter.Add("nserlum", db.Lte(500000))
	//filter.Add("nserlum", db.Ne(123456))
	//filter.Add("nserlum", db.In([]any{654321}))
	//filter.Add("nserlum", db.Nin([]any{654321}))
	//filter.Add("nserlum", db.Exists(true))
	///???filter.Add("gato", db.Type("amarelo"))
	//filter.Add("gato", db.Regex("v.rd+"))

	findParams := FindParams{
		Collection: "users",
		Connection: "123456abc",
		Database:   "api-kdl-test",
		Filter:     filter,
		Fields:     []string{"idPeca", "coletaRede", "etiqueta", "status", "tempo", "gato", "nserlum", "Opa"},
		Options: FindOptions{
			Sort: Sort("etiqueta", false),
		},
	}
	dataResult := Find(findParams)

	var models []examples.FdibModel
	dataResult.Print()
	dataResult.Models(&models)

	//dataResult.toAPI()
	fmt.Println("dataResult--->>>", models, time.Since(start))
	//fmt.Println("dataResult--->>>", model, time.Since(start))
	//fmt.Println("dataResult--->>>", models[1].Etiqueta, time.Since(start))

	// Output:
	// 1257894000000
	// 2009-11-10 23:00:00 +0000 UTC
}

func ExampleDeleteOne() {
	start := time.Now()

	filter := NewFilter()
	filter.Add("quantity", Gte(20))

	//db.inventory.find( { quantity: { $gte: 20 } } )
	//db.inventory.find(bson.D{bson.E{Key: "quantity", Value: bson.E{Key: "$gte", Value: 20}}})

	//filter.Id("6384f0e452ed0e02aa02d688")
	//filter.Add("nserlum", db.Eq(654321))
	//filter.Add("status", db.Eq(true))
	//filter.Add("nserlum", db.Gt(500000))
	//filter.Add("nserlum", db.Gte(500000))
	//filter.Add("nserlum", db.Lt(500000))
	//filter.Add("nserlum", db.Lte(500000))
	//filter.Add("nserlum", db.Ne(123456))
	//filter.Add("nserlum", db.In([]any{654321}))
	//filter.Add("nserlum", db.Nin([]any{654321}))
	//filter.Add("nserlum", db.Exists(true))
	///???filter.Add("gato", db.Type("amarelo"))
	//filter.Add("gato", db.Regex("v.rd+"))

	findParams := FindParams{
		Collection: "users",
		Connection: "123456abc",
		Database:   "api-kdl-test",
		Filter:     filter,
		Fields:     []string{"idPeca", "coletaRede", "etiqueta", "status", "tempo", "gato", "nserlum", "Opa"},
		Options: FindOptions{
			Sort: Sort("etiqueta", false),
		},
	}
	dataResult := Find(findParams)

	var models []examples.FdibModel
	dataResult.Print()
	dataResult.Models(&models)

	//dataResult.toAPI()
	fmt.Println("dataResult--->>>", models, time.Since(start))
	//fmt.Println("dataResult--->>>", model, time.Since(start))
	//fmt.Println("dataResult--->>>", models[1].Etiqueta, time.Since(start))

	// Output:
	// 1257894000000
	// 2009-11-10 23:00:00 +0000 UTC
}
