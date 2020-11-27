package app

import (
	"os"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

//StartApplication : To start the application
func StartApplication() {
	mapUrls()
	router.Run(os.Getenv("PORT"))
}
