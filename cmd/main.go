package main

import (
	"log"

	"github.com/Yongwen6580/CoinListAPI/pkg/api"
)

func main() {
	//call the List() Method from CoinGecko Structs
	cg := &api.CoinGecko{}
	coins, err := cg.List()
	for _, coin := range coins {
		log.Printf("ID: %v, Name: %v, Symbol: %v\n", coin.Id, coin.Name, coin.Symbol)
	}
	//call the get GetTrending() method
	topCoin, err := cg.GetTrending()
	if err != nil {
		log.Fatal(err)
	}
	for _, topC := range topCoin {
		log.Printf(" Coin ID: %v, Name: %v, Symbol: %v, price_btc: %v, Market Cap Rank: %v\n", topC.CoinID, topC.Name, topC.Symbol, topC.Price, topC.MarketCapRank)
	}
	//call the get GetTokenPrice() method
	price, err := cg.GetTokenPrice("bitcoin")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Price of bitcoin: $%v\n", price.Usd)
}
