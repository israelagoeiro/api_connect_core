package mysql_test

import (
	"fmt"
	"github.com/israelagoeiro/api_connect_core/mysql"
	"time"
)

func ExampleInsertOne() {
	start := time.Now()

	dataModel := mysql.FdibModel{
		IdPeca:     78945612,
		ColetaRede: time.Now(),
		Etiqueta:   "IP08526A",
		Nserlum:    654321,
		Status:     true,
		Tempo: mysql.TempoModel{
			LigadoParcial:   time.Now(),
			LigadoDecorrido: time.Now(),
		},
	}

	dataMap := map[string]any{
		"IdPeca":          78945612,
		"ColetaRede":      time.Now().Format("2006-01-02 15:04:05"),
		"Etiqueta":        "IP08526A",
		"Nserlum":         654321,
		"Status":          true,
		"LigadoParcial":   time.Now().Format("2006-01-02"),
		"LigadoDecorrido": dataModel.Tempo.LigadoDecorrido.Format("2006-01-02"),
	}

	input := mysql.NewInsertInput()
	input.Map(dataMap)
	input.Data("IdPeca", dataModel.IdPeca)
	input.Data("LigadoParcial", dataModel.Tempo.LigadoParcial.Format("2006-01-02 15:04:05"))

	fields := mysql.NewFields()
	fields.All()
	fields.Add("id", "idPeca", "nserlum", "coletaRede", "etiqueta", "status", "ligadoParcial", "ligadoDecorrido")

	findParams := mysql.FindParams{
		Collection: "users",
		Fields:     fields,
	}

	insertParams := mysql.InsertParams{
		Collection: "users",
		Input:      input,
		FindParams: findParams,
	}
	dataResult := mysql.InsertOne(insertParams)

	model := mysql.FdibModel{}
	dataResult.Model(&model)
	dataResult.Print()
	//insertResult.toAPI()

	//fmt.Println("id--->>>", dataResult.Id, time.Since(start))
	fmt.Println("insertResult--->>>", model, time.Since(start))
	fmt.Println("insertResult--->>>", model.Etiqueta, time.Since(start))

	// Output:
	// 1257894000000
	// 2009-11-10 23:00:00 +0000 UTC
}

func ExampleInsertMany() {
	start := time.Now()
	/*dataModel := mysql.FdibModel{
		IdPeca:     78945612,
		ColetaRede: time.Now(),
		Etiqueta:   "IP08526A",
		Nserlum:    654321,
		Status:     true,
		Tempo: mysql.TempoModel{
			LigadoParcial:   time.Now(),
			LigadoDecorrido: time.Now(),
		},
	}*/

	input := mysql.NewInsertInput()
	//input.Model(dataModel)
	input.Data("nserlum", 123456)

	insertParams := mysql.InsertParams{
		Collection: "users",
		Input:      input,
	}
	dataResult := mysql.InsertMany(insertParams)

	model := mysql.FdibModel{}
	dataResult.Model(&model)
	dataResult.Print()
	//insertResult.toAPI()

	//fmt.Println("id--->>>", dataResult.Id, time.Since(start))
	fmt.Println("insertResult--->>>", model, time.Since(start))
	fmt.Println("insertResult--->>>", model.Etiqueta, time.Since(start))

	// Output:
	// 1257894000000
	// 2009-11-10 23:00:00 +0000 UTC
}
