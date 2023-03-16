package main

import (
	"context"
	"fmt"
	pb "github.com/Yongwen6580/CoinListAPI/pkg/gRPC/coinService"
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
	client := pb.NewCoinGeckoClient(conn)

	//testing list method
	coins, err := client.List(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("could not list coins: %v", err)
	}
	for _, coin := range coins.Coins {
		log.Printf("ID: %s, Name: %s, Symbol: %s\n", coin.Id, coin.Name, coin.Symbol)
	}

	//testing trending method
	trendingCoin, err := client.GetTrendingCoins(context.Background(), &pb.GetTrendingCoinsRequest{})
	if err != nil {
		log.Fatalf("could not get Treading: %v", err)
	}
	for _, trending := range trendingCoin.TopCoin {
		log.Printf("ID: %s, Coin ID: %d, Name: %s, Symbol: %s, Price btc: %.10f, Market Cap Change: %d\n", trending.Id, trending.CoinID, trending.Name, trending.Symbol, trending.Price, trending.MarketCap)
	}

	//testing price method
	var inputName string
	fmt.Println("Please enter the name of the coin for price:")
	fmt.Scanln(&inputName)
	price, err := client.GetTokenPrice(context.Background(), &pb.GetTokenPriceRequest{Name: inputName})
	if err != nil {
		log.Fatalf("could not get token price: %v", err)
	}
	log.Printf("Price of %s: $%.2f\n", inputName, price.Usd)

	//testing the third method

}
