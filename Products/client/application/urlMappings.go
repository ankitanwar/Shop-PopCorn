package application

import "github.com/ankitanwar/e-Commerce/Products/client/controllers"

func urlMapping() {
	router.GET("/ping", controllers.PingController.Ping)
	router.POST("/items/sellitem", controllers.ItemController.Create)
	router.GET("/items/:id", controllers.ItemController.Get)
	router.DELETE("/items/:id", controllers.ItemController.Delete)
	router.POST("/items/buy/:itemsID", controllers.ItemController.Buy)
	router.PATCH("/items/:id", controllers.ItemController.Update)
	router.POST("/seller/items/:id", controllers.ItemController.SellerView)
}
