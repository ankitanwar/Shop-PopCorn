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
}

var (
	restClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8086",
		Timeout: 100 * time.Millisecond,
	}
)

func (item *itemServicesStruct) GetItemDetails(itemID string) (*ItemValue, *errors.RestError) {
	if len(itemID) < 0 {
		return nil, errors.NewBadRequest("please Enter the valid Item ID")
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
