package application

import (
	"github.com/ankitanwar/userLoginWithOAuth/Oauth/services"

	mongod "github.com/ankitanwar/userLoginWithOAuth/Oauth/clients"
	"github.com/ankitanwar/userLoginWithOAuth/Oauth/http"
	"github.com/ankitanwar/userLoginWithOAuth/Oauth/repository/db"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

//StartApplication : To start the application
func StartApplication() {
	mongod.Ping()
	dbRepository := db.NewRepository()
	userRepository := db.NewRestRepository()
	atService := services.NewService(dbRepository, userRepository)
	atHandler := http.NewHandler(atService)
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)
	router.POST("/oauth/access_token", atHandler.Create)
	router.Run(":8083")
}
