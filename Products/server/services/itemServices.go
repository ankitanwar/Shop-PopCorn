package services

import (
	"context"
	"errors"
	"time"

	itemspb "github.com/ankitanwar/e-Commerce/Products/proto"
	db "github.com/ankitanwar/e-Commerce/Products/server/database"
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
	item.QuantitySold = 0
	item.Status = "Available"
	res, err := db.SaveItem(item)
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
			QuantitySold:      0,
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
	item, err := db.SearchByID(oid)
	if err != nil {
		return nil, err
	}
	response := &itemspb.GetItemResposne{}
	response.Item = item
	return response, nil

}

//Update : To update the item by particular ID
func (s *ItemService) Update(ctx context.Context, req *itemspb.UpdateItemRequest) (*itemspb.UpdateItemResponse, error) {
	itemID := req.GetID()
	oid, err := primitive.ObjectIDFromHex(itemID)
	if err != nil {
		return nil, err
	}
	updatedDetial := req.GetItem()
	savedDetail, err := db.DetailView(oid)
	if err != nil {
		return nil, err
	}
	if savedDetail.Seller != updatedDetial.Seller {
		return nil, errors.New("Item Not Found")
	}
	if updatedDetial.Title != "" {
		savedDetail.Title = updatedDetial.Title
	}
	if updatedDetial.Description != "" {
		savedDetail.Description = updatedDetial.Description
	}
	if updatedDetial.AvailableQuantity != 0 {
		savedDetail.AvailableQuantity += updatedDetial.AvailableQuantity
	}
	if updatedDetial.Status != "" {
		savedDetail.Status = updatedDetial.Status
	}
	err = db.UpdateItem(savedDetail)
	if err != nil {
		return nil, err
	}
	response := &itemspb.UpdateItemResponse{}
	response.Item = savedDetail
	return response, nil

}

//Delete : To delete the item by particular ID
func (s *ItemService) Delete(ctx context.Context, req *itemspb.DeleteItemRequest) (*itemspb.DeleteItemResponse, error) {
	oid, _ := primitive.ObjectIDFromHex(req.GetId())
	err := db.DeleteByID(oid)
	if err != nil {
		return nil, err
	}
	response := &itemspb.DeleteItemResponse{}
	response.Operation = "Item Has Been Deleted Successfully"
	return response, nil
}

//Buy : To buy the given item
func (s *ItemService) Buy(c context.Context, req *itemspb.BuyItemRequest) (*itemspb.BuyItemResponse, error) {
	itemID := req.GetItemID()
	// userID := req.GetUserID()
	oid, err := primitive.ObjectIDFromHex(itemID)
	item, err := db.BuyItem(oid)
	if err != nil {
		return nil, err
	}
	if item.AvailableQuantity <= 0 {
		return nil, errors.New("Item OF Stock")
	}
	item.AvailableQuantity--
	item.SoldQuantity++
	err = db.UpdateQunatity(item)
	if err != nil {
		return nil, errors.New("Error While Buying The Item")
	}
	// address, addressErr := user.GetUserAddress.GetAddress(userID)
	// if addressErr != nil {
	// 	return nil, errors.New("Invalid address")
	// }
	deliveryTime := time.Now()
	deliveryTime.Format("01-02-2006")
	deliveryTime = deliveryTime.AddDate(0, 0, 10)
	response := &itemspb.BuyItemResponse{}
	response.ExceptedDateOfDilvery = deliveryTime.String()
	response.Title = item.Title
	response.Price = item.Price
	return response, nil

}

//SellerView : To Give seller detail view about the item they are selling
func (s *ItemService) SellerView(c context.Context, view *itemspb.SellerViewRequest) (*itemspb.SellerViewRespsonse, error) {
	return nil, nil
}

//SearchItem : To steam the items available with given name
func (s *ItemService) SearchItem(req *itemspb.SearchItemRequest, stream itemspb.ItemService_SearchItemServer) error {
	return nil
}
