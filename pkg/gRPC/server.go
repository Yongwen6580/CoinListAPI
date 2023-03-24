package main

import (
	"context"
	cg "github.com/Yongwen6580/CoinListAPI/pkg/api"
	pb "github.com/Yongwen6580/CoinListAPI/pkg/gRPC/proto"
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

func (s *server) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	coinGecko := &cg.CoinGecko{}
	coins, err := coinGecko.List()
	if err != nil {
		return nil, err
	}
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
	coinGecko := &cg.CoinGecko{}
	priceInput := req.Name
	price, err := coinGecko.GetTokenPrice(priceInput)
	if err != nil {
		return nil, err
	}
	return &pb.GetTokenPriceResponse{Usd: price.Usd}, nil
}

func (s *server) GetTrendingCoins(ctx context.Context, req *pb.GetTrendingCoinsRequest) (*pb.GetTrendingCoinsResponse, error) {
	coinGecko := &cg.CoinGecko{}
	topCoin, err := coinGecko.GetTrending()
	if err != nil {
		return nil, err
	}
	resp := &pb.GetTrendingCoinsResponse{}
	for _, trending := range topCoin[:7] {
		resp.TopCoin = append(resp.TopCoin, &pb.Coin{
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
	pb.RegisterCoinGeckoServer(s, &server{})
	log.Println("Starting server on port :50053...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
