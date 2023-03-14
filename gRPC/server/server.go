package main

import (
	"context"
	"fmt"
	pb "github.com/Yongwen6580/CoinListAPI/gRPC/coinService"
	main2 "github.com/Yongwen6580/CoinListAPI/method"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedCoinGeckoServer
}

func (s *server) mustEmbedUnimplementedCoinGeckoServer() {
	//TODO implement me
	panic("implement me")
}

func (s *server) List(ctx context.Context, req *pb.Empty) (*pb.ListResponse, error) {
	cg := &main2.CoinGecko{}
	coins := cg.List()
	var resp pb.ListResponse
	for _, coin := range coins {
		resp.Coins = append(resp.Coins, &pb.Coin{
			Id:     coin.Id,
			Name:   coin.Name,
			Symbol: coin.Symbol,
		})
	}
	return &resp, nil
}

func (s *server) GetTokenPrice(ctx context.Context, req *pb.GetTokenPriceRequest) (*pb.GetTokenPriceResponse, error) {
	cg := &main2.CoinGecko{}
	priceInput := req.Name
	price, err := cg.GetTokenPrice(priceInput)
	if err != nil {
		log.Fatal(err)
	}
	return &pb.GetTokenPriceResponse{Usd: price.Usd}, nil
}

func (s *server) GetCategories(ctx context.Context, req *pb.GetCategoriesRequest) (*pb.GetCategoriesResponse, error) {
	cg := &main2.CoinGecko{}
	categories, err := cg.GetCategories()
	if err != nil {
		return nil, err
	}
	var resp pb.GetCategoriesResponse
	for _, category := range categories {
		resp.Categories = append(resp.Categories, &pb.Coin{
			Id:                  category.Id,
			MarketCap:           category.MarketCap,
			MarketCapChange_24H: category.MarketCapChange,
		})
	}
	return &resp, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCoinGeckoServer(s, &server{})
	fmt.Println("Starting server on port :50053...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
