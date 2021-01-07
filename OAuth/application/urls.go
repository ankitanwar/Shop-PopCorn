package application

import controller "github.com/ankitanwar/e-Commerce/new/controllers"

func mapURL() {
	router.POST("/access_token", controller.CreateAccessToken)
	router.GET("/validate/:userID/:access_token")
}
