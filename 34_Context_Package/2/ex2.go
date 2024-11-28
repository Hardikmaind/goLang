package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	//mongo db connection uri
	clientoptions := options.Client().ApplyURI("mongodb://localhost:27017")

	//set a timeout using a context
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	//connect to the mongo db and ping the server
	client, err := mongo.Connect(ctx, clientoptions) // if this takes 6 seconds
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary()) //? <= CONTEXT TIMEOUT ERROR		//* and if the ping function takes 5 seconds then it will go over the timeout of our context (which is 10) and total time req is 11.
	if err != nil {
		log.Fatal(err)
	}





	fmt.Println("")
	fmt.Println("========================================CONTEXT 2.===================================================")




	
	//* CONTEXT 2. example context for searching the the db. if takes longer than 10 seecnds then we will cancel it
	ctx2, cancel2 := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel2()

	//Select the database and the collection
	collection := client.Database("api").Collection("users")

	//Query the database
	var res bson.M
	filter := bson.M{"name": "Alex"}
	err2 := collection.FindOne(ctx2, filter).Decode(&res)
	if err2 != nil {
		log.Fatal(err)
	}
}
