package addressdb

import (
	"context"

	"github.com/ankitanwar/e-Commerce/User/domain/users"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	//NotFound : if the given data is not present in the database
	NotFound = "mongo: no documents in result"
)

func createNewAddress(userID string, address *users.UserAddress) error {
	addAddress := &users.Address{}
	addAddress.UserID = userID
	addAddress.List = append(addAddress.List, *address)
	_, err := collection.InsertOne(context.Background(), addAddress)
	if err != nil {
		return err
	}
	return nil

}

//GetUserAddress : To get the all given address from the databse
func GetUserAddress(userID string) (*users.Address, error) {
	address := &users.Address{}
	filter := bson.M{"_id": userID}
	err := collection.FindOne(context.Background(), filter).Decode(address)
	if err != nil {
		return nil, err
	}
	return address, nil

}

//AddAddress : To add the given address into the database
func AddAddress(userID string, address *users.UserAddress) error {
	storedAddress := &users.Address{}
	filter := bson.M{"_id": userID}
	err := collection.FindOne(context.Background(), filter).Decode(storedAddress)
	if err != nil {
		if err.Error() == NotFound {
			err = createNewAddress(userID, address)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	addAddress := bson.M{"$push": bson.M{"addresses": address}}
	_, err = collection.UpdateOne(context.Background(), filter, addAddress)
	if err != nil {
		return err
	}
	return nil
}

//RemoveAddress : To remove the given item from the address
func RemoveAddress(userID string) {

}
