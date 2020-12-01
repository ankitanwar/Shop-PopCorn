package controllers

import (
	"context"
	"net/http"

	itemspb "github.com/ankitanwar/userLoginWithOAuth/Items-A.P.I/proto"
	"github.com/ankitanwar/userLoginWithOAuth/Items-A.P.I/server/domian/items"
	oauth "github.com/ankitanwar/userLoginWithOAuth/interactWithOAuth/oAuth"
	"github.com/gin-gonic/gin"
)

var (
	//ItemController : Methods available for items
	ItemController itemControllerInterface = &itemControllerStruct{}
)

type itemControllerInterface interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
	Search(c *gin.Context)
}
type itemControllerStruct struct {
}

//Create : To create new Item
func (i *itemControllerStruct) Create(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err)
		return
	}
	itemRequest := &itemspb.Item{}
	if err := c.ShouldBindJSON(itemRequest); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	itemRequest.Seller = int64(oauth.GetClientID(c.Request))
	result, createErr := items.Services.Create(context.Background(), &itemspb.CreateItemRequest{Item: itemRequest})
	if createErr != nil {
		c.JSON(http.StatusInternalServerError, createErr)
		return
	}
	c.JSON(http.StatusOK, result)

}

//Get : To get the particaular item by given ID
func (i *itemControllerStruct) Get(c *gin.Context) {

}

//Search : To search for the items with given ID
func (i *itemControllerStruct) Search(c *gin.Context) {

}
