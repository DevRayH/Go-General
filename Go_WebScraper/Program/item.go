package main

type Item struct {
	Name    string
	Price   float32
	Details string
	Rating  string
}

func newItem(name string, price float32, details string, rating string) *Item {
	var item Item = Item{name, price, details, rating}
	return &item
}