package controllers

import (
	"context"
	"errors"
	"net/http"

	itemspb "github.com/ankitanwar/e-Commerce/Items-A.P.I/proto"
	"github.com/ankitanwar/e-Commerce/Items-A.P.I/server/domian/items"
	oauth "github.com/ankitanwar/e-Commerce/interactWithOAuth/oAuth"
	"github.com/gin-gonic/gin"
)

var (
	//ItemController : Methods available for items
	ItemController itemControllerInterface = &itemControllerStruct{}
)

type itemControllerInterface interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
	Delete(c *gin.Context)
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
func (i *itemControllerStruct) Buy(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err)
		return
	}
	itemID := c.Param("itemsID")
	userID := int64(oauth.GetClientID(c.Request))
	buyRequest := &itemspb.BuyItemRequest{ItemID: itemID, UserID: userID}
	item, err := items.Services.Buy(context.Background(), buyRequest)
	if err != nil {
		if err == errors.New("Out Of Stock") {
			c.JSON(http.StatusNotFound, "Item Is Currently Out Of stock")
		} else {
			c.JSON(http.StatusInternalServerError, "Some Error has been occured")
		}
		return
	}

	c.JSON(http.StatusOK, item)
}

//Get : To get the particaular item by given ID
func (i *itemControllerStruct) Get(c *gin.Context) {
	oid := c.Param("id")
	toFind := &itemspb.GetItemRequest{
		ID: oid,
	}
	res, err := items.Services.Get(context.Background(), toFind)
	if err != nil {
		c.JSON(http.StatusFound, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

//Search : To search for the items with given ID
func (i *itemControllerStruct) Delete(c *gin.Context) {
	oid := c.Param("id")
	toDelete := &itemspb.DeleteItemRequest{
		Id: oid,
	}
	res, err := items.Services.Delete(context.Background(), toDelete)
	if err != nil {
		c.JSON(http.StatusFound, err)
		return
	}
	c.JSON(http.StatusOK, res)
}
