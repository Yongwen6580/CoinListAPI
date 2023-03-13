package main

import (
	"CoinList/gRPC/coinService"
	pb "CoinList/gRPC/coinService"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	client := coinService.NewCoinGeckoClient(conn)

	// Prompt user for input.
	var input string
	fmt.Println("Please select an option，enter the corresponding number:")
	fmt.Println("1. List Coins")
	fmt.Println("2. Get Token Price")
	fmt.Println("3. Get Categories")
	fmt.Scanln(&input)

	// Call the appropriate gRPC method based on user input.
	switch input {
	case "1":
		coins, err := client.List(context.Background(), &pb.Empty{})
		if err != nil {
			log.Fatalf("could not list coins: %v", err)
		}
		for _, coin := range coins.Coins {
			log.Printf("ID: %s, Name: %s, Symbol: %s\n", coin.Id, coin.Name, coin.Symbol)
		}
	case "2":
		var inputName string
		fmt.Println("Please enter the name of the coin for price:")
		fmt.Scanln(&inputName)
		price, err := client.GetTokenPrice(context.Background(), &coinService.GetTokenPriceRequest{Name: inputName})
		if err != nil {
			log.Fatalf("could not get token price: %v", err)
		}
		log.Printf("Price of %s: $%.2f\n", inputName, price.Usd)
	case "3":
		categories, err := client.GetCategories(context.Background(), &pb.GetCategoriesRequest{})
		if err != nil {
			log.Fatalf("could not get categories: %v", err)
		}
		for _, category := range categories.Categories {
			log.Printf("ID: %s, Market Cap: %f, Market Cap Change: %f\n", category.Id, category.MarketCap, category.MarketCapChange_24H)
		}
	default:
		log.Fatalf("invalid input")
	}
}
