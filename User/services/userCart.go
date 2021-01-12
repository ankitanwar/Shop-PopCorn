package services

import (
	"context"
	"errors"
	"fmt"

	mongo "github.com/ankitanwar/e-Commerce/User/databasource/mongoUserCart"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	//NotFound : If the id is not present in the database
	NotFound = "mongo: no documents in result"
)

//Order : Details of the orders
type Order struct {
	userID   int             `bson:"user_id"`
	products []productDetail `bson:"products"`
}

type productDetail struct {
	itemID string `bson:"itemID"`
	title  string `bson:"title"`
	price  int    `bson:"price"`
}

var (
	//UserCart : Services Available fot user cart
	UserCart userCartInerface = &userCartServices{}
)

type userCartServices struct {
}
type userCartInerface interface {
	ViewCart(int) ([]productDetail, error)
	AddToCart(string, int, int, string) error
	RemoveFromCart(int, string) error
}

func (u *userCartServices) ViewCart(userID int) ([]productDetail, error) {
	filter := bson.M{"user_id": userID}
	orders := &Order{}
	err := mongo.Collection.FindOne(context.Background(), filter).Decode(orders)
	if err != nil {
		return nil, err
	}
	return orders.products, nil

}

func (u *userCartServices) AddToCart(itemID string, userID, price int, name string) error {
	fmt.Println("Inside function", userID, itemID, price, name)
	filter := bson.M{"user_id": userID}
	orders := &Order{}
	err := mongo.Collection.FindOne(context.Background(), filter).Decode(orders)
	if err != nil {
		if err.Error() == NotFound {
			orders.userID = userID
			t := productDetail{
				itemID: itemID,
				title:  name,
				price:  price,
			}
			orders.products = append(orders.products, t)
			fmt.Println("The value of order is ", orders)
			_, err := mongo.Collection.InsertOne(context.Background(), orders)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}

	return nil

}

func (u *userCartServices) RemoveFromCart(userID int, itemID string) error {
	filter := bson.M{"user_id": userID}
	orders := &Order{}
	err := mongo.Collection.FindOne(context.Background(), filter).Decode(orders)
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
