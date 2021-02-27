package controller

import (
	"net/http"

	"github.com/ankitanwar/e-Commerce/Oauth/domain"
	"github.com/ankitanwar/e-Commerce/Oauth/services"
	"github.com/gin-gonic/gin"
)

//CreateAccessToken : To get the new access token
func CreateAccessToken(c *gin.Context) {
	req := &domain.LoginRequest{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Unable To Bind")
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
	userID := c.Request.Header.Get("X-Caller-Id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, "Invalid userID")
	}
	token := c.Request.Header.Get("access_token")
	valid, validErr := services.ValidateAccessToken(userID, token)
	if validErr != nil {
		c.JSON(validErr.Status, validErr)
		return
	}
	c.JSON(http.StatusOK, valid)
	return
}
