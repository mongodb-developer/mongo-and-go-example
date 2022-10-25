package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// Connect to my cluster
	client, err := mongo.Connect(
		context.Background(),
		options.Client().ApplyURI("<ATLAS_URI>"))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.Background())

	// List databases
	databases, err := client.ListDatabaseNames(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
	fmt.Println(databases)

}
