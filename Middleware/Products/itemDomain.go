package product

//Item : Item and its fields
type Item struct {
	Title             string `json:"Title"`
	Price             int    `json:"Price"`
	Status            string `json:"Status"`
	AvailableQuantity int    `json:"AvailableQuantity"`
}
