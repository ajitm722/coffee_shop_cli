package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	// Root command
	var rootCmd = &cobra.Command{
		Use:   "./coffee-cli",
		Short: "CLI for managing coffee orders",
	}

	// Subcommand: View all orders
	var viewCmd = &cobra.Command{
		Use:   "view",
		Short: "View all coffee orders",
		Run: func(cmd *cobra.Command, args []string) {
			resp, err := http.Get(viper.GetString("api.url") + "/orders")
			if err != nil {
				fmt.Println("Error: Unable to fetch orders,", err)
				return
			}
			defer resp.Body.Close()

			body, _ := io.ReadAll(resp.Body)
			if resp.StatusCode == http.StatusOK {
				fmt.Println("All Orders:")
				fmt.Println(string(body))
			} else {
				fmt.Printf("Failed to fetch orders: %s\n", resp.Status)
			}
		},
	}

	// Subcommand: Place a new order
	var placeCmd = &cobra.Command{
		Use:   "place [client] [drink] [size] [quantity] [comment]",
		Short: "Place a new coffee order",
		Args:  cobra.MinimumNArgs(4), // At least 4 arguments (comment is optional)
		Run: func(cmd *cobra.Command, args []string) {
			client, drink, size, quantity := args[0], args[1], args[2], args[3]
			comment := ""
			if len(args) > 4 {
				comment = args[4] // Optional comment
			}

			order := map[string]interface{}{
				"client":   client,
				"drink":    drink,
				"size":     size,
				"quantity": quantity,
				"comment":  comment,
			}

			orderJSON, err := json.Marshal(order)
			if err != nil {
				fmt.Println("Error: Unable to create order payload,", err)
				return
			}

			resp, err := http.Post(viper.GetString("api.url")+"/orders", "application/json", bytes.NewBuffer(orderJSON))
			if err != nil {
				fmt.Println("Error: Unable to place order,", err)
				return
			}
			defer resp.Body.Close()

			body, _ := io.ReadAll(resp.Body)
			if resp.StatusCode == http.StatusCreated {
				fmt.Println("Order placed successfully:")
				fmt.Println(string(body))
			} else {
				fmt.Printf("Failed to place order: %s\n", resp.Status)
			}
		},
	}

	// Subcommand: Get order by ID
	var getCmd = &cobra.Command{
		Use:   "get [order_id]",
		Short: "Get details of a specific order by its ID",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			orderID := args[0]
			resp, err := http.Get(viper.GetString("api.url") + "/orders/" + orderID)
			if err != nil {
				fmt.Println("Error: Unable to fetch order,", err)
				return
			}
			defer resp.Body.Close()

			body, _ := io.ReadAll(resp.Body)
			if resp.StatusCode == http.StatusOK {
				fmt.Printf("Details of Order ID %s:\n", orderID)
				fmt.Println(string(body))
			} else {
				fmt.Printf("Failed to fetch order: %s\n", resp.Status)
			}
		},
	}

	// Add subcommands to root command
	rootCmd.AddCommand(viewCmd, placeCmd, getCmd)

	// Set up Viper for configuration
	viper.SetDefault("api.url", "http://localhost:8080")

	// Execute the root command
	rootCmd.Execute()
}
