package items

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	itemspb "github.com/ankitanwar/e-Commerce/Items-A.P.I/proto"
	db "github.com/ankitanwar/e-Commerce/Items-A.P.I/server/database"
	"github.com/grpc/grpc-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Item : Struct it contains all the value item has
type Item struct {
	ID                primitive.ObjectID `bson:"id"`
	Seller            int64              `bson:"seller"`
	Title             string             `bson:"title"`
	Description       string             `bson:"description"`
	Price             int64              `bson:"price"`
	AvailableQuantity int64              `bson:"available_quantity"`
	SoldQuantity      int64              `bson:"sold_quantity"`
	Status            string             `bson:"status"`
}

//ItemService : Services Available for items

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:4040")
	if err != nil {
		log.Fatalln("Unable to listen")

		panic(err)
	}
	opts := []grpc.ServerOption{}

	var ser *grpc.Server
	ser = grpc.NewServer(opts...)
	itemspb.RegisterItemServiceServer(ser, &ItemService{})

	go func() {
		if err := ser.Serve(lis); err != nil {
			log.Fatalln("Unable to server")
			panic(err)
		}
	}()

	//Waiting for the stop signal
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	fmt.Println("Stopping the server")
	ser.Stop()
	fmt.Println("closing the listner")
	lis.Close()
	fmt.Println("Closing MongoDB Sever")
	db.Client.Disconnect(context.TODO())

}
