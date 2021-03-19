package controllers

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	oauth "github.com/ankitanwar/e-Commerce/Middleware/oAuth"
	connect "github.com/ankitanwar/e-Commerce/Products/client/connectToServer"
	itemspb "github.com/ankitanwar/e-Commerce/Products/proto"
	"github.com/gin-gonic/gin"
)

var (
	//ItemController : Methods available for items
	ItemController itemControllerInterface = &itemControllerStruct{}
)

//getUserID : To get the ID of the user
func getUserID(req *http.Request) string {
	userID := req.Header.Get("X-Caller-Id")
	return userID
}

type itemControllerInterface interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
	Delete(c *gin.Context)
	Buy(c *gin.Context)
	Update(c *gin.Context)
	SellerView(c *gin.Context)
	SearchByName(c *gin.Context)
}
type itemControllerStruct struct {
}

//Create : To create new Item
func (i *itemControllerStruct) Create(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err)
		return
	}
	createItem := &itemspb.CreateItemRequest{}
	if err := c.ShouldBindJSON(createItem); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	userID := getUserID(c.Request)
	createItem.Seller = userID
	result, createErr := connect.Client.Create(context.Background(), createItem)
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
	userID := getUserID(c.Request)
	buyRequest := &itemspb.BuyItemRequest{ItemID: itemID, UserID: userID}
	item, err := connect.Client.Buy(context.Background(), buyRequest)
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
	res, err := connect.Client.Get(context.Background(), toFind)
	if err != nil {
		c.JSON(http.StatusFound, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

//Search : To search for the items with given ID
func (i *itemControllerStruct) Delete(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err)
		return
	}
	itemID := c.Param("id")
	userID := getUserID(c.Request)
	ItemDetail := &itemspb.DeleteItemRequest{
		ItemID: itemID,
		UserID: userID,
	}
	res, err := connect.Client.Delete(context.Background(), ItemDetail)
	if err != nil {
		c.JSON(http.StatusFound, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

//Update : To update the given item
func (i *itemControllerStruct) Update(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err)
		return
	}
	userID := oauth.GetCallerID(c.Request)
	itemID := c.Param("id")
	itemToBeUpdated := &itemspb.UpdateItemRequest{
		UserID: userID,
		ItemID: itemID,
	}
	err := c.ShouldBind(itemToBeUpdated)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error Fetching The Item Details")
		return
	}
	fmt.Println("The value of items is", itemToBeUpdated)
	response, err := connect.Client.Update(context.Background(), itemToBeUpdated)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error While Updating The Item")
		return
	}
	c.JSON(http.StatusAccepted, response)
	return
}

//SellerView : To give seller the detail information about the product he is selling
func (i *itemControllerStruct) SellerView(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err)
		return
	}
	userID := oauth.GetCallerID(c.Request)
	itemID := c.Param("id")
	viewRequest := &itemspb.SellerViewRequest{
		UserID: userID,
		ItemID: itemID,
	}
	view, err := connect.Client.SellerView(context.Background(), viewRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error while fetching the database ")
		return
	}
	c.JSON(http.StatusAccepted, view)
	return
}

//SearchByName : To search the items by name
func (i *itemControllerStruct) SearchByName(c *gin.Context) {
	itemName := c.Param("itemName")
	if len(itemName) == 0 {
		c.JSON(http.StatusBadRequest, "Please Enter The valid Item Name")
		return
	}
	searchItemRequest := &itemspb.SearchItemRequest{
		Name: itemName,
	}
	items, err := connect.Client.SearchItem(context.Background(), searchItemRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error While Fetching The Items")
		return
	}
	for {
		details, err := items.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("error while fetching the item with given Details", err)
		}
		c.JSON(http.StatusAccepted, details.Item)

	}

}
