package main

import (
	"log"
)

func main() {
	//call the List() Method from CoinGecko Structs
	cg := &CoinGecko{}
	//coins := cg.List()
	//
	//for _, coin := range coins {
	//	log.Printf("ID: %s, Name: %s, Symbol: %s\n", coin.Id, coin.Name, coin.Symbol)
	//}
	//
	////call the GetTokenPrice() Method
	//price, err := cg.GetTokenPrice("bitcoin")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Printf("Price of bitcoin: $%.2f\n", price.Usd)

	//call the GetCategories() Method from categories Structs
	//categories, err := cg.GetCategories()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//for i, category := range categories {
	//	categories[i].Name = fmt.Sprintf("ID: %s", category.Id) // Store the value to the slice
	//	categories[i].Name = fmt.Sprintf("Coins: %d", category.CoinCount)
	//	categories[i].Name = fmt.Sprintf("Markets: %d\n", category.MarketCount)
	//	fmt.Printf("%s: %s\n", category.Id, category.Name)
	//	fmt.Printf("\tCoins: %d\n", category.CoinCount)
	//	fmt.Printf("\tMarkets: %d\n", category.MarketCount)
	//}

	categories, err := cg.GetCategories()
	if err != nil {
		log.Fatal(err)
	}

	for _, category := range categories {
		log.Printf("ID: %s, Markets: %f, MarketCapChange%f\n", category.Id, category.MarketCap, category.MarketCapChange)
	}
}
