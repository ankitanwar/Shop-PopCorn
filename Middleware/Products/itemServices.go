package product

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ankitanwar/GoAPIUtils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
)

const (
	headerXAccessTokenID = "X-Token-ID"
	headerXCallerID      = "X-Caller-ID"
)

var (
	ItemSerivce itemServiceInterface = &itemServicesStruct{}
	headers                          = make(http.Header)
	restClient                       = rest.RequestBuilder{
		BaseURL: "http://localhost:8086",
		Timeout: 100 * time.Millisecond,
	}
)

type itemServicesStruct struct {
}

type itemServiceInterface interface {
	GetItemDetails(string) (*ItemValue, *errors.RestError)
	BuyItem(*http.Request, string) *errors.RestError
}

func checkItemID(itemID string) *errors.RestError {
	if len(itemID) < 0 {
		return errors.NewBadRequest("Invalid Item ID")
	}
	return nil
}

//GetCallerID : To get the caller id from the url
func GetCallerID(request *http.Request) string {
	if request == nil {
		return ""
	}
	callerID := request.Header.Get(headerXCallerID)
	return callerID
}

//GetTokenID : To get the access token ID of the given user
func GetTokenID(request *http.Request) string {
	if request == nil {
		return ""
	}
	callerID := request.Header.Get(headerXAccessTokenID)
	return callerID
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

func (item *itemServicesStruct) BuyItem(req *http.Request, itemID string) *errors.RestError {
	userID := GetCallerID(req)
	tokenID := GetTokenID(req)
	headers.Set(headerXCallerID, userID)
	headers.Set(headerXAccessTokenID, tokenID)
	err := checkItemID(itemID)
	if err != nil {
		return err
	}
	res := restClient.Post(fmt.Sprintf("/checkout/%s", itemID), nil)
	if res.StatusCode < 300 {
		return nil
	}
	return errors.NewBadRequest("Unable to purchase items")
}
