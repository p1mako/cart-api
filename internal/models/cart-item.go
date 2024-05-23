package models

type CartItem struct {
	Id       int
	CartId   int
	Product  string
	Quantity int
}
