package main

import (
	"fmt"
	"github.com/Yongwen6580/CoinListAPI/method"
	"log"
)

func main() {
	//call the List() Method from CoinGecko Structs
	cg := &method.CoinGecko{}
	coins := cg.List()
	for _, coin := range coins {
		log.Printf("ID: %s, Name: %s, Symbol: %s\n", coin.Id, coin.Name, coin.Symbol)
	}

	//call the GetTokenPrice() Method
	var inputPrice string
	fmt.Println("Please enter the name of coin for price")
	fmt.Scanln(&inputPrice)
	price, err := cg.GetTokenPrice(inputPrice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Price of bitcoin: $%.2f\n", price.Usd)

	//call the get GetCategories() method
	//categories, err := cg.GetCategories()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//for _, category := range categories {
	//	log.Printf("ID: %s, Markets: %f, MarketCapChange%f\n", category.Id, category.MarketCap, category.MarketCapChange)
	//}
}
