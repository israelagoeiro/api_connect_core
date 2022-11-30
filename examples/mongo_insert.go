package examples

import (
	"fmt"
	"github.com/israelagoeiro/api_connect_core/db"
	"time"
)

func MongoInsert() {
	start := time.Now()
	dataModel := FdibModel{
		IdPeca:     "78945612",
		ColetaRede: time.Now(),
		Etiqueta:   "IP08526A",
		Nserlum:    654321,
		Status:     true,
		Tempo: TempoModel{
			LigadoParcial:   time.Now(),
			LigadoDecorrido: time.Now(),
		},
	}

	input := db.NewMongoInputInsert()
	input.Model(dataModel)
	input.Data("nserlum", 123456)
	input.DataMap(map[string]any{
		"casa":      1,
		"gato":      "verde",
		"Opa":       1.5,
		"bola.azul": "circular",
		"bolax": map[string]any{
			"azulx": "quadrado",
		},
		"list": []any{"A", "B", "C", 1, 2, 3},
	})

	findParams := db.MongoFindParams{
		Collection: "users",
		Connection: "123456abc",
		Database:   "api-kdl-test",
		Fields:     []string{"idPeca", "coletaRede", "etiqueta", "status", "tempo", "bolax", "bola.azul"},
	}

	insertParams := db.MongoInsertParams{
		Collection: "users",
		Connection: "123456abc",
		Database:   "api-kdl-test",
		Input:      input,
		FindParams: findParams,
		DataLog: db.DataLog{
			Action:       "INSERT_NSERLUM",
			SaveHistory:  true,
			SaveInfo:     false,
			SaveAnalytic: false,
		},
	}
	dataResult := db.InsertOne(insertParams)

	model := FdibModel{}
	dataResult.Model(&model)
	dataResult.Print()
	//insertResult.toAPI()

	fmt.Println("id--->>>", dataResult.Id, time.Since(start))
	fmt.Println("insertResult--->>>", model, time.Since(start))
	fmt.Println("insertResult--->>>", model.Etiqueta, time.Since(start))
}
