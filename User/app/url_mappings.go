package app

import "github.com/ankitanwar/Shop-PopCorn/User/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)
	router.POST("/users", controllers.CreateUser)
	router.GET("/users", controllers.GetUser)
	router.PATCH("/users", controllers.UpdateUser)
	router.DELETE("/users", controllers.DeleteUser)
	router.POST("/user/verify", controllers.VerifyUser)
	router.GET("/user/address", controllers.GetAddress)
	router.POST("/user/address", controllers.AddAddress)
	router.GET("/user/specificaddress/:addressID", controllers.GetAddressWithID)
	router.DELETE("/user/address/:addressID", controllers.RemoveAddress)
}
