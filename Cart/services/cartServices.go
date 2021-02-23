package services

import (
	"github.com/ankitanwar/GoAPIUtils/errors"
	cartdatabase "github.com/ankitanwar/e-Commerce/Cart/database"
	domain "github.com/ankitanwar/e-Commerce/Cart/domain"
	product "github.com/ankitanwar/e-Commerce/Middleware/Products"
)

//AddToCart : To add the given product details into the cart of the given user
func AddToCart(userID, itemID string, productDetails *product.ItemValue) *errors.RestError {
	user := &domain.User{}
	user.UserID = userID
	item := &domain.Item{}
	item.ItemID = itemID
	item.Price = productDetails.Price
	item.Title = productDetails.Title
	saveToDB := cartdatabase.AddToCart(userID, *item)
	if saveToDB != nil {
		return saveToDB
	}
	return nil

}

//RemoveFromCart : To remove the given item from the cart
func RemoveFromCart(userID, itemID string) *errors.RestError {
	removeErr := cartdatabase.RemoveFromCart(userID, itemID)
	if removeErr != nil {
		return errors.NewInternalServerError("Error while removing the item from the cart")
	}
	return nil
}
