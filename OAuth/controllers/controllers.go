package controller

import (
	"net/http"

	"github.com/ankitanwar/Shop-PopCorn/Oauth/domain"
	"github.com/ankitanwar/Shop-PopCorn/Oauth/services"
	"github.com/gin-gonic/gin"
)

//CreateAccessToken : To get the new access token
func CreateAccessToken(c *gin.Context) {
	req := &domain.LoginRequest{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Unable To Fetch The Details")
		return
	}
	token, restErr := services.CreateAccessToken(req)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusAccepted, token)

}

//ValidateAccessToken : To validate the access token
func ValidateAccessToken(c *gin.Context) {
	userID := c.Request.Header.Get("X-Caller-ID")
	if userID == "" {
		c.JSON(http.StatusBadRequest, "Invalid userID")
	}
	token := c.Request.Header.Get("X-Token-ID")
	valid, validErr := services.ValidateAccessToken(userID, token)
	if validErr != nil {
		c.JSON(validErr.Status, validErr)
		return
	}
	c.JSON(http.StatusOK, valid)
	return
}

//RemoveAccessToken : To logout the user and remove the access Token
func RemoveAccessToken(c *gin.Context) {
	userID := c.Request.Header.Get("X-Caller-Id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, "Invalid Caller ID")
		return
	}
	accessToken := c.Request.Header.Get("X-Token-ID")
	removeErr := services.RemoveAccessToken(userID, accessToken)
	if removeErr != nil {
		c.JSON(removeErr.Status, removeErr.Message)
		return
	}
	c.JSON(http.StatusAccepted, "User Has Been logged Out Successfully")
}
