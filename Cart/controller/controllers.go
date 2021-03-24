package cotrollers

import (
	"net/http"

	"github.com/ankitanwar/GoAPIUtils/errors"
	"github.com/ankitanwar/Shop-PopCorn/Cart/services"
	product "github.com/ankitanwar/Shop-PopCorn/Middleware/Products"
	oauth "github.com/ankitanwar/Shop-PopCorn/Middleware/oAuth"
	"github.com/ankitanwar/Shop-PopCorn/Middleware/user"
	"github.com/gin-gonic/gin"
)

func getUserAddress(req *http.Request, addressID string) (*user.SpecificAddress, *errors.RestError) {
	address, addressErr := user.GetUserAddress.GetAddress(req, addressID)
	if addressErr != nil {
		return nil, addressErr
	}
	return address, nil
}

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
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err.Message)
		return
	}
	userID := getCallerID(c.Request)
	itemInCart, err := services.ViewCart(userID)
	if err != nil {
		c.JSON(err.Status, err.Message)
		return
	}
	c.JSON(http.StatusAccepted, itemInCart)
	return
}

//Checkout : To checkout all the items from the cart
func Checkout(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err.Message)
		return
	}
	userID := getCallerID(c.Request)
	addressID := c.Param("addressID")
	address, addressErr := getUserAddress(c.Request, addressID)
	if addressErr != nil {
		c.JSON(addressErr.Status, addressErr.Message)
		return
	}
	response, err := services.Checkout(c.Request, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Unable To Checkout The cart")
		return
	}
	response.Country = address.Country
	response.HouseNumber = address.HouseNumber
	response.Street = address.Street
	response.Phone = address.Phone
	response.State = address.State
	c.JSON(http.StatusAccepted, response)

}
