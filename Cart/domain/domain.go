package domin

//User : Structure of userCart
type User struct {
	UserID string `bson:"_id"`
	Items  []Item `bson:"items"`
}

//Item : Item struct
type Item struct {
	ItemID   string `bson:"itemID"`
	Title    string `bson:"name"`
	Price    int    `bson:"price"`
	Quantity int    `bson:"quantity"`
}
