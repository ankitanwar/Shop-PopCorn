package items

import itemspb "github.com/ankitanwar/e-Commerce/Products/proto"

//UserHistory : To keep track of all the items the user has ordered
type UserHistory struct {
	UserID int            `bson:"userID"`
	orders []itemspb.Item `bson:"orders"`
}
