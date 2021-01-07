package mongo

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
	Client *mongo.Client

	//Collection :- Pointer to the MongoDB collection
	Collection *mongo.Collection
)

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	Client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalln("Unable to connect to mongoDB Sever")
		panic(err)
	}
	err = Client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalln("Error while ping")
		panic(err)
	}
	Collection = Client.Database("Users").Collection("AccessToken")
	err = Client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		fmt.Println("Error while connecting to the mongoDB database")
		panic(err)
	}
}
