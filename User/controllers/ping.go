package controllers

import (
	"github.com/gin-gonic/gin"
)

//Ping : To check whether the server is up and running
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}
