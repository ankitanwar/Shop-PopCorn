package services

import (
	"context"
	"errors"
	"time"

	itemspb "github.com/ankitanwar/Shop-PopCorn/Products/proto"
	db "github.com/ankitanwar/Shop-PopCorn/Products/server/database"
	domain "github.com/ankitanwar/Shop-PopCorn/Products/server/domian/items"
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

func BuyItem(itemID string) (*domain.BuyItem, error) {
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
	deliveryTime := time.Now()
	deliveryTime.Format("01-02-2006")
	deliveryTime = deliveryTime.AddDate(0, 0, 10)
	result := &domain.BuyItem{
		DeliveryTime: deliveryTime.String(),
		Price:        item.Price,
		ItemName:     item.Name,
	}
	return result, nil

}

//Create : To Create the item
func (s *ItemService) Create(ctx context.Context, req *itemspb.CreateItemRequest) (*itemspb.CreateItemResposne, error) {
	if len(req.Name) == 0 {
		return nil, errors.New("Please Enter The Valid Name")
	}
	item := &domain.Item{
		Seller:            req.GetSeller(),
		Name:              req.GetName(),
		Description:       req.GetDescription(),
		Price:             req.GetPrice(),
		AvailableQuantity: req.GetAvailableQuantity(),
		SoldQuantity:      0,
		Status:            "Available",
	}
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
	itemID := req.GetID()
	oid, err := getOID(itemID)
	if err != nil {
		return nil, errors.New("Invalid Item ID")
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
	itemID := req.GetItemID()
	userID := req.GetUserID()
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
	if req.GetName() != "" {
		savedDetail.Name = req.GetName()
	}
	if req.Description != "" {
		savedDetail.Description = req.GetDescription()
	}
	if req.AvailableQuantity != 0 {
		savedDetail.AvailableQuantity += req.GetAvailableQuantity()
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
	itemID := req.GetItemID()
	userID := req.GetUserID()
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
	itemID := req.GetItemID()
	buyItem, err := BuyItem(itemID)
	if err != nil {
		return nil, errors.New("Unable To Buy The Product")
	}
	response := &itemspb.BuyItemResponse{
		ExceptedDateOfDilvery: buyItem.DeliveryTime,
		Title:                 buyItem.ItemName,
		Price:                 buyItem.Price,
	}
	return response, nil

}

//SellerView : To Give seller detail view about the item they are selling
func (s *ItemService) SellerView(c context.Context, req *itemspb.SellerViewRequest) (*itemspb.SellerViewRespsonse, error) {
	userID := req.GetUserID()
	itemID := req.GetItemID()
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
	name := req.GetName()
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

func (S *ItemService) CheckOut(ctx context.Context, req *itemspb.CheckoutRequest) (*itemspb.CheckOutResponse, error) {
	itemID := req.GetItemID()
	buyItem, err := BuyItem(itemID)
	if err != nil {
		return nil, errors.New("Unable To Buy The Product")
	}
	response := &itemspb.CheckOutResponse{
		ExceptedDateOfDilvery: buyItem.DeliveryTime,
		Title:                 buyItem.ItemName,
		Price:                 buyItem.Price,
	}
	return response, nil
}
