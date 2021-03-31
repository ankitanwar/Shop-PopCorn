package services

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ankitanwar/GoAPIUtils/errors"
	cartdatabase "github.com/ankitanwar/Shop-PopCorn/Cart/database"
	domain "github.com/ankitanwar/Shop-PopCorn/Cart/domain"
	product "github.com/ankitanwar/Shop-PopCorn/Middleware/Products"
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

//ViewCart : To view all the items in the cart
func ViewCart(userID string) (*[]domain.Item, *errors.RestError) {
	userCart, err := cartdatabase.ViewCart(userID)
	if err != nil {
		return nil, errors.NewInternalServerError("Error while fetching the cart")
	}
	return &userCart.Items, nil
}

//Checkout : To checkout all the given items in the cart
func Checkout(req *http.Request, userID string) (*domain.CheckoutResponse, *errors.RestError) {
	cart, err := cartdatabase.Checkout(userID)
	if err != nil {
		return nil, errors.NewInternalServerError("Unable To Fetch Cart Items")
	}
	response := &domain.CheckoutResponse{}
	deliveryTime := time.Now()
	deliveryTime.Format("01-02-2006")
	deliveryTime = deliveryTime.AddDate(0, 0, 10)
	i := 0
	for i < len(cart.Items) {
		currentItem := cart.Items[0]
		itemID := currentItem.ItemID
		Buyerr := product.ItemSerivce.BuyItem(req, itemID)
		if Buyerr == nil {
			details := &domain.ByResponse{
				Price:       int64(currentItem.Price),
				DeliverDate: deliveryTime.String(),
				Title:       currentItem.Title,
			}
			response.Products = append(response.Products, *details)
			response.TotalCost += int64(currentItem.Price)
			fmt.Println("The value of response is", response)
		} else {
			return nil, Buyerr
		}
		i++
	}
	return response, nil

}
