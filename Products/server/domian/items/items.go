package items

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Item : Struct it contains all the value item has
type Item struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	Seller            string             `bson:"seller"`
	Name              string             `bson:"name"`
	Description       string             `bson:"description"`
	Price             int64              `bson:"price"`
	AvailableQuantity int64              `bson:"availablequantity"`
	SoldQuantity      int64              `bson:"quantitysold"`
	Status            string             `bson:"status"`
}

type BuyItem struct {
	ItemName     string
	DeliveryTime string
	Price        int64
}
