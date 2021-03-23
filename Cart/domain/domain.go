package domin

//User : Structure of userCart
type User struct {
	UserID string `bson:"_id"`
	Items  []Item `bson:"items" json:"items"`
}

//Item : Item struct
type Item struct {
	ItemID string `bson:"itemID" json:"itemID"`
	Title  string `bson:"name" json:"name"`
	Price  int    `bson:"price" json:"price"`
}

type ByResponse struct {
	Price       int64  `json:"price"`
	DeliverDate string `json:"ExceptedDateOfDilvery"`
	Title       string `json:"Title"`
}

type CheckoutResponse struct {
	HouseNumber string       `json:"houseNo"`
	Street      string       `json:"street"`
	State       string       `json:"state"`
	Country     string       `json:"country"`
	Phone       string       `json:"phone"`
	Products    []ByResponse `json:"products"`
	TotalCost   int64        `json:"total_cost"`
}
