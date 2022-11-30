package main

import (
	"github.com/israelagoeiro/api_connect_core/db"
	"github.com/israelagoeiro/api_connect_core/examples"
	"github.com/joho/godotenv"
	"log"
)

func loadEnv() {
	// load .env file
	err := godotenv.Load("config.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func init() {
	loadEnv()
}

func main() {
	db.DbMong = db.MongoDBOpen()
	//examples.MongoInsert()
	examples.MongoFind()
	//examples.MongoUpdate()
	examples.MongoDelete()
}
