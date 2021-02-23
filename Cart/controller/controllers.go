package cotrollers

import (
	"net/http"

	"github.com/ankitanwar/GoAPIUtils/errors"
	"github.com/ankitanwar/e-Commerce/Cart/services"
	product "github.com/ankitanwar/e-Commerce/Middleware/Products"
	oauth "github.com/ankitanwar/e-Commerce/Middleware/oAuth"
	"github.com/gin-gonic/gin"
)

func getCallerID(request *http.Request) string {
	userID := request.Header.Get("X-Caller-Id")
	return userID
}

func getItemID(itemID string) (string, *errors.RestError) {
	if itemID == "" {
		return "", errors.NewBadRequest("Invalid Item ID")
	}
	return itemID, nil
}

//AddToCart : To add the given item into the cart
func AddToCart(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err)
		return
	}
	userID := getCallerID(c.Request)
	itemID, getItemErr := getItemID(c.Param("itemID"))
	if getItemErr != nil {
		c.JSON(getItemErr.Status, getItemErr.Message)
		return
	}
	details, err := product.ItemSerivce.GetItemDetails(itemID)
	if err != nil {
		c.JSON(err.Status, err.Message)
		return
	}
	err = services.AddToCart(userID, itemID, details)
	if err != nil {
		c.JSON(err.Status, err.Message)
		return
	}
	c.JSON(http.StatusAccepted, "Item has been added to the cart successfully")
	return

}

//RemoveFromCart : To remove the given item from the cart
func RemoveFromCart(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err.Message)
		return
	}
	userID := getCallerID(c.Request)
	itemID, err := getItemID(c.Param("itemID"))
	if err != nil {
		c.JSON(err.Status, err.Message)
	}
	err = services.RemoveFromCart(userID, itemID)
	if err != nil {
		c.JSON(err.Status, err.Message)
		return
	}
	c.JSON(http.StatusAccepted, "Item has been removed successfully")
	return
}

//ViewCart : To view the cart of the particular user
func ViewCart(c *gin.Context) {

}

//Checkout : To checkout all the items from the cart
func Checkout(c *gin.Context) {

}
