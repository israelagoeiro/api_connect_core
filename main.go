package main

import (
	"github.com/israelagoeiro/api_connect_core/db"
	"github.com/israelagoeiro/api_connect_core/examples"
	"github.com/israelagoeiro/api_connect_core/util"
)

func init() {
	util.LoadEnv()
}

func main() {
	db.DbMong = db.MongoDBOpen()

	examples.MongoInsertMany()
	examples.MongoInsertOne()

	examples.MongoFind()
	examples.MongoFindOne()

	examples.MongoUpdateMany()
	examples.MongoUpdateOne()

	examples.MongoDeleteMany()
	examples.MongoDeleteOne()
	examples.FindOneAndUpdate()
}
