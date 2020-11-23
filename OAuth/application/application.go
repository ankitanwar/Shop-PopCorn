package application

import (
	mongodb "github.com/ankitanwar/OAuth/clients"
	"github.com/ankitanwar/OAuth/http"
	"github.com/ankitanwar/OAuth/repository/db"
	"github.com/ankitanwar/OAuth/services"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

//StartApplication : To start the application
func StartApplication() {
	mongodb.Ping()
	dbRepository := db.NewRepository()
	userRepository := db.NewRestRepository()
	atService := services.NewService(dbRepository, userRepository)
	atHandler := http.NewHandler(atService)
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)
	router.POST("/oauth/access_token", atHandler.Create)
	router.Run(":8080")
}
