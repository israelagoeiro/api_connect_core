package mysql_test

import (
	"fmt"
	"github.com/israelagoeiro/api_connect_core/mysql"
	"testing"
	"time"
)

func TestUpdateInputMap(t *testing.T) {
	dataMap := map[string]any{
		"IdPeca":        78945000,
		"Etiqueta":      "IP085000",
		"Nserlum":       654000,
		"Status":        0,
		"LigadoParcial": time.Now().Format("2006-01-02"),
	}

	input := mysql.NewUpdateInput()
	input.Map(dataMap)

	expected := "SET (Etiqueta='IP085000',IdPeca='78945000',LigadoParcial='2022-12-05',Nserlum='654000',Status='0')"
	obtained := input.Values()

	if expected != obtained {
		fmt.Println("expected", expected)
		fmt.Println("obtained", obtained)
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}

func TestUpdateInputData(t *testing.T) {
	input := mysql.NewUpdateInput()
	input.Data("IdPeca", 78945612)
	input.Data("Etiqueta", "IP08526A")
	input.Data("Nserlum", 654321)
	input.Data("Status", 1)
	input.Data("LigadoParcial", time.Now().Format("2006-01-02"))

	expected := "SET (Etiqueta='IP08526A',IdPeca='78945612',LigadoParcial='2022-12-05',Nserlum='654321',Status='1')"
	obtained := input.Values()

	if expected != obtained {
		fmt.Println("expected", expected)
		fmt.Println("obtained", obtained)
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}
