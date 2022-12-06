package mongo_test

import (
	"fmt"
	"github.com/israelagoeiro/api_connect_core/mongo"

	"time"
)

func ExampleInsertOne() {
	start := time.Now()
	dataModel := mongo.FdibModel{
		IdPeca:     "78945612",
		ColetaRede: time.Now(),
		Etiqueta:   "IP08526A",
		Nserlum:    654321,
		Status:     true,
		Tempo: mongo.TempoModel{
			LigadoParcial:   time.Now(),
			LigadoDecorrido: time.Now(),
		},
	}

	input := mongo.NewInput()
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

	findParams := mongo.FindParams{
		Collection: "users",
		Connection: "123456abc",
		Database:   "api-kdl-test",
		Fields:     []string{"idPeca", "coletaRede", "etiqueta", "status", "tempo", "bolax", "bola.azul"},
	}

	insertParams := mongo.InsertParams{
		Collection: "users",
		Connection: "123456abc",
		Database:   "api-kdl-test",
		Input:      input,
		FindParams: findParams,
		DataLog: mongo.DataLog{
			Action:       "INSERT_NSERLUM",
			SaveHistory:  true,
			SaveInfo:     false,
			SaveAnalytic: false,
		},
	}
	dataResult := mongo.InsertOne(insertParams)

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

func ExampleInsertMany() {
	start := time.Now()
	dataModel := mongo.FdibModel{
		IdPeca:     "78945612",
		ColetaRede: time.Now(),
		Etiqueta:   "IP08526A",
		Nserlum:    654321,
		Status:     true,
		Tempo: mongo.TempoModel{
			LigadoParcial:   time.Now(),
			LigadoDecorrido: time.Now(),
		},
	}

	input := mongo.NewInput()
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

	findParams := mongo.FindParams{
		Collection: "users",
		Connection: "123456abc",
		Database:   "api-kdl-test",
		Fields:     []string{"idPeca", "coletaRede", "etiqueta", "status", "tempo", "bolax", "bola.azul"},
	}

	insertParams := mongo.InsertParams{
		Collection: "users",
		Connection: "123456abc",
		Database:   "api-kdl-test",
		Input:      input,
		FindParams: findParams,
		DataLog: mongo.DataLog{
			Action:       "INSERT_NSERLUM",
			SaveHistory:  true,
			SaveInfo:     false,
			SaveAnalytic: false,
		},
	}
	dataResult := mongo.InsertMany(insertParams)

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
