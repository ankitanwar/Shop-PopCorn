package main

import (
	"log"

	"github.com/ankitanwar/user-api/app"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Could not connect to env ", err)
		panic(err)
	}
	app.StartApplication()
}
