// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package mongo provides functionality for measuring and displaying time.
//
// The calendrical calculations always assume a Gregorian calendar, with
// no leap seconds.
//
// Monotonic Clocks
//
// Operating systems provide both a “wall clock,” which is subject to
// changes for clock synchronization, and a “monotonic clock,” which is
// not. The general rule is that the wall clock is for telling time and
// the monotonic clock is for measuring time. Rather than split the API,
// in this package the Time returned by time.Now contains both a wall
// clock reading and a monotonic clock reading; later time-telling
// operations use the wall clock reading, but later time-measuring
// operations, specifically comparisons and subtractions, use the
// monotonic clock reading.
package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/israelagoeiro/api_connect_core/util"
	"os"
	"strconv"
)

var Client *sql.DB

func NewClient(filename string) {
	util.LoadEnv(filename)
	Client = newClient()
}

func newClient() *sql.DB {
	hostname := os.Getenv("MYSQL_DB_HOSTNAME")
	port := os.Getenv("MYSQL_DB_PORT")
	username := os.Getenv("MYSQL_DB_USERNAME")
	password := os.Getenv("MYSQL_DB_PASSWORD")
	schema := os.Getenv("MYSQL_DB_SCHEMA")
	maxIdleConns, _ := strconv.Atoi(os.Getenv("MAX_IDLE_CONNS"))

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, hostname, port, schema)
	client, err := sql.Open("mysql", dataSourceName)
	client.SetMaxIdleConns(maxIdleConns)

	if err != nil {
		fmt.Println("Error NewClient:Mysql", err.Error())
	}

	fmt.Println("NewClient: Connected to Mysql")
	return client
}

func Disconnect() {
	fmt.Println("Connection to MYSQL closed.")
	err := Client.Close()
	if err != nil {
		fmt.Println("Error TestInsertOne:Client.Close", err.Error())
	}
}
