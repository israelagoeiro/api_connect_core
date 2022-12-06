package mysql_test

import (
	"fmt"
	"github.com/israelagoeiro/api_connect_core/mysql"
	"testing"
	"time"
)

func TestInsertInputMap(t *testing.T) {
	dataMap := map[string]any{
		"IdPeca":        78945612,
		"Etiqueta":      "IP08526A",
		"Nserlum":       654321,
		"Status":        true,
		"LigadoParcial": time.Now().Format("2006-01-02"),
	}

	input := mysql.NewInsertInput()
	input.Map(dataMap)
	field, values := input.Values()

	wantField := "(Etiqueta,IdPeca,LigadoParcial,Nserlum,Status)"
	wantValues := "('IP08526A','78945612','2022-12-05','654321','true')"
	fmt.Println("expected", field, values)
	fmt.Println("obtained", wantField, wantValues)

	if field != wantField || values != wantValues {
		t.Errorf("expected %q, obtained %q", field, wantField)
		t.Errorf("expected %q, obtained %q", values, wantValues)
	}
}

func TestInsertInputData(t *testing.T) {
	input := mysql.NewInsertInput()
	input.Data("IdPeca", 78945612)
	input.Data("Etiqueta", "IP08526A")
	input.Data("Nserlum", 654321)
	input.Data("Status", true)
	input.Data("LigadoParcial", time.Now().Format("2006-01-02"))

	field, values := input.Values()

	wantField := "(Etiqueta,IdPeca,LigadoParcial,Nserlum,Status)"
	wantValues := "('IP08526A','78945612','2022-12-05','654321','true')"
	fmt.Println("expected", field, values)
	fmt.Println("obtained", wantField, wantValues)

	if field != wantField || values != wantValues {
		t.Errorf("expected %q, obtained %q", field, wantField)
		t.Errorf("expected %q, obtained %q", values, wantValues)
	}
}
