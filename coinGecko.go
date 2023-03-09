package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type CoinGecko struct {
}

// List cg pionter, List() name of the method, type is []Coin
func (cg *CoinGecko) List() []Coin {
	resp, err := http.Get("https://api.coingecko.com/api/v3/coins/list")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	var coins []Coin //slice
	if err := json.NewDecoder(resp.Body).Decode(&coins); err != nil {
		panic(err)
	}
	/*
	   The response body of the request is then decoded into a slice of Coin structs (coins) using json.NewDecoder and json.Unmarshal.
	    If there's an error decoding the response, the program panics.
	*/
	for _, coin := range coins {
		log.Printf("ID: %s, Name: %s, Symbol: %s\n", coin.Id, coin.Name, coin.Symbol)
	}
	return coins
}
