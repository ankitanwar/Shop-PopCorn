package items

import (
	"context"
	"fmt"
	"time"

	itemspb "github.com/ankitanwar/e-Commerce/Products/proto"
	db "github.com/ankitanwar/e-Commerce/Products/server/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	//Services : services available
	Services itemspb.ItemServiceServer = &ItemService{}
)

//ItemService : Services Available for items
type ItemService struct {
}

//Create : To Create the item
func (s *ItemService) Create(ctx context.Context, req *itemspb.CreateItemRequest) (*itemspb.CreateItemResposne, error) {
	item := req.GetItem()
	data := Item{
		Seller:            item.GetSeller(),
		Title:             item.GetTitle(),
		Description:       item.GetDescription(),
		Price:             item.GetPrice(),
		AvailableQuantity: item.GetAvailableQuantity(),
		Status:            item.GetStatus(),
	}
	res, err := db.Collection.InsertOne(context.Background(), data)
	if err != nil {
		return nil, err
	}
	oid, _ := res.InsertedID.(primitive.ObjectID)
	return &itemspb.CreateItemResposne{
		Item: &itemspb.Item{
			ID:                oid.Hex(),
			Seller:            item.GetSeller(),
			Title:             item.GetTitle(),
			Description:       item.GetDescription(),
			Price:             item.GetPrice(),
			AvailableQuantity: item.GetAvailableQuantity(),
			Status:            item.GetStatus(),
		},
	}, nil

}

//Get : To Get The Item By Particular Id
func (s *ItemService) Get(ctx context.Context, req *itemspb.GetItemRequest) (*itemspb.GetItemResposne, error) {
	itemID := req.GetID()
	oid, err := primitive.ObjectIDFromHex(itemID)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": oid}
	res := &itemspb.Item{}
	findErr := db.Collection.FindOne(context.Background(), filter).Decode(res)
	if findErr != nil {
		return nil, err
	}
	return &itemspb.GetItemResposne{
		Item: res,
	}, nil

}

//Update : To update the item by particular ID
func (s *ItemService) Update(ctx context.Context, req *itemspb.UpdateItemRequest) (*itemspb.UpdateItemResponse, error) {
	itemID := req.GetID()
	oid, err := primitive.ObjectIDFromHex(itemID)
	if err != nil {
		return nil, err
	}
	item := req.GetItem()
	data := Item{
		Seller:            item.GetSeller(),
		Title:             item.GetTitle(),
		Description:       item.GetDescription(),
		Price:             item.GetPrice(),
		AvailableQuantity: item.GetAvailableQuantity(),
		Status:            item.GetStatus(),
	}
	fmt.Println(oid, data)
	return nil, nil
}

//Delete : To delete the item by particular ID
func (s *ItemService) Delete(ctx context.Context, req *itemspb.DeleteItemRequest) (*itemspb.DeleteItemResponse, error) {
	oid, _ := primitive.ObjectIDFromHex(req.GetId())
	filter := bson.M{"id": oid}
	res, err := db.Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	if res.DeletedCount == 0 {
		return nil, err
	}
	return &itemspb.DeleteItemResponse{
		Operation: req.GetId(),
	}, nil
}

//Buy : To buy the given item
func (s *ItemService) Buy(c context.Context, req *itemspb.BuyItemRequest) (*itemspb.BuyItemResponse, error) {
	itemID := req.GetItemID()
	oid, err := primitive.ObjectIDFromHex(itemID)
	userID := req.GetUserID()
	filter := bson.M{"_id": oid}
	result := &itemspb.Item{}
	err = db.Collection.FindOne(context.Background(), filter).Decode(result)
	fmt.Println("The value of result is ", result)
	if err != nil {
		fmt.Println("find err ")
		return nil, err
	}
	deliveryTime := time.Now()
	deliveryTime.Format("01-02-2006")
	deliveryTime = deliveryTime.AddDate(0, 0, 10)
	err = UserService.SaveOrder(userID, itemID, result.Description, int(result.Price))
	return &itemspb.BuyItemResponse{
		ExceptedDateOfDilvery: deliveryTime.String(),
		Title:                 result.Title,
		Price:                 result.Price,
	}, nil
}
