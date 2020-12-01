package application

import "github.com/ankitanwar/userLoginWithOAuth/Items-A.P.I/client/controllers"

func urlMapping() {
	router.GET("/ping", controllers.PingController.Ping)
	router.POST("/items", controllers.ItemController.Create)
	router.GET("/items/:id", controllers.ItemController.Get)
	router.GET("/items/search/:title", controllers.ItemController.Search)
}
