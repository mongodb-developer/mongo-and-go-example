package main

import (
	"context"
	"fmt"
	"time"

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

	// Create a context with a 30 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Define my document struct
	type Post struct {
		Title string `bson:"title,omitempty"`
		Body  string `bson:"body,omitempty"`
	}

	// Get my collection instance
	collection := client.Database("blog").Collection("posts")

	// Insert documents
	docs := []interface{}{
		Post{Title: "World", Body: "Hello World"},
		Post{Title: "Mars", Body: "Hello Mars"},
		Post{Title: "Pluto", Body: "Hello Pluto"},
	}

	res, err := collection.InsertMany(ctx, docs)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Inserted: %+v\n", res)

	// Find and print all posts in the collection
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		panic(err)
	}
	defer cur.Close(ctx)

	var posts []Post
	if err := cur.All(ctx, &posts); err != nil {
		panic(err)
	}
	fmt.Printf("Posts: %+v\n", posts)

}
