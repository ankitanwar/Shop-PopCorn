package connect

import (
	"fmt"

	itemspb "github.com/ankitanwar/Shop-PopCorn/Products/proto"
	"google.golang.org/grpc"
)

var (
	Client itemspb.ItemServiceClient
	CC     *grpc.ClientConn
)

//ConnectServer : To Connect To the gRPC server
func ConnectServer() {
	opts := grpc.WithInsecure()
	var err error
	CC, err = grpc.Dial("localhost:4040", opts)
	if err != nil {
		fmt.Println("Error while connection to the server", err.Error())
		panic(err)
	}
	Client = itemspb.NewItemServiceClient(CC)
	fmt.Println("Connection to Server is successfull")
}
