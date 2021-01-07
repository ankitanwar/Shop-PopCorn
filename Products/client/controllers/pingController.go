package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	pong = "pong"

	//PingController : Controller for the ping
	PingController pingInterface = &pingStruct{}
)

type pingInterface interface {
	Ping(*gin.Context)
}

type pingStruct struct {
}

//Ping : To ping to the database to test the server is running or not
func (p *pingStruct) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
	return
}
