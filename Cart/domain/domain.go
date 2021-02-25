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
