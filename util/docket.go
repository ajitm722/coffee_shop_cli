package util

import (
	"coffee_shop/model"
	"fmt"
)

// PrintDocket simulates the printing of a coffee shop order docket.
func PrintDocket(order model.Order) {
	fmt.Println("=================================")
	fmt.Printf("Order ID: %d\n", order.ID)
	fmt.Printf("Client: %s\n", order.Client)
	fmt.Printf("Drink: %s\n", order.Drink)
	fmt.Printf("Size: %s\n", order.Size)
	if order.Comment != "" {
		fmt.Printf("Comment: %s\n", order.Comment)
	}
	fmt.Println("=================================")
}
