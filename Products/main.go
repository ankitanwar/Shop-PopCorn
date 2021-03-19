package main

import (
	"os"
	"os/signal"

	"github.com/ankitanwar/e-Commerce/Products/client/application"
	server "github.com/ankitanwar/e-Commerce/Products/server/startServer"
)

func main() {
	go func() {
		server.StartServer()
	}()
	go func() {
		application.StartClient()
	}()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}
