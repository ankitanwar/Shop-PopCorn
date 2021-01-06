package services

import (
	"context"
	"errors"

	cart "github.com/ankitanwar/e-Commerce/User/databasource/mongoUserCart"
	"go.mongodb.org/mongo-driver/bson"
)

//Order : Details of the orders
type Order struct {
	userID   int             `bson:"user_id"`
	products []productDetail `bson:"products"`
}

type productDetail struct {
	itemID int
	title  string
	price  int
}

var (
	//UserCart : Services Available fot user cart
	UserCart userCartInerface = &userCartServices{}
)

type userCartServices struct {
}
type userCartInerface interface {
	ViewCart(int) ([]productDetail, error)
	Checkout(int)
	AddToCart(int, int, int, string) error
	RemoveFromCart(int, int) error
}

func (u *userCartServices) ViewCart(userID int) ([]productDetail, error) {
	filter := bson.M{"user_id": userID}
	orders := &Order{}
	err := cart.Collection.FindOne(context.Background(), filter).Decode(orders)
	if err != nil {
		return nil, err
	}
	return orders.products, nil

}

func (u *userCartServices) Checkout(userID int) {

}
func (u *userCartServices) AddToCart(userID, itemID, price int, name string) error {
	filter := bson.M{"user_id": userID}
	orders := &Order{}
	err := cart.Collection.FindOne(context.Background(), filter).Decode(orders)
	if err != nil {
		return err
	}
	t := productDetail{
		itemID: itemID,
		title:  name,
		price:  price,
	}
	orders.products = append(orders.products, t)
	_, updateError := cart.Collection.UpdateOne(context.Background(), filter, orders)
	if updateError != nil {
		return updateError
	}

	return nil

}

func (u *userCartServices) RemoveFromCart(userID, itemID int) error {
	filter := bson.M{"user_id": userID}
	orders := &Order{}
	err := cart.Collection.FindOne(context.Background(), filter).Decode(orders)
	if err != nil {
		return err
	}
	var found = false
	for i := 0; i < len(orders.products); i++ {
		if orders.products[i].itemID == itemID {
			x := orders.products[:i]
			y := orders.products[i+1:]
			x = append(x, y...)
			orders.products = x
			found = true
			break
		}
	}
	if found == false {
		return errors.New("Item not found")
	}
	return nil

}
