package handler

import (
	"coffee_shop/model"
	"coffee_shop/util"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	orders []model.Order // In-memory order storage
	mu     sync.Mutex    // Mutex for thread-safe operations
)

// PlaceOrder handles placing a new order.
func PlaceOrder(c *gin.Context) {
	var order model.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	// Assign unique ID and add to order list
	order.ID = len(orders) + 1
	orders = append(orders, order)

	// Simulate docket printing
	fmt.Println("Printing docket:")
	util.PrintDocket(order)

	c.JSON(http.StatusCreated, gin.H{"message": "Order placed successfully", "order": order})
}

// GetOrders returns all orders.
func GetOrders(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	if len(orders) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No orders found"})
		return
	}
	c.JSON(http.StatusOK, orders)
}

// GetOrderById returns an order by its ID.
func GetOrderById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for _, order := range orders {
		if order.ID == id {
			c.JSON(http.StatusOK, order)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
}
