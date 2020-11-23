package app

import "github.com/ankitanwar/user-api/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)

	router.POST("/users", controllers.CreateUser)
	router.GET("/users/:user_id", controllers.GetUser)
	router.PUT("/users/:user_id", controllers.UpdateUser)
	router.PATCH("/users/:user_id", controllers.UpdateUser)
	router.DELETE("/users/:user_id", controllers.DeleteUser)
	router.GET("/internal/users/search", controllers.FindByStatus)
	router.POST("/user/login", controllers.Login)
}
