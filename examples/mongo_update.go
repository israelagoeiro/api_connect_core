package examples

import (
	"fmt"
	"github.com/israelagoeiro/api_connect_core/db"
	"time"
)

func MongoUpdate() {
	start := time.Now()

	filter := db.NewMongoFilter()
	filter.Id("6384f0e452ed0e02aa02d688")

	input := db.NewMongoInputUpdate()
	input.Set("nserlum", 999999)
	input.SetMap(map[string]any{
		"bola": "azul",
		"mala": "verde",
	})

	findParams := db.MongoFindParams{
		Collection: "users",
		Connection: "123456abc",
		Database:   "api-kdl-test",
		Fields:     []string{"idPeca", "coletaRede", "etiqueta", "status", "tempo"},
	}

	updateParams := db.MongoUpdateParams{
		Collection: "users",
		Connection: "123456abc",
		Database:   "api-kdl-test",
		Info:       nil,
		Input:      input,
		Filter:     filter,
		FindParams: findParams,
		UpdateOptions: db.UpdateOptions{
			ReturnOriginal: true,
			Upsert:         true,
		},
		DataLog: db.DataLog{
			Action:       "UPDATE_NSERLUM",
			SaveChange:   true,
			SaveHistory:  true,
			SaveInfo:     false,
			SaveAnalytic: false,
		},
	}
	dataResult := db.UpdateOne(updateParams)

	model := FdibModel{}
	dataResult.Model(&model)
	dataResult.Print()
	//insertResult.toAPI()

	fmt.Println("id--->>>", dataResult.Id, time.Since(start))
	fmt.Println("insertResult--->>>", model, time.Since(start))
	fmt.Println("insertResult--->>>", model.Etiqueta, time.Since(start))

}
