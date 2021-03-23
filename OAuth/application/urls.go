package application

import controller "github.com/ankitanwar/Shop-PopCorn/Oauth/controllers"

func mapURL() {
	router.POST("/login", controller.CreateAccessToken)
	router.GET("/validate", controller.ValidateAccessToken)
	router.DELETE("/logout", controller.RemoveAccessToken)
}
