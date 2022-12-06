package mysql_test

import (
	"fmt"
	"github.com/israelagoeiro/api_connect_core/mysql"
	"testing"
	"time"
)

func TestUpdateOneMap(t *testing.T) {
	mysql.NewClient("./../config.env")

	dataMap := map[string]any{
		"IdPeca":        78945000,
		"Etiqueta":      "IP085000",
		"Nserlum":       654000,
		"Status":        0,
		"LigadoParcial": time.Now().Format("2006-01-02"),
	}

	input := mysql.NewUpdateInput()
	input.Map(dataMap)

	filter := mysql.NewFilter()
	filter.Where("etiqueta = 'IP08530A'")

	insertParams := mysql.UpdateParams{
		Collection: "users",
		Input:      input,
		Filter:     filter,
	}
	insertParams.PrintQuery()
	updateResult := mysql.UpdateOne(insertParams)

	fmt.Println("updateResult-->>", updateResult)

	/*


		expected := 1
		obtained := dataResult.RowsAffected()
		fmt.Println("expected", expected)
		fmt.Println("obtained", obtained)

		if expected != obtained {
			fmt.Println("expected", expected)
			fmt.Println("obtained", obtained)
			t.Errorf("expected %q, obtained %q", expected, obtained)
		}*/
}
