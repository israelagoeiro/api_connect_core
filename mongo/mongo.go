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
package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

var Client *mongo.Client

func NewClient() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_DB_CONNECT")))

	if err != nil {
		fmt.Println("Error NewClient", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 90*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("NewClient: Connected to MongoDB")

	return client
}

func Disconnect() {
	if Client == nil {
		return
	}

	err := Client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// TODO optional you can log your closed MongoDB client
	fmt.Println("Connection to MongoDB closed.")
}

func GetCollection(database string, collectionName string) *mongo.Collection {
	return Client.Database(database).Collection(collectionName)
}
