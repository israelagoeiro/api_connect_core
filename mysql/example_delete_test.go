// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mysql_test

import (
	"fmt"
	"github.com/israelagoeiro/api_connect_core/mysql"
	"time"
)

func ExampleDeleteMany() {
	start := time.Now()

	fields := mysql.NewFields()
	fields.All()
	fields.Add("id", "idPeca", "nserlum", "coletaRede", "etiqueta", "status", "ligadoParcial", "ligadoDecorrido")

	findParams := mysql.FindParams{
		Collection: "users",
		Fields:     fields,
	}

	deleteParams := mysql.DeleteParams{
		Collection: "users",
		FindParams: findParams,
	}
	dataResult := mysql.DeleteMany(deleteParams)

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

func ExampleDeleteOne() {
	start := time.Now()

	fields := mysql.NewFields()
	fields.All()
	fields.Add("id", "idPeca", "nserlum", "coletaRede", "etiqueta", "status", "ligadoParcial", "ligadoDecorrido")

	findParams := mysql.FindParams{
		Collection: "users",
		Fields:     fields,
	}

	deleteParams := mysql.DeleteParams{
		Collection: "users",
		FindParams: findParams,
	}
	dataResult := mysql.DeleteOne(deleteParams)

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
