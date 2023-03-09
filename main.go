package main

import (
	"log"
)

func main() {
	//call the List() Method from CoinGecko Structs
	cg := &CoinGecko{}
	coins := cg.List()
	for _, coin := range coins {
		log.Printf("ID: %s, Name: %s, Symbol: %s\n", coin.Id, coin.Name, coin.Symbol)
	}

	//call the GetTokenPrice() Method
	price, err := cg.GetTokenPrice("bitcoin")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Price of bitcoin: $%.2f\n", price.Usd)

	//call the get GetCategories() method
	categories, err := cg.GetCategories()
	if err != nil {
		log.Fatal(err)
	}
	for _, category := range categories {
		log.Printf("ID: %s, Markets: %f, MarketCapChange%f\n", category.Id, category.MarketCap, category.MarketCapChange)
	}
}
