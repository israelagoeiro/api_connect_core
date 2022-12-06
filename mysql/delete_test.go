package mysql_test

import (
	"fmt"
	"github.com/israelagoeiro/api_connect_core/mysql"
	"testing"
)

func TestDeleteMany(t *testing.T) {
	mysql.NewClient("./../config.env")

	filter := mysql.NewFilter()
	filter.Where("etiqueta = 'IP08526A'")
	filter.OrderBy("etiqueta ASC")
	filter.Limit(1)

	deleteParams := mysql.DeleteParams{
		Collection: "users",
		Filter:     filter,
	}
	deleteParams.PrintQuery()
	deleteResult := mysql.DeleteMany(deleteParams)

	mysql.Disconnect()

	expected := 1
	obtained := deleteResult.TotalAfected()
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if expected > obtained {
		fmt.Println("expected", expected)
		fmt.Println("obtained", obtained)
		t.Errorf("expected %d, obtained %d", expected, obtained)
	}
}

func TestDeleteOne(t *testing.T) {
	mysql.NewClient("./../config.env")

	filter := mysql.NewFilter()
	filter.Where("etiqueta = 'IP08526A'")
	filter.OrderBy("etiqueta ASC")

	deleteParams := mysql.DeleteParams{
		Collection: "users",
		Filter:     filter,
	}
	deleteParams.PrintQuery()
	deleteResult := mysql.DeleteOne(deleteParams)

	mysql.Disconnect()

	expected := 1
	obtained := deleteResult.TotalAfected()
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if expected > obtained {
		fmt.Println("expected", expected)
		fmt.Println("obtained", obtained)
		t.Errorf("expected %d, obtained %d", expected, obtained)
	}
}
