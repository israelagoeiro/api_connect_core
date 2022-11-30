package examples

import (
	"fmt"
	"github.com/israelagoeiro/api_connect_core/db"
	"time"
)

func MongoDeleteOne() {
	start := time.Now()

	findParams := db.MongoFindParams{
		Collection: "users",
		Connection: "123456abc",
		Database:   "api-kdl-test",
		Fields:     []string{"idPeca", "coletaRede", "etiqueta", "status", "tempo", "bolax", "bola.azul"},
	}

	deleteParams := db.MongoDeleteParams{
		Collection: "users",
		Connection: "123456abc",
		Database:   "api-kdl-test",
		FindParams: findParams,
		DataLog: db.DataLog{
			Action:       "INSERT_NSERLUM",
			SaveHistory:  true,
			SaveInfo:     false,
			SaveAnalytic: false,
		},
	}
	dataResult := db.DeleteOne(deleteParams)

	model := FdibModel{}
	dataResult.Model(&model)
	dataResult.Print()
	//insertResult.toAPI()

	fmt.Println("id--->>>", dataResult.Id, time.Since(start))
	fmt.Println("insertResult--->>>", model, time.Since(start))
	fmt.Println("insertResult--->>>", model.Etiqueta, time.Since(start))
}

func MongoDeleteMany() {
	start := time.Now()

	findParams := db.MongoFindParams{
		Collection: "users",
		Connection: "123456abc",
		Database:   "api-kdl-test",
		Fields:     []string{"idPeca", "coletaRede", "etiqueta", "status", "tempo", "bolax", "bola.azul"},
	}

	deleteParams := db.MongoDeleteParams{
		Collection: "users",
		Connection: "123456abc",
		Database:   "api-kdl-test",
		FindParams: findParams,
		DataLog: db.DataLog{
			Action:       "INSERT_NSERLUM",
			SaveHistory:  true,
			SaveInfo:     false,
			SaveAnalytic: false,
		},
	}
	dataResult := db.DeleteMany(deleteParams)

	model := FdibModel{}
	dataResult.Model(&model)
	dataResult.Print()
	//insertResult.toAPI()

	fmt.Println("id--->>>", dataResult.Id, time.Since(start))
	fmt.Println("insertResult--->>>", model, time.Since(start))
	fmt.Println("insertResult--->>>", model.Etiqueta, time.Since(start))
}
