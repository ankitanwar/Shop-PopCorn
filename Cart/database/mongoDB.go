package cartdatabase

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	//Client :- Pointer to the MongoDB Client
	client *mongo.Client
	//Collection :- Pointer to the MongoDB collection
	collection *mongo.Collection
)

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalln("Unable to connect to mongoDB Sever")
		panic(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalln("Error while ping")
		panic(err)
	}
	collection = client.Database("Cart").Collection("Cart")
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		fmt.Println("Error while connecting to the mongoDB database")
		panic(err)
	}
}
