package application

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	mongo "github.com/ankitanwar/Shop-PopCorn/Oauth/database"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

//StartApplication : To start the application
func StartApplication() {
	mapURL()
	go func() {
		router.Run(":8090")
	}()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	fmt.Println("Closing the Connection with server")
	err := mongo.Client.Disconnect(context.Background())
	if err != nil {
		log.Fatalln("Error While Closing The Server")
	}
}
