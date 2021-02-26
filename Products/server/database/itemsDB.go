package db

import (
	"context"

	itemspb "github.com/ankitanwar/e-Commerce/Products/proto"
	domain "github.com/ankitanwar/e-Commerce/Products/server/domian/items"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//SaveItem : To save the given item into the database
func SaveItem(item *itemspb.Item) (*mongo.InsertOneResult, error) {
	res, err := collection.InsertOne(context.Background(), item)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//SearchByID : To search the given item by ID
func SearchByID(id primitive.ObjectID) (*itemspb.ViewItem, error) {
	filter := bson.M{"_id": id}
	item := &itemspb.ViewItem{}
	err := collection.FindOne(context.Background(), filter).Decode(item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

//DeleteByID : To delete the item by given ID
func DeleteByID(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}

//BuyItem : To buy the given item
func BuyItem(id primitive.ObjectID) (*domain.Item, error) {
	item := &domain.Item{}
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.Background(), filter).Decode(item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

//UpdateItem :To update the given item
func UpdateItem(item *itemspb.Item) error {
	id := item.ID
	filter := bson.M{"_id": id}
	_, err := collection.UpdateOne(context.Background(), filter, item)
	if err != nil {
		return err
	}
	return nil
}

//DetailView : To give the detail view about the item
func DetailView(itemID primitive.ObjectID) (*itemspb.Item, error) {
	filter := bson.M{"_id": itemID}
	item := &itemspb.Item{}
	err := collection.FindOne(context.Background(), filter).Decode(item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

//UpdateQunatity : To update the qunatity of the given item
func UpdateQunatity(item *domain.Item) error {
	itemID := item.ID.String()
	filter := bson.M{"_id": itemID}
	update := bson.M{"$set": bson.A{"AvailableQuantity", item.AvailableQuantity, "SoldQuantity", item.SoldQuantity}}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil

}
