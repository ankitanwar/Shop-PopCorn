package mongo

import (
	"context"

	"github.com/ankitanwar/Shop-PopCorn/Oauth/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateAccessToken(accessToken *domain.AccessToken) error {
	opts := options.Update().SetUpsert(true)
	filter := bson.M{"_id": accessToken.UserID, "email": accessToken.Email}
	update := bson.M{"$set": bson.M{"access_token": accessToken.Token}}
	_, err := Collection.UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
		return err
	}
	return nil

}

func GetAccessToken(userID string) (*domain.AccessToken, error) {
	filter := bson.M{"_id": userID}
	details := &domain.AccessToken{}
	err := Collection.FindOne(context.Background(), filter).Decode(details)
	if err != nil {
		return nil, err
	}
	return details, nil
}

func RemoveAccessToken(userID string) error {
	filter := bson.M{"_id": userID}
	_, err := Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}
