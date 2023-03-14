package main

import (
	"context"
	"fmt"
	"github.com/Yongwen6580/CoinListAPI/gRPC/coinService"
	pb "github.com/Yongwen6580/CoinListAPI/gRPC/coinService"
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

	//testing 1st method
	coins, err := client.List(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("could not list coins: %v", err)
	}
	for _, coin := range coins.Coins {
		log.Printf("ID: %s, Name: %s, Symbol: %s\n", coin.Id, coin.Name, coin.Symbol)
	}
	//testing 3rd method
	categories, err := client.GetCategories(context.Background(), &pb.GetCategoriesRequest{})
	if err != nil {
		log.Fatalf("could not get categories: %v", err)
	}
	for _, category := range categories.Categories {
		log.Printf("ID: %s, Market Cap: %f, Market Cap Change: %f\n", category.Id, category.MarketCap, category.MarketCapChange_24H)
	}

	//testing 2nd method
	var inputName string
	fmt.Println("Please enter the name of the coin for price:")
	fmt.Scanln(&inputName)
	price, err := client.GetTokenPrice(context.Background(), &coinService.GetTokenPriceRequest{Name: inputName})
	if err != nil {
		log.Fatalf("could not get token price: %v", err)
	}
	log.Printf("Price of %s: $%.2f\n", inputName, price.Usd)
}
