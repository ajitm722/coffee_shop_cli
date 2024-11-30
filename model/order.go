package model

// Order represents a coffee shop order.
type Order struct {
	ID      int    `json:"id"`
	Client  string `json:"client"`
	Drink   string `json:"drink"`
	Size    string `json:"size"`
	Comment string `json:"comment"`
}
