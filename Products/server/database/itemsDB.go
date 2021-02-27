package db

import (
	"context"

	domain "github.com/ankitanwar/e-Commerce/Products/server/domian/items"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//SaveItem : To save the given item into the database
func SaveItem(item *domain.Item) (*mongo.InsertOneResult, error) {
	res, err := collection.InsertOne(context.Background(), item)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//SearchByID : To search the given item by ID
func SearchByID(itemID primitive.ObjectID) (*domain.Item, error) {
	filter := bson.M{"_id": itemID}
	item := &domain.Item{}
	err := collection.FindOne(context.Background(), filter).Decode(item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

//DeleteByID : To delete the item by given ID
func DeleteByID(itemID primitive.ObjectID) error {
	filter := bson.M{"_id": itemID}
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}

//BuyItem : To buy the given item
func BuyItem(itemID primitive.ObjectID) (*domain.Item, error) {
	item := &domain.Item{}
	filter := bson.M{"_id": itemID}
	err := collection.FindOne(context.Background(), filter).Decode(item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

//UpdateItem :To update the given item
func UpdateItem(item *domain.Item) error {
	filter := bson.M{"_id": item.ID}
	update := bson.M{"$set": bson.M{"name": item.Name, "desciption": item.Description, "availablequantity": item.AvailableQuantity, "status": item.Status}}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

//DetailView : To give the detail view about the item
func DetailView(itemID primitive.ObjectID) (*domain.Item, error) {
	filter := bson.M{"_id": itemID}
	item := &domain.Item{}
	err := collection.FindOne(context.Background(), filter).Decode(item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

//UpdateQunatity : To update the qunatity of the given item
func UpdateQunatity(item *domain.Item) error {
	filter := bson.M{"_id": item.ID}
	update := bson.M{"$set": bson.M{"availablequantity": item.AvailableQuantity, "quantitysold": item.SoldQuantity}}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil

}
