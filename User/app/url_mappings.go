package app

import "github.com/ankitanwar/e-Commerce/User/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)
	router.POST("/users", controllers.CreateUser)
	router.GET("/users/:user_id", controllers.GetUser)
	router.PATCH("/users", controllers.UpdateUser)
	router.DELETE("/users", controllers.DeleteUser)
	router.GET("/internal/users/search", controllers.FindByStatus)
	router.POST("/user/login", controllers.Login)
	router.GET("/user/address", controllers.GetAddress)
	router.POST("/user/address", controllers.AddAddress)
}
