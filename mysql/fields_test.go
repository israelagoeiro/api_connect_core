package mysql_test

import (
	"fmt"
	"github.com/israelagoeiro/api_connect_core/mysql"
	"testing"
)

func TestFieldsAdd(t *testing.T) {
	collection := "users"
	var fields = mysql.NewFields()
	fields.Add("id", "idPeca", "nserlum", "coletaRede", "etiqueta", "SUM(salary) AS 'Total Salary'")

	expected := "id,idPeca,nserlum,coletaRede,etiqueta,SUM(salary) AS 'Total Salary'"
	obtained := fields.Values(collection)
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if expected != obtained {
		fmt.Println("expected", expected)
		fmt.Println("obtained", obtained)
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}

func TestFieldsAll(t *testing.T) {
	collection := "users"
	var fields = mysql.NewFields()
	fields.All()

	expected := "*"
	obtained := fields.Values(collection)
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if expected != obtained {
		fmt.Println("expected", expected)
		fmt.Println("obtained", obtained)
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}
