package cartdatabase

import (
	"context"
	"fmt"

	"github.com/ankitanwar/GoAPIUtils/errors"
	domin "github.com/ankitanwar/e-Commerce/Cart/domain"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	//NotFound : if the given data is not present in the database
	NotFound = "mongo: no documents in result"
)

func createNew(userID string, item domin.Item) error {
	filter := bson.A{"_id", userID, "items", item}
	_, err := Collection.InsertOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}

//AddToCart : To add the item into the cart
func AddToCart(userID string, item domin.Item) *errors.RestError {
	filter := bson.D{{"_id", userID}}
	findErr := Collection.FindOne(context.Background(), filter)
	fmt.Println(findErr)
	return nil
}
