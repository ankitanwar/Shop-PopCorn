package application

import (
	"fmt"
	"os"
	"os/signal"

	connect "github.com/ankitanwar/e-Commerce/Products/client/connectToServer"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

//StartClient : To start the client service
func StartClient() {
	urlMapping()
	connect.ConnectServer()
	go func() {
		router.Run(":8086")
	}()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	fmt.Println("Closing the Connection with server")
	connect.CC.Close()

}
