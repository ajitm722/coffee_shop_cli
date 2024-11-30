package main

import (
	"coffee_shop/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Routes
	r.GET("/orders", handler.GetOrders)
	r.GET("/orders/:id", handler.GetOrderById)
	r.POST("/orders", handler.PlaceOrder)

	// Start server
	r.Run(":8080") // Start on localhost:8080
}
