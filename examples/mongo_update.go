package examples

import (
	"fmt"
	"github.com/israelagoeiro/api_connect_core/mongo"
	"time"
)

func MongoUpdateOne() {
	start := time.Now()

	filter := mongo.NewFilter()
	filter.Id("6384f0e452ed0e02aa02d688")

	input := mongo.NewUpdateInput()
	input.Set("nserlum", 999999)
	input.SetMap(map[string]any{
		"bola": "azul",
		"mala": "verde",
	})

	findParams := mongo.FindParams{
		Collection: "users",
		Connection: "123456abc",
		Database:   "api-kdl-test",
		Fields:     []string{"idPeca", "coletaRede", "etiqueta", "status", "tempo"},
	}

	updateParams := mongo.UpdateParams{
		Collection: "users",
		Connection: "123456abc",
		Database:   "api-kdl-test",
		Info:       nil,
		Input:      input,
		Filter:     filter,
		FindParams: findParams,
		UpdateOptions: mongo.UpdateOptions{
			ReturnOriginal: true,
			Upsert:         true,
		},
		DataLog: mongo.DataLog{
			Action:       "UPDATE_NSERLUM",
			SaveChange:   true,
			SaveHistory:  true,
			SaveInfo:     false,
			SaveAnalytic: false,
		},
	}
	dataResult := mongo.UpdateOne(updateParams)

	model := FdibModel{}
	dataResult.Model(&model)
	dataResult.Print()
	//insertResult.toAPI()

	fmt.Println("id--->>>", dataResult.Id, time.Since(start))
	fmt.Println("insertResult--->>>", model, time.Since(start))
	fmt.Println("insertResult--->>>", model.Etiqueta, time.Since(start))

}

func MongoUpdateMany() {
	start := time.Now()

	filter := mongo.NewFilter()
	filter.Id("6384f0e452ed0e02aa02d688")

	input := mongo.NewUpdateInput()
	input.Set("nserlum", 999999)
	input.SetMap(map[string]any{
		"bola": "azul",
		"mala": "verde",
	})

	findParams := mongo.FindParams{
		Collection: "users",
		Connection: "123456abc",
		Database:   "api-kdl-test",
		Fields:     []string{"idPeca", "coletaRede", "etiqueta", "status", "tempo"},
	}

	updateParams := mongo.UpdateParams{
		Collection: "users",
		Connection: "123456abc",
		Database:   "api-kdl-test",
		Info:       nil,
		Input:      input,
		Filter:     filter,
		FindParams: findParams,
		UpdateOptions: mongo.UpdateOptions{
			ReturnOriginal: true,
			Upsert:         true,
		},
		DataLog: mongo.DataLog{
			Action:       "UPDATE_NSERLUM",
			SaveChange:   true,
			SaveHistory:  true,
			SaveInfo:     false,
			SaveAnalytic: false,
		},
	}
	dataResult := mongo.UpdateMany(updateParams)

	model := FdibModel{}
	dataResult.Model(&model)
	dataResult.Print()
	//insertResult.toAPI()

	fmt.Println("id--->>>", dataResult.Id, time.Since(start))
	fmt.Println("insertResult--->>>", model, time.Since(start))
	fmt.Println("insertResult--->>>", model.Etiqueta, time.Since(start))

}

func FindOneAndUpdate() {
	start := time.Now()

	filter := mongo.NewFilter()
	filter.Id("6384f0e452ed0e02aa02d688")

	input := mongo.NewUpdateInput()
	input.Set("nserlum", 999999)
	input.SetMap(map[string]any{
		"bola": "azul",
		"mala": "verde",
	})

	findParams := mongo.FindParams{
		Collection: "users",
		Connection: "123456abc",
		Database:   "api-kdl-test",
		Fields:     []string{"idPeca", "coletaRede", "etiqueta", "status", "tempo"},
	}

	updateParams := mongo.UpdateParams{
		Collection: "users",
		Connection: "123456abc",
		Database:   "api-kdl-test",
		Info:       nil,
		Input:      input,
		Filter:     filter,
		FindParams: findParams,
		UpdateOptions: mongo.UpdateOptions{
			ReturnOriginal: true,
			Upsert:         true,
		},
		DataLog: mongo.DataLog{
			Action:       "UPDATE_NSERLUM",
			SaveChange:   true,
			SaveHistory:  true,
			SaveInfo:     false,
			SaveAnalytic: false,
		},
	}
	dataResult := mongo.FindOneAndUpdate(updateParams)

	model := FdibModel{}
	dataResult.Model(&model)
	dataResult.Print()
	//insertResult.toAPI()

	fmt.Println("id--->>>", dataResult.Id, time.Since(start))
	fmt.Println("insertResult--->>>", model, time.Since(start))
	fmt.Println("insertResult--->>>", model.Etiqueta, time.Since(start))

}
