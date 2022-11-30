package examples

import (
	"fmt"
	"github.com/israelagoeiro/api_connect_core/db"
	"time"
)

func MongoFind() {
	start := time.Now()

	filter := db.NewMongoFilter()
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

	findParams := db.MongoFindParams{
		Collection: "users",
		Connection: "123456abc",
		Database:   "api-kdl-test",
		Filter:     filter,
		Fields:     []string{"idPeca", "coletaRede", "etiqueta", "status", "tempo", "gato", "nserlum", "Opa"},
		Options: db.FindOptions{
			Sort: db.Sort("etiqueta", false),
		},
	}
	dataResult := db.Find(findParams)

	models := []FdibModel{}
	dataResult.Print()
	dataResult.Models(&models)

	//dataResult.toAPI()
	fmt.Println("dataResult--->>>", models, time.Since(start))
	//fmt.Println("dataResult--->>>", model, time.Since(start))
	//fmt.Println("dataResult--->>>", models[1].Etiqueta, time.Since(start))
}