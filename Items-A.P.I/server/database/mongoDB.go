package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	//Client :- grpc Client
	Client *mongo.Client

	//Collection :- grpc Collection Database
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
	Collection = Client.Database("grpc").Collection("Items")
}
