package addressdb

import (
	"context"
	"fmt"

	"github.com/ankitanwar/Shop-PopCorn/User/domain/users"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	//NotFound : if the given data is not present in the database
	NotFound = "mongo: no documents in result"
)

//GetUserAddress : To get the all given address from the databse
func GetUserAddress(userID string) (*users.Address, error) {
	address := &users.Address{}
	filter := bson.M{"_id": userID}
	err := collection.FindOne(context.Background(), filter).Decode(address)
	if err != nil {
		fmt.Println("The value of databse err is", err)
		return nil, err
	}
	return address, nil

}

//AddAddress : To add the given address into the database
func AddAddress(userID string, address *users.UserAddress) error {
	opts := options.Update().SetUpsert(true)
	filter := bson.M{"_id": userID}
	addAddress := bson.M{"$push": bson.M{"addresses": address}}
	_, err := collection.UpdateOne(context.Background(), filter, addAddress, opts)
	if err != nil {
		return err
	}
	return nil
}

//RemoveAddress : To remove the given item from the address
func RemoveAddress(userID string, addressID string) error {
	filter := bson.M{"_id": userID}
	remove := bson.M{"$pull": bson.M{"addresses": bson.M{"id": addressID}}}
	_, err := collection.UpdateOne(context.Background(), filter, remove)
	if err != nil {
		return err
	}
	return nil
}
