package main

import (
	pb "github.com/Yongwen6580/CoinListAPI/pkg/gRPC/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"testing"
)

const (
	address = "localhost:50053"
)

func TestList(t *testing.T) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewCoinGeckoClient(conn)
	// Call the List() method
	coins, err := client.List(context.Background(), &pb.ListRequest{})
	if err != nil {
		t.Fatalf("could not list coins: %v", err)
	}
	if len(coins.Coins) == 0 {
		t.Errorf("expected more than 0 coins, got %v", len(coins.Coins))
	}
}

func TestGetTrending(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewCoinGeckoClient(conn)
	trendingCoin, err := client.GetTrendingCoins(context.Background(), &pb.GetTrendingCoinsRequest{})
	if err != nil {
		t.Fatalf("could not get Treading: %v", err)
	}
	if len(trendingCoin.TopCoin) == 0 {
		t.Errorf("expected more than 0 top coins, got %v", len(trendingCoin.TopCoin))
	}
}

func TestGetTokenPrice(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewCoinGeckoClient(conn)
	price, err := client.GetTokenPrice(context.Background(), &pb.GetTokenPriceRequest{Name: "bitcoin"})
	if err != nil {
		t.Fatalf("could not get token price: %v", err)
	}
	if price.Usd <= 0 {
		t.Errorf("expected price of %v to be greater than 0, got %.2v", "bitcoin", price.Usd)
	}
}
