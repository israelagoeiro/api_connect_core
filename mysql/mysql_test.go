package mysql_test

import (
	"fmt"
	"github.com/israelagoeiro/api_connect_core/mysql"
	"testing"
)

func TestNewClient(t *testing.T) {
	mysql.NewClient("./../config.env")

	fmt.Println("expected", mysql.Client)

	if mysql.Client == nil {
		t.Errorf("NewClient %s", "false")
	}
}
