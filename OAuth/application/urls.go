package application

import controller "github.com/ankitanwar/e-Commerce/Oauth/controllers"

func mapURL() {
	router.POST("/access_token", controller.CreateAccessToken)
	router.GET("/validate", controller.ValidateAccessToken)
}
