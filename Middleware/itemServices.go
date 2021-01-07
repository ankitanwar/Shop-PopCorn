package services

import (
	"encoding/json"
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
	GetItemDetails(string) (*Item, *errors.RestError)
}

var (
	restClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8086",
		Timeout: 100 * time.Millisecond,
	}
)

func (item *itemServicesStruct) GetItemDetails(itemID string) (*Item, *errors.RestError) {
	if len(itemID) < 0 {
		return nil, errors.NewBadRequest("please Enter the valid Item ID")
	}
	res := restClient.Get("/itemID")
	if res.Response == nil || res == nil {
		return nil, errors.NewInternalServerError("Error while fetching the item details")
	}
	product := &Item{}
	if res.StatusCode < 299 {
		err := json.Unmarshal(res.Bytes(), &product)
		if err != nil {
			return nil, errors.NewInternalServerError("Error while unmarshalling the data")
		}
		return product, nil
	}
	return nil, errors.NewInternalServerError("Error while getting the items details")
}
