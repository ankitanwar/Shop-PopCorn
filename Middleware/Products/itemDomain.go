package product

//ItemStruct : To get the value of item
type Item struct {
	Item ItemValue `json:"item"`
}

//ItemValue : Item and its fields
type ItemValue struct {
	Title             string `json:"Title"`
	Price             int    `json:"Price"`
	Status            string `json:"Status"`
	AvailableQuantity int    `json:"AvailableQuantity"`
}

type BuyResponse struct {
	Price       int64  `json:"price"`
	DeliverDate string `json:"ExceptedDateOfDilvery"`
	Title       string `json:"Title"`
}
