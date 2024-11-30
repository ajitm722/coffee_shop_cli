package main

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "coffeecli",
		Short: "CLI for placing and viewing coffee orders",
	}

	// Subcommand: View all orders
	var viewCmd = &cobra.Command{
		Use:   "view",
		Short: "View all coffee orders",
		Run: func(cmd *cobra.Command, args []string) {
			resp, err := http.Get(viper.GetString("api.url") + "/orders")
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			fmt.Println(string(body))
		},
	}

	// Subcommand: Place a new order
	var placeCmd = &cobra.Command{
		Use:   "place [client] [coffee] [size] [quantity]",
		Short: "Place a new coffee order",
		Args:  cobra.ExactArgs(4),
		Run: func(cmd *cobra.Command, args []string) {
			client, coffee, size, quantity := args[0], args[1], args[2], args[3]
			orderJSON := fmt.Sprintf(`{"client": "%s", "coffee": "%s", "size": "%s", "quantity": %s}`,
				client, coffee, size, quantity)
			resp, err := http.Post(viper.GetString("api.url")+"/orders", "application/json", bytes.NewBuffer([]byte(orderJSON)))
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			fmt.Println(string(body))
		},
	}

	rootCmd.AddCommand(viewCmd, placeCmd)

	// Set up Viper for configuration
	viper.SetDefault("api.url", "http://localhost:8080")
	rootCmd.Execute()
}

