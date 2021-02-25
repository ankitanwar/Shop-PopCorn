package product

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ankitanwar/GoAPIUtils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	ItemSerivce itemServiceInterface = &itemServicesStruct{}
)

type itemServicesStruct struct {
}

type itemServiceInterface interface {
	GetItemDetails(string) (*ItemValue, *errors.RestError)
	BuyItem(string) *errors.RestError
}

var (
	restClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8086",
		Timeout: 100 * time.Millisecond,
	}
)

func checkItemID(itemID string) *errors.RestError {
	if len(itemID) < 0 {
		return errors.NewBadRequest("Invalid Item ID")
	}
	return nil
}

func (item *itemServicesStruct) GetItemDetails(itemID string) (*ItemValue, *errors.RestError) {
	err := checkItemID(itemID)
	if err != nil {
		return nil, err
	}
	res := restClient.Get(fmt.Sprintf("/items/%s", itemID))
	if res.Response == nil || res == nil {
		return nil, errors.NewInternalServerError("Error while fetching the item details")
	}
	product := &Item{}
	if res.StatusCode < 299 {
		err := json.Unmarshal(res.Bytes(), &product)
		if err != nil {
			return nil, errors.NewInternalServerError("Error while unmarshalling the data")
		}
		r := &ItemValue{}
		r.AvailableQuantity = product.Item.AvailableQuantity
		r.Price = product.Item.Price
		r.Status = product.Item.Status
		r.Title = product.Item.Title
		return r, nil
	}
	return nil, errors.NewInternalServerError("Error while getting the items details")
}

func (item *itemServicesStruct) BuyItem(itemID string) *errors.RestError {
	err := checkItemID(itemID)
	if err != nil {
		return err
	}
	res := restClient.Post(fmt.Sprintf("/items/buy/%s", itemID), nil)
	if res.StatusCode == 200 {
		return nil
	}
	return errors.NewBadRequest("Unable to purchase items")
}
