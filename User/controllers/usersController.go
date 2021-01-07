package controllers

import (
	"net/http"
	"strconv"

	"github.com/ankitanwar/GoAPIUtils/errors"
	bookoauth "github.com/ankitanwar/bookStore-OAuth/oAuth"
	oauth "github.com/ankitanwar/e-Commerce/Middleware/oAuth"
	"github.com/ankitanwar/e-Commerce/User/domain/users"
	"github.com/ankitanwar/e-Commerce/User/services"
	"github.com/gin-gonic/gin"
)

func getUserid(userIDParam string) (int, *errors.RestError) {
	userID, userErr := strconv.Atoi(userIDParam)
	if userErr != nil {
		err := errors.NewBadRequest("Enter the valid used id")
		return 0, err
	}
	return userID, nil
}

//CreateUser : To create the user
func CreateUser(c *gin.Context) {
	var newUser users.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		err := errors.NewBadRequest("Invalid Request")
		c.JSON(err.Status, err)
		return
	}

	result, saverr := services.UserServices.CreateUser(newUser)
	if saverr != nil {
		c.JSON(saverr.Status, saverr)
		return
	}
	c.JSON(http.StatusCreated, result.MarshallUser(oauth.IsPublic(c.Request)))
}

//GetUser : To get the user from the database
func GetUser(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err)
		return
	}
	userid, userErr := getUserid(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}
	user, err := services.UserServices.GetUser(userid)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, user.MarshallUser(bookoauth.IsPublic(c.Request)))

}

//UpdateUser :To Update the value of particaular user
func UpdateUser(c *gin.Context) {
	var user = users.User{}
	userid, userErr := getUserid(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}
	user.ID = userid
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequest("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	isPartial := c.Request.Method == http.MethodPatch

	updatedUser, err := services.UserServices.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, updatedUser.MarshallUser(bookoauth.IsPublic(c.Request)))
}

//DeleteUser :To Delete the user with given id
func DeleteUser(c *gin.Context) {
	userid, userErr := getUserid(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}
	if err := services.UserServices.DeleteUser(userid); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"Status": "User Deleted"})
}

//FindByStatus : To find all the users by given status
func FindByStatus(c *gin.Context) {
	status := c.Query("status")
	users, err := services.UserServices.FindByStatus(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, users.MarshallUser(bookoauth.IsPublic(c.Request)))

}

//Login : to verify user email and password
func Login(c *gin.Context) {
	verifyUser := users.LoginRequest{}
	if err := c.ShouldBindJSON(&verifyUser); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	user, err := services.UserServices.LoginUser(verifyUser)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, user.MarshallUser(bookoauth.IsPublic(c.Request)))

}

//GetCart : To get the items in the cart
func GetCart(c *gin.Context) {
	err := oauth.AuthenticateRequest(c.Request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	userID, err := getUserid(c.Param("userID"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	products, viewError := services.UserCart.ViewCart(userID)
	if viewError != nil {
		c.JSON(http.StatusInternalServerError, viewError)
		return
	}
	c.JSON(http.StatusAccepted, products)
}

//AddToCart : Adding item to the cart
func AddToCart(c *gin.Context) {
	err := oauth.AuthenticateRequest(c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	userID, err := getUserid(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	itemID := 0
	price := 100
	itemTitle := "testing"
	addError := services.UserCart.AddToCart(userID, itemID, price, itemTitle)
	if addError != nil {
		c.JSON(http.StatusInternalServerError, addError)
		return
	}
	c.JSON(http.StatusAccepted, "Item Has been added Successfully")

}

//DeleteFromCart : To delete the item from the user cart
func DeleteFromCart(c *gin.Context) {
	err := oauth.AuthenticateRequest(c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	userID, err := getUserid(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	itemID := 0
	delError := services.UserCart.RemoveFromCart(userID, itemID)
	if delError != nil {
		c.JSON(http.StatusInternalServerError, delError)
		return
	}
	c.JSON(http.StatusAccepted, "Item Has been deleted Successfully")

}

//GetAddress : To Get the address of the given user
func GetAddress(c *gin.Context) {
	err := oauth.AuthenticateRequest(c.Request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	userID, err := getUserid(c.Param("userID"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	address, err := services.UserServices.GetAddress(userID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusAccepted, address)
}

//AddAddress : To Get the address of the given user
func AddAddress(c *gin.Context) {
	err := oauth.AuthenticateRequest(c.Request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	address := &users.UserAddress{}
	bindErr := c.ShouldBindJSON(address)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, "Error while binding to the json")
		return
	}
	userID, err := getUserid(c.Param("userID"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	err = services.UserServices.AddAddress(userID, *address)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusAccepted, address)
}
