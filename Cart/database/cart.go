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
	t := &domin.User{}
	t.UserID = userID
	t.Items = append(t.Items, item)
	res, err := collection.InsertOne(context.Background(), t)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

//AddToCart : To add the item into the cart
func AddToCart(userID string, item domin.Item) *errors.RestError {
	user := &domin.User{}
	filter := bson.M{"_id": userID}
	err := collection.FindOne(context.Background(), filter).Decode(user)
	if err != nil {
		if err.Error() == NotFound {
			err = createNew(userID, item)
			if err != nil {
				return errors.NewInternalServerError("Error while adding the item into the cart")
			}
			return nil
		}
		return errors.NewInternalServerError("Error while adding the item into the cart")
	}
	PushToCart := bson.M{"$push": bson.M{"items": item}}
	_, err = collection.UpdateOne(context.Background(), filter, PushToCart)
	if err != nil {
		return errors.NewInternalServerError("Error While Adding Item into the cart")
	}
	return nil
}

//RemoveFromCart : To remove the item from the cart
func RemoveFromCart(userID, itemID string) error {
	filter := bson.M{"_id": userID}
	remove := bson.M{"$pull": bson.M{"items": bson.M{"$in": bson.A{bson.M{"itemID": itemID}}}}}
	_, err := collection.UpdateOne(context.Background(), filter, remove)
	fmt.Println("The value of err is", err)
	if err != nil {
		return err
	}
	return nil
}
