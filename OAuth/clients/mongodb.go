package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	//Client :pointer to the mongo clinet
	Client *mongo.Client
)

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	Client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

}

//GetSession : To create the session with mongoDB Dont Forget To close the connection after invoking it
func GetSession() (context.Context, context.CancelFunc) {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	if err != nil {
		fmt.Println("Error while connection to the database")
		panic(err)
	}

	return ctx, cancel

}

//Ping To the database
func Ping() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := Client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("Error while ping the database")
		panic(err)
	}

}
