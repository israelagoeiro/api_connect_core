package mysql_test

import (
	"fmt"
	"github.com/israelagoeiro/api_connect_core/mysql"
	"testing"
	"time"
)

func TestInsertMany(t *testing.T) {
	mysql.NewClient("./../config.env")

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

	data1 := map[string]any{
		"IdPeca":          78945123,
		"ColetaRede":      time.Now().Format("2006-01-02 15:04:05"),
		"Etiqueta":        "IP08530A",
		"Nserlum":         654321,
		"Status":          1,
		"LigadoParcial":   time.Now().Format("2006-01-02"),
		"LigadoDecorrido": dataModel.Tempo.LigadoDecorrido.Format("2006-01-02"),
	}
	data2 := map[string]any{
		"IdPeca":          78945456,
		"ColetaRede":      time.Now().Format("2006-01-02 15:04:05"),
		"Etiqueta":        "IP08531A",
		"Nserlum":         654321,
		"Status":          1,
		"LigadoParcial":   time.Now().Format("2006-01-02"),
		"LigadoDecorrido": dataModel.Tempo.LigadoDecorrido.Format("2006-01-02"),
	}
	data3 := map[string]any{
		"IdPeca":          78945789,
		"ColetaRede":      time.Now().Format("2006-01-02 15:04:05"),
		"Etiqueta":        "IP08532A",
		"Nserlum":         654321,
		"Status":          0,
		"LigadoParcial":   time.Now().Format("2006-01-02"),
		"LigadoDecorrido": dataModel.Tempo.LigadoDecorrido.Format("2006-01-02"),
	}

	input := mysql.NewInsertInput()
	input.Map(data1)
	input.Map(data2)
	input.Map(data3)

	insertParams := mysql.InsertParams{
		Collection: "users",
		Input:      input,
	}
	insertParams.PrintQuery()

	//dataResult := mysql.InsertMany(insertParams)
	dataResult := mysql.InsertQuery("INSERT INTO users (ColetaRede,Etiqueta,IdPeca,LigadoDecorrido,LigadoParcial,Nserlum,Status) VALUES ('2022-12-05 16:23:59','IP08530A','78945123','2022-12-05','2022-12-05','654321','1'),('2022-12-05 16:23:59','IP08531A','78945456','2022-12-05','2022-12-05','654321','1'),('2022-12-05 16:23:59','IP08532A','78945789','2022-12-05','2022-12-05','654321','0');")

	expected := 3
	obtained := dataResult.RowsAffected()

	mysql.Disconnect()

	if expected != obtained {
		fmt.Println("expected", expected)
		fmt.Println("obtained", obtained)
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}

func TestInsertOneMap(t *testing.T) {
	mysql.NewClient("./../config.env")

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
		"Status":          1,
		"LigadoParcial":   time.Now().Format("2006-01-02"),
		"LigadoDecorrido": dataModel.Tempo.LigadoDecorrido.Format("2006-01-02"),
	}

	input := mysql.NewInsertInput()
	input.Map(dataMap)

	insertParams := mysql.InsertParams{
		Collection: "users",
		Input:      input,
	}
	insertParams.PrintQuery()

	dataResult := mysql.InsertOne(insertParams)

	expected := 1
	obtained := dataResult.RowsAffected()
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if expected != obtained {
		fmt.Println("expected", expected)
		fmt.Println("obtained", obtained)
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}

func TestInsertOneData(t *testing.T) {
	mysql.NewClient("./../config.env")

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

	input := mysql.NewInsertInput()
	input.Data("IdPeca", 78945612)
	input.Data("ColetaRede", time.Now().Format("2006-01-02"))
	input.Data("Etiqueta", "IP08526A")
	input.Data("Nserlum", 654321)
	input.Data("Status", 1)
	input.Data("LigadoParcial", time.Now().Format("2006-01-02"))
	input.Data("LigadoDecorrido", dataModel.Tempo.LigadoDecorrido.Format("2006-01-02"))

	insertParams := mysql.InsertParams{
		Collection: "users",
		Input:      input,
	}
	insertParams.PrintQuery()

	dataResult := mysql.InsertOne(insertParams)

	expected := 1
	obtained := dataResult.RowsAffected()

	mysql.Disconnect()

	if expected != obtained {
		fmt.Println("expected", expected)
		fmt.Println("obtained", obtained)
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}

func TestInsertOneAndFindParams(t *testing.T) {
	mysql.NewClient("./../config.env")

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
		"Status":          1,
		"LigadoParcial":   time.Now().Format("2006-01-02"),
		"LigadoDecorrido": dataModel.Tempo.LigadoDecorrido.Format("2006-01-02"),
	}

	input := mysql.NewInsertInput()
	input.Map(dataMap)

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

	expected := 1
	obtained := dataResult.RowsAffected()

	mysql.Disconnect()

	t.Errorf("expected %q, obtained %q", expected, obtained)

	/*if expected != obtained {
		fmt.Println("expected", expected)
		fmt.Println("obtained", obtained)
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}*/
}

func TestInsertSimucLogsdia(t *testing.T) {
	mysql.NewClient("./../config.env")

	dataMap := map[string]any{
		"codusr":   3,
		"nsercon":  2245041,
		"nserlum":  2031024428,
		"consumo":  524,
		"tmpac":    100,
		"framesv":  1,
		"hdestd":   time.Now().Format("2006-01-02"),
		"atuestd":  time.Now().Format("2006-01-02"),
		"ndestd":   time.Now().Format("2006-01-02"),
		"hacendeu": time.Now().Format("2006-01-02"),
		"hapagou":  time.Now().Format("2006-01-02"),
		"whacc":    125,
		"dtestd":   25,
	}
	input := mysql.NewInsertInput()
	input.Map(dataMap)

	insertParams := mysql.InsertParams{
		Collection: "logsdia",
		Input:      input,
	}
	insertParams.PrintQuery()

	mysql.Disconnect()

	/*dataResult := mysql.InsertOne(insertParams)

	expected := 1
	obtained := dataResult.RowsAffected()
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)*/

	//t.Errorf("expected %q, obtained %q", expected, obtained)

	/*if expected != obtained {
		fmt.Println("expected", expected)
		fmt.Println("obtained", obtained)
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}*/
}
