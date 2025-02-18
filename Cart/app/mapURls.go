package application

import controllers "github.com/ankitanwar/Shop-PopCorn/Cart/controller"

func mapUrls() {
	router.POST("/user/cart/:itemID", controllers.AddToCart)
	router.DELETE("/user/cart/:itemID", controllers.RemoveFromCart)
	router.GET("/user/cart", controllers.ViewCart)
	router.POST("/cart/checkout/:addressID", controllers.Checkout)
}
