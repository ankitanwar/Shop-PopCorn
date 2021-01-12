package items

import (
	"context"
	"fmt"

	"github.com/ankitanwar/GoAPIUtils/errors"
	"github.com/ankitanwar/e-Commerce/Middleware/user"
	itemspb "github.com/ankitanwar/e-Commerce/Products/proto"
	db "github.com/ankitanwar/e-Commerce/Products/server/database"
)

var (
	//UserService : Services available for user struct
	UserService userInterface = &userStruct{}
)

type userStruct struct{}

type userInterface interface {
	SaveOrder(string, string, string, int) error
}

type sales struct {
	UserID      string `bson:"userID"`
	ProductID   string `bson:"productID"`
	Description string `bson:description`
	Price       int    `bson:"price"`
}

//UserHistory : To keep track of all the items the user has ordered
type UserHistory struct {
	UserID int            `bson:"userID"`
	orders []itemspb.Item `bson:"orders"`
}

func getAddress(userID string) (*user.Address, *errors.RestError) {
	add, err := user.GetUserAddress.GetAddress(userID)
	if err != nil {
		return nil, err
	}
	return add, nil
}

func (u *userStruct) SaveOrder(userID, productID, description string, price int) error {
	s := &sales{}
	s.UserID = userID
	s.ProductID = productID
	s.Description = description
	s.Price = price
	fmt.Println("The value of s is ", s)
	_, err := db.Sales.InsertOne(context.Background(), s)
	if err != nil {
		return err
	}
	return nil

}
