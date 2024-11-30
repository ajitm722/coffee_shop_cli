package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	// Root command
	var rootCmd = &cobra.Command{
		Use:   "coffeecli",
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
		Use:   "place [client] [coffee] [size] [quantity]",
		Short: "Place a new coffee order",
		Args:  cobra.ExactArgs(4),
		Run: func(cmd *cobra.Command, args []string) {
			client, drink, size, quantity := args[0], args[1], args[2], args[3]
			orderJSON := fmt.Sprintf(`{"client": "%s", "drink": "%s", "size": "%s", "quantity": %s}`,
				client, drink, size, quantity)

			resp, err := http.Post(viper.GetString("api.url")+"/orders", "application/json", bytes.NewBuffer([]byte(orderJSON)))
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
