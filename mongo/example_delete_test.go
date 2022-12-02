package mongo_test

import (
	"fmt"
	"github.com/israelagoeiro/api_connect_core/mongo"
	"time"
)

func ExampleDeleteMany() {
	start := time.Now()

	findParams := mongo.FindParams{
		Collection: "users",
		Connection: "123456abc",
		Database:   "api-kdl-test",
		Fields:     []string{"idPeca", "coletaRede", "etiqueta", "status", "tempo", "bolax", "bola.azul"},
	}

	deleteParams := mongo.DeleteParams{
		Collection: "users",
		Connection: "123456abc",
		Database:   "api-kdl-test",
		FindParams: findParams,
		DataLog: mongo.DataLog{
			Action:       "INSERT_NSERLUM",
			SaveHistory:  true,
			SaveInfo:     false,
			SaveAnalytic: false,
		},
	}
	dataResult := mongo.DeleteMany(deleteParams)

	model := mongo.FdibModel{}
	dataResult.Model(&model)
	dataResult.Print()
	//insertResult.toAPI()

	fmt.Println("id--->>>", dataResult.Id, time.Since(start))
	fmt.Println("insertResult--->>>", model, time.Since(start))
	fmt.Println("insertResult--->>>", model.Etiqueta, time.Since(start))

	// Output:
	// 1257894000000
	// 2009-11-10 23:00:00 +0000 UTC
}

func ExampleDeleteOne() {
	start := time.Now()

	findParams := mongo.FindParams{
		Collection: "users",
		Connection: "123456abc",
		Database:   "api-kdl-test",
		Fields:     []string{"idPeca", "coletaRede", "etiqueta", "status", "tempo", "bolax", "bola.azul"},
	}

	deleteParams := mongo.DeleteParams{
		Collection: "users",
		Connection: "123456abc",
		Database:   "api-kdl-test",
		FindParams: findParams,
		DataLog: mongo.DataLog{
			Action:       "INSERT_NSERLUM",
			SaveHistory:  true,
			SaveInfo:     false,
			SaveAnalytic: false,
		},
	}
	dataResult := mongo.DeleteOne(deleteParams)

	model := mongo.FdibModel{}
	dataResult.Model(&model)
	dataResult.Print()
	//insertResult.toAPI()

	fmt.Println("id--->>>", dataResult.Id, time.Since(start))
	fmt.Println("insertResult--->>>", model, time.Since(start))
	fmt.Println("insertResult--->>>", model.Etiqueta, time.Since(start))

	// Output:
	// 1257894000000
	// 2009-11-10 23:00:00 +0000 UTC
}
