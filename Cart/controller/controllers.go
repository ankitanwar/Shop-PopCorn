package cotrollers

import (
	"fmt"
	"net/http"

	"github.com/ankitanwar/GoAPIUtils/errors"
	product "github.com/ankitanwar/e-Commerce/Middleware/Products"
	oauth "github.com/ankitanwar/e-Commerce/Middleware/oAuth"
	"github.com/gin-gonic/gin"
)

func getCallerID(request *http.Request) string {
	if request == nil {
		return ""
	}
	return request.Header.Get("X-Caller-Id")
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
	id, getItemErr := getItemID(c.Param("itemID"))
	if getItemErr != nil {
		c.JSON(getItemErr.Status, getItemErr.Message)
		return
	}
	details, err := product.ItemSerivce.GetItemDetails(id)
	if err != nil {
		c.JSON(err.Status, err.Message)
		return
	}
	fmt.Println(details)

}

//RemoveFromCart : To remove the given item from the cart
func RemoveFromCart(c *gin.Context) {

}

//ViewCart : To view the cart of the particular user
func ViewCart(c *gin.Context) {

}

//Checkout : To checkout all the items from the cart
func Checkout(c *gin.Context) {

}
