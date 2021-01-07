package controller

import (
	"net/http"
	"strconv"

	"github.com/ankitanwar/e-Commerce/new/domain"
	"github.com/ankitanwar/e-Commerce/new/services"
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
	userID := c.Param("userID")
	id, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	token := c.Param("access_token")
	valid, validErr := services.ValidateAccessToken(id, token)
	if validErr != nil {
		c.JSON(validErr.Status, validErr)
		return
	}
	c.JSON(http.StatusOK, valid)
	return
}
