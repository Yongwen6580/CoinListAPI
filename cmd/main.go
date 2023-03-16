package main

import (
	"fmt"
	"github.com/Yongwen6580/CoinListAPI/pkg/api"
	"log"
)

func main() {
	//call the List() Method from CoinGecko Structs
	cg := &api.CoinGecko{}
	coins, err := cg.List()
	for _, coin := range coins {
		log.Printf("ID: %s, Name: %s, Symbol: %s\n", coin.Id, coin.Name, coin.Symbol)
	}

	//call the get GetTrending() method
	topCoin, err := cg.GetTrending()
	if err != nil {
		log.Fatal(err)
	}
	for _, topC := range topCoin {
		log.Printf(" Coin ID: %d, Name: %s, Symbol: %s, price_btc: %9f, Market Cap Rank: %d\n", topC.CoinID, topC.Name, topC.Symbol, topC.Price, topC.MarketCapRank)
	}
	//call the get GetTokenPrice() method
	var inputPrice string
	fmt.Println("Please enter the name of coin for price")
	fmt.Scanln(&inputPrice)
	price, err := cg.GetTokenPrice(inputPrice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Price of bitcoin: $%.2f\n", price.Usd)
}
