package items

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Item : Struct it contains all the value item has
type Item struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	Seller            string             `bson:"seller"`
	Title             string             `bson:"title"`
	Description       string             `bson:"description"`
	Price             int64              `bson:"price"`
	AvailableQuantity int64              `bson:"available_quantity"`
	SoldQuantity      int64              `bson:"sold_quantity"`
	Status            string             `bson:"status"`
}

//StartServer : To start the the item server
func StartServer() {
	// lis, err := net.Listen("tcp", "0.0.0.0:4040")
	// if err != nil {
	// 	log.Fatalln("Unable to listen")

	// 	panic(err)
	// }
	// opts := []grpc.ServerOption{}
	// srv := grpc.NewServer(opts...)
	// itemspb.RegisterItemServiceServer(srv, &ItemService{})

	// go func() {
	// 	if err := srv.Serve(lis); err != nil {
	// 		log.Fatalln("Unable to server")
	// 		panic(err)
	// 	}
	// }()

	// //Waiting for the stop signal
	// ch := make(chan os.Signal, 1)
	// signal.Notify(ch, os.Interrupt)
	// <-ch
	// fmt.Println("Stopping the server")
	// srv.Stop()
	// fmt.Println("closing the listner")
	// lis.Close()
	// fmt.Println("Closing MongoDB Sever")
	// db.Client.Disconnect(context.TODO())

}
