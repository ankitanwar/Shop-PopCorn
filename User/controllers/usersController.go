package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ankitanwar/GoAPIUtils/errors"
	bookoauth "github.com/ankitanwar/bookStore-OAuth/oAuth"
	"github.com/ankitanwar/e-Commerce/User/domain/users"
	"github.com/ankitanwar/e-Commerce/User/services"
	oauth "github.com/ankitanwar/e-Commerce/interactWithOAuth/oAuth"
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
		fmt.Println("This is working")
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
	fmt.Println("id and password at context is", verifyUser.Password, verifyUser.Email)
	user, err := services.UserServices.LoginUser(verifyUser)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, user.MarshallUser(bookoauth.IsPublic(c.Request)))

}
