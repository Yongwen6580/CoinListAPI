package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type CoinGecko struct {
}

func (cg *CoinGecko) List() ([]Coin, error) {
	resp, err := http.Get("https://api.coingecko.com/api/v3/coins/list")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var coins []Coin //slice
	if err := json.NewDecoder(resp.Body).Decode(&coins); err != nil {
		return nil, err
	}
	return coins, nil
}

func (cg *CoinGecko) GetTokenPrice(tokenId string) (*Price, error) {
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
	return &Price{Usd: usdPrice}, nil
}

func (cg *CoinGecko) GetTrending() ([]Top, error) {
	resp, err := http.Get("https://api.coingecko.com/api/v3/search/trending")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get trending coins, status code: %d", resp.StatusCode)
	}
	var trending struct {
		Coins []struct {
			Item Top `json:"item"`
		} `json:"coins"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&trending); err != nil {
		return nil, err
	}
	var topCoins []Top
	for _, c := range trending.Coins {
		topCoins = append(topCoins, c.Item)
	}
	return topCoins[:7], nil
}
