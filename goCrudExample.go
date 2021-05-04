package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	/*
			Connect to my cluster
	*/
	client, err := mongo.NewClient(options.Client().ApplyURI("<ATLAS_URI>"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)


	// Interact with data
		type Post struct {
	    Title   string   `bson:"title,omitempty"`
	    Body    string   `bson:"body,omitempty"`
	 }


	 /*
	 		Get my collection instance
	 */
		collection := client.Database("blog").Collection("posts")


		/*
	  	Insert documents
	  */
		docs := []interface{}{
			bson.D{{"title" , "World"},{"body" , "Hello World"}},
			bson.D{{"title" , "Mars"},{"body" , "Hello Mars"}},
			bson.D{{"title" , "Pluto" }, {"body" , "Hello Pluto"}},
		}

		res, insertErr := collection.InsertMany(ctx, docs)
		if insertErr != nil {
			log.Fatal(insertErr)
		}
    fmt.Println(res);
	/*
			Iterate a cursor
		*/
		cur, currErr := collection.Find(ctx, bson.D{})

		if currErr != nil { panic(currErr) }
		defer cur.Close(ctx)

		var posts []Post
		if err = cur.All(ctx, &posts); err != nil {
	       panic(err)
		}
		fmt.Println(posts)


}
