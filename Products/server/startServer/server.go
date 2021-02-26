package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	itemspb "github.com/ankitanwar/e-Commerce/Products/proto"
	items "github.com/ankitanwar/e-Commerce/Products/server/services"
	"google.golang.org/grpc"
)

//StartServer : To start the the item server
func StartServer() {
	lis, err := net.Listen("tcp", "0.0.0.0:4040")
	if err != nil {
		log.Fatalln("Unable to listen")

		panic(err)
	}
	opts := []grpc.ServerOption{}
	srv := grpc.NewServer(opts...)
	itemspb.RegisterItemServiceServer(srv, &items.ItemService{})

	go func() {
		if err := srv.Serve(lis); err != nil {
			log.Fatalln("Unable to server")
			panic(err)
		}
	}()

	//Waiting for the stop signal
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	fmt.Println("Stopping the server")
	srv.Stop()
	fmt.Println("closing the listner")
	lis.Close()
	fmt.Println("Closing MongoDB Sever")
	db.Client.Disconnect(context.TODO())

}
