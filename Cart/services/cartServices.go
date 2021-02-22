package services

import (
	domain "github.com/ankitanwar/e-Commerce/Cart/domain"
	product "github.com/ankitanwar/e-Commerce/Middleware/Products"
)

//AddToCart : To add the given product details into the cart of the given user
func AddToCart(userID, itemID string, productDetails product.ItemValue) {
	user := &domain.User{}
	user.UserID = userID
	item := &domain.Item{}
	item.ItemID = itemID
	item.Price = productDetails.Price
	item.Title = productDetails.Title
}
