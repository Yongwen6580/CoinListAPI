package main

import (
	"context"
	pb "github.com/Yongwen6580/CoinListAPI/pkg/api"
	"github.com/Yongwen6580/CoinListAPI/pkg/gRPC/coinService"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	coinService.UnimplementedCoinGeckoServer
}

func (s *server) mustEmbedUnimplementedCoinGeckoServer() {
	//TODO implement me
	panic("implement me")
}

func (s *server) List(ctx context.Context, req *coinService.Empty) (*coinService.ListResponse, error) {
	cg := &pb.CoinGecko{}
	coins, err := cg.List()
	if err != nil {
		return nil, err
	}
	var resp coinService.ListResponse
	for _, coin := range coins {
		resp.Coins = append(resp.Coins, &coinService.Coin{
			Id:     coin.Id,
			Name:   coin.Name,
			Symbol: coin.Symbol,
		})
	}
	return &resp, nil
}

func (s *server) GetTokenPrice(ctx context.Context, req *coinService.GetTokenPriceRequest) (*coinService.GetTokenPriceResponse, error) {
	cg := &pb.CoinGecko{}
	priceInput := req.Name
	price, err := cg.GetTokenPrice(priceInput)
	if err != nil {
		return nil, err
	}
	return &coinService.GetTokenPriceResponse{Usd: price.Usd}, nil
}

func (s *server) GetTrendingCoins(ctx context.Context, req *coinService.GetTrendingCoinsRequest) (*coinService.GetTrendingCoinsResponse, error) {
	cg := &pb.CoinGecko{}
	topCoin, err := cg.GetTrending()
	if err != nil {
		return nil, err
	}
	resp := &coinService.GetTrendingCoinsResponse{}
	for _, trending := range topCoin[:7] {
		resp.TopCoin = append(resp.TopCoin, &coinService.Coin{
			Id:        trending.Id,
			CoinID:    trending.CoinID,
			Name:      trending.Name,
			Symbol:    trending.Symbol,
			Price:     float32(trending.Price),
			MarketCap: int64(trending.MarketCapRank),
		})
	}
	return resp, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	coinService.RegisterCoinGeckoServer(s, &server{})
	log.Println("Starting server on port :50053...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
