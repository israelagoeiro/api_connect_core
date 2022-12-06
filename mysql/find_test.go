package mysql_test

import (
	"database/sql"
	"fmt"
	"github.com/israelagoeiro/api_connect_core/mysql"
	"testing"
	"time"
)

func TestFind(t *testing.T) {
	start := time.Now()

	mysql.NewClient("./../config.env")

	fields := mysql.NewFields()
	//fields.All()
	fields.Add("id", "idPeca", "nserlum", "coletaRede", "etiqueta", "status", "ligadoParcial", "ligadoDecorrido")

	filter := mysql.NewFilter()
	//filter.Where("etiqueta = 'IP08526B'")
	//filter.And("idPeca = '78945612'")
	filter.Limit(1)

	findParams := mysql.FindParams{
		Collection: "users",
		Fields:     fields,
		Filter:     filter,
	}
	findParams.PrintQuery()

	dataResult := mysql.Find(findParams)

	var listResult []mysql.FdibModel
	dataResult.Rows(func(rows *sql.Rows) {
		model := mysql.FdibModel{}
		err := rows.Scan(&model.Id, &model.IdPeca, &model.Nserlum, &model.ColetaRede, &model.Etiqueta, &model.Status, &model.Tempo.LigadoParcial, &model.Tempo.LigadoDecorrido)
		if err != nil {
			fmt.Println("Error TestFindOne.rows.Scan", err.Error(), findParams.Fields.Values(findParams.Collection))
		}
		listResult = append(listResult, model)
	})
	fmt.Println("dataResult--->listResult", len(listResult), listResult, time.Since(start))

	mysql.Disconnect()

	expected := 1
	obtained := len(listResult)
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if expected > obtained {
		fmt.Println("expected", expected)
		fmt.Println("obtained", obtained)
		t.Errorf("expected %d, obtained %d", expected, obtained)
	}
}
func TestFindQuery(t *testing.T) {
	start := time.Now()

	mysql.NewClient("./../config.env")
	dataResult := mysql.FindQuery("SELECT users.id,users.idPeca,users.nserlum,users.coletaRede,users.etiqueta,users.status,users.ligadoParcial,users.ligadoDecorrido FROM users LIMIT 2 OFFSET 2")

	var listResult []mysql.FdibModel
	dataResult.Rows(func(rows *sql.Rows) {
		model := mysql.FdibModel{}
		err := rows.Scan(&model.Id, &model.IdPeca, &model.Nserlum, &model.ColetaRede, &model.Etiqueta, &model.Status, &model.Tempo.LigadoParcial, &model.Tempo.LigadoDecorrido)
		if err != nil {
			fmt.Println("Error TestFindOne.rows.Scan", err.Error())
		}
		listResult = append(listResult, model)
	})
	fmt.Println("dataResult--->listResult", len(listResult), listResult, time.Since(start))

	mysql.Disconnect()

	expected := 1
	obtained := len(listResult)
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if expected > obtained {
		fmt.Println("expected", expected)
		fmt.Println("obtained", obtained)
		t.Errorf("expected %d, obtained %d", expected, obtained)
	}
}

func TestFindOneFieldsAll(t *testing.T) {
	start := time.Now()

	mysql.NewClient("./../config.env")

	fields := mysql.NewFields()
	fields.All()

	filter := mysql.NewFilter()
	//filter.Where("etiqueta = 'IP08526B'")
	//filter.And("idPeca = '78945612'")
	filter.OrderBy("idPeca ASC")

	findParams := mysql.FindParams{
		Collection: "users",
		Fields:     fields,
		Filter:     filter,
	}
	findParams.PrintQuery()

	dataResult := mysql.FindOne(findParams)

	var listResult []mysql.FdibModel
	dataResult.Rows(func(rows *sql.Rows) {
		model := mysql.FdibModel{}
		err := rows.Scan(&model.Id, &model.IdPeca, &model.Nserlum, &model.ColetaRede, &model.Etiqueta, &model.Status, &model.Tempo.LigadoParcial, &model.Tempo.LigadoDecorrido)
		if err != nil {
			fmt.Println("Error TestFindOne.rows.Scan", err.Error(), findParams.Fields.Values(findParams.Collection))
		}
		listResult = append(listResult, model)
	})

	fmt.Println("findParams--->", findParams)
	fmt.Println("dataResult--->listResult", listResult)
	fmt.Println("dataResult--->", dataResult, time.Since(start))

	mysql.Disconnect()
}
func TestFindOneFieldsAdd(t *testing.T) {
	start := time.Now()

	mysql.NewClient("./../config.env")

	fields := mysql.NewFields()
	fields.Add("id", "idPeca AS idPeca", "nserlum", "coletaRede", "etiqueta", "status", "ligadoParcial", "ligadoDecorrido")

	filter := mysql.NewFilter()
	filter.Where("etiqueta = 'IP08526B'")
	filter.And("idPeca = '78945612'")

	findParams := mysql.FindParams{
		Collection: "users",
		Fields:     fields,
		Filter:     filter,
	}
	findParams.PrintQuery()

	dataResult := mysql.FindOne(findParams)

	var listResult []mysql.FdibModel
	dataResult.Rows(func(rows *sql.Rows) {
		model := mysql.FdibModel{}
		err := rows.Scan(&model.Id, &model.IdPeca, &model.Nserlum, &model.ColetaRede, &model.Etiqueta, &model.Status, &model.Tempo.LigadoParcial, &model.Tempo.LigadoDecorrido)
		if err != nil {
			fmt.Println("Error TestFindOne.rows.Scan", err.Error(), findParams.Fields.Values(findParams.Collection))
		}
		listResult = append(listResult, model)
	})

	fmt.Println("findParams--->", findParams)
	fmt.Println("dataResult--->listResult", listResult)
	fmt.Println("dataResult--->", dataResult, time.Since(start))

	mysql.Disconnect()
}
