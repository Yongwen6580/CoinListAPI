package main

import (
	"context"
	"log"

	pb "github.com/Yongwen6580/CoinListAPI/pkg/gRPC/proto"
	"google.golang.org/grpc"
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
	coins, err := client.List(context.Background(), &pb.ListRequest{})
	if err != nil {
		log.Fatalf("could not list coins: %v", err)
	}
	for _, coin := range coins.Coins {
		log.Printf("ID: %v, Name: %v, Symbol: %v\n", coin.Id, coin.Name, coin.Symbol)
	}

	//testing trending method
	trendingCoin, err := client.GetTrendingCoins(context.Background(), &pb.GetTrendingCoinsRequest{})
	if err != nil {
		log.Fatalf("could not get Treading: %v", err)
	}
	for _, trending := range trendingCoin.TopCoin {
		log.Printf("ID: %v, Coin ID: %v, Name: %v, Symbol: %v, Price btc: %v, Market Cap Change: %v\n", trending.Id, trending.CoinID, trending.Name, trending.Symbol, trending.Price, trending.MarketCap)
	}

	//testing price method
	name := "bitcoin"
	price, err := client.GetTokenPrice(context.Background(), &pb.GetTokenPriceRequest{Name: name})
	if err != nil {
		log.Fatalf("could not get token price: %v", err)
	}
	log.Printf("Price of %v: $%v\n", name, price.Usd)
}
