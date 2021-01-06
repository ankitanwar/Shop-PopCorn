package items

import itemspb "github.com/ankitanwar/e-Commerce/Items-A.P.I/proto"

//UserHistory : To keep track of all the items the user has ordered
type UserHistory struct {
	UserID int            `bson:"userID"`
	orders []itemspb.Item `bson:"orders"`
}
