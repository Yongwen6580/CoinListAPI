package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type CoinGecko struct {
}

// List cg receiver, List() name of the method, type is []Coin
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

func (cg *CoinGecko) GetTokenPrice(tokenId string) (*Coin, error) {

	url := "https://api.coingecko.com/api/v3/simple/price?ids=" + tokenId + "&vs_currencies=usd"
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching price for token %s: %s", tokenId, err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status %d", resp.StatusCode)
	}
	/*
		Store the prices of tokens returned by the Coingecko API.
		The outer map has a key for each token ID, and the corresponding value is an inner map that
		contains the prices of the token in USD
	*/
	var price map[string]map[string]float64
	if err := json.NewDecoder(resp.Body).Decode(&price); err != nil {
		return nil, err
	}
	//assigns the USD price to the usdPrice variable and assigns a boolean value to the ok variable that
	//indicates whether the USD price was found in the price map for the given token ID.
	usdPrice, ok := price[tokenId]["usd"]
	if !ok {
		return nil, fmt.Errorf("USD price not found for token %s", tokenId)
	}
	return &Coin{Usd: usdPrice}, nil
}

func (cg *CoinGecko) GetCategories() ([]Coin, error) {
	resp, err := http.Get("https://api.coingecko.com/api/v3/coins/categories")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get categories, status code: %d", resp.StatusCode)
	}

	var categories []Coin
	if err := json.NewDecoder(resp.Body).Decode(&categories); err != nil {
		return nil, err
	}

	for i, category := range categories {
		categories[i].Name = fmt.Sprintf("ID: %s, Markets: %f, MarketCapChange%f\n", category.Id, category.MarketCap, category.MarketCapChange)
		log.Printf("ID: %s, Markets: %f, MarketCapChange%f\n", category.Id, category.MarketCap, category.MarketCapChange)
	}

	return categories, nil
}
