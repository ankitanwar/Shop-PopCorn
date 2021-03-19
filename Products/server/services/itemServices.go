package services

import (
	"context"
	"errors"
	"time"

	itemspb "github.com/ankitanwar/e-Commerce/Products/proto"
	db "github.com/ankitanwar/e-Commerce/Products/server/database"
	domain "github.com/ankitanwar/e-Commerce/Products/server/domian/items"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	//Services : services available
	Services itemspb.ItemServiceServer = &ItemService{}
)

//ItemService : Services Available for items
type ItemService struct {
}

func getOID(id string) (*primitive.ObjectID, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	if len(oid) == 0 {
		return nil, errors.New("Invalid oid")
	}
	return &oid, nil
}

//Create : To Create the item
func (s *ItemService) Create(ctx context.Context, req *itemspb.CreateItemRequest) (*itemspb.CreateItemResposne, error) {
	if len(req.Name) == 0 {
		return nil, errors.New("Please Enter The Valid Name")
	}
	item := &domain.Item{
		Seller:            req.Seller,
		Name:              req.Name,
		Description:       req.Description,
		Price:             req.Price,
		AvailableQuantity: req.AvailableQuantity,
		SoldQuantity:      0,
		Status:            "Available",
	}
	println("The value of createItem is", item)
	res, err := db.SaveItem(item)
	if err != nil {
		return nil, err
	}
	oid, _ := res.InsertedID.(primitive.ObjectID)
	response := &itemspb.CreateItemResposne{
		Id:                oid.String(),
		Seller:            item.Seller,
		Title:             item.Name,
		Description:       item.Description,
		Price:             item.Price,
		AvailableQuantity: item.AvailableQuantity,
		Status:            item.Status,
		QuantitySold:      item.SoldQuantity,
	}
	return response, nil

}

//Get : To Get The Item By Particular Id
func (s *ItemService) Get(ctx context.Context, req *itemspb.GetItemRequest) (*itemspb.GetItemResposne, error) {
	itemID := req.ID
	oid, err := getOID(itemID)
	if err != nil {
		return nil, errors.New("Invalid Item iD")
	}
	item, err := db.SearchByID(*oid)
	if err != nil {
		return nil, err
	}
	result := &itemspb.ViewItem{
		ID:    item.ID.String(),
		Title: item.Name,
		Price: item.Price,
	}
	response := &itemspb.GetItemResposne{
		Item: result,
	}
	return response, nil

}

//Update : To update the item by particular ID
func (s *ItemService) Update(ctx context.Context, req *itemspb.UpdateItemRequest) (*itemspb.UpdateItemResponse, error) {
	itemID := req.ItemID
	userID := req.UserID
	oid, err := getOID(itemID)
	if err != nil {
		return nil, errors.New("Invalid oid")
	}
	savedDetail, err := db.DetailView(*oid)
	if err != nil {
		return nil, err
	}
	if savedDetail.Seller != userID {
		return nil, errors.New("Item Not Found")
	}
	if req.Name != "" {
		savedDetail.Name = req.Name
	}
	if req.Description != "" {
		savedDetail.Description = req.Description
	}
	if req.AvailableQuantity != 0 {
		savedDetail.AvailableQuantity += req.AvailableQuantity
	}
	err = db.UpdateItem(savedDetail)
	if err != nil {
		return nil, err
	}
	response := &itemspb.UpdateItemResponse{
		ItemID:            itemID,
		Seller:            savedDetail.Seller,
		Title:             savedDetail.Name,
		Description:       savedDetail.Description,
		Price:             savedDetail.Price,
		AvailableQuantity: savedDetail.AvailableQuantity,
		Status:            savedDetail.Status,
		QuantitySold:      savedDetail.SoldQuantity,
	}
	return response, nil

}

//Delete : To delete the item by particular ID
func (s *ItemService) Delete(ctx context.Context, req *itemspb.DeleteItemRequest) (*itemspb.DeleteItemResponse, error) {
	itemID := req.ItemID
	userID := req.UserID
	oid, err := getOID(itemID)
	if err != nil {
		return nil, errors.New("Invalid oid")
	}
	savedItem, err := db.DetailView(*oid)
	if err != nil {
		return nil, errors.New("Item Not Found")
	}
	if savedItem.Seller != userID {
		return nil, errors.New("Item Not Found")
	}

	err = db.DeleteByID(*oid)
	if err != nil {
		return nil, errors.New("Error While Removing The Item From The Database")
	}
	response := &itemspb.DeleteItemResponse{}
	response.Message = "Item Has Been Deleted Successfully"
	return response, nil
}

//Buy : To buy the given item
func (s *ItemService) Buy(c context.Context, req *itemspb.BuyItemRequest) (*itemspb.BuyItemResponse, error) {
	itemID := req.ItemID
	// userID := req.GetUserID()
	oid, err := getOID(itemID)
	if err != nil {
		return nil, errors.New("Invalid oid")
	}
	item, err := db.BuyItem(*oid)
	if err != nil {
		return nil, errors.New("Error While Buying The Item")
	}
	if item.AvailableQuantity <= 0 {
		return nil, errors.New("Item Of Stock")
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
	response := &itemspb.BuyItemResponse{
		ExceptedDateOfDilvery: deliveryTime.String(),
		Title:                 item.Name,
		Price:                 item.Price,
	}
	return response, nil

}

//SellerView : To Give seller detail view about the item they are selling
func (s *ItemService) SellerView(c context.Context, req *itemspb.SellerViewRequest) (*itemspb.SellerViewRespsonse, error) {
	userID := req.UserID
	itemID := req.ItemID
	oid, err := getOID(itemID)
	if err != nil {
		return nil, errors.New("Invalid oid")
	}
	item, err := db.DetailView(*oid)
	if err != nil {
		return nil, err
	}
	if item.Seller != userID {
		return nil, errors.New("Item Not Found")
	}
	respone := &itemspb.SellerViewRespsonse{
		Seller:            item.Seller,
		Title:             item.Name,
		Description:       item.Description,
		Price:             item.Price,
		AvailableQuantity: item.AvailableQuantity,
		QuantitySold:      item.SoldQuantity,
		Status:            item.Status,
	}
	return respone, nil
}

//SearchItem : To steam the items available with given name
func (s *ItemService) SearchItem(req *itemspb.SearchItemRequest, stream itemspb.ItemService_SearchItemServer) error {
	name := req.Name
	result, err := db.SearchByName(name)
	if err != nil {
		return err
	}
	var items []domain.Item
	result.All(context.Background(), &items)
	for i := 0; i < len(items); i++ {
		current := items[i]
		item := &itemspb.ViewItem{
			ID:    current.ID.Hex(),
			Title: current.Name,
			Price: current.Price,
		}
		response := &itemspb.SearchItemResponse{
			Item: item,
		}
		stream.Send(response)

	}
	return nil
}
