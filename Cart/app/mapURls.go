package application

import controllers "github.com/ankitanwar/e-Commerce/Cart/controller"

func mapUrls() {
	router.POST("/user/cart/:itemID", controllers.AddToCart)
	router.DELETE("/user/cart/:itemID", controllers.RemoveFromCart)
	router.GET("/user/cart", controllers.ViewCart)
}
