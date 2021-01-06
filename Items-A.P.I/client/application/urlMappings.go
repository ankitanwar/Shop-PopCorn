package application

import "github.com/ankitanwar/e-Commerce/Items-A.P.I/client/controllers"

func urlMapping() {
	router.GET("/ping", controllers.PingController.Ping)
	router.POST("/items/:access_token", controllers.ItemController.Create)
	router.GET("/items/:id", controllers.ItemController.Get)
	router.DELETE("/items/search/:access_token", controllers.ItemController.Delete)
}
