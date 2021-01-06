package main

import (
	"log"
	"os"

	"github.com/ankitanwar/e-Commerce/User/app"
	"github.com/joho/godotenv"
)

var (
	username = os.Getenv("MYSQL_USER")
	password = os.Getenv("MYSQL_PASSWORD")
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Could not connect to env ", err)
		panic(err)
	}
	app.StartApplication()
}
