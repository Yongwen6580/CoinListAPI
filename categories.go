package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type CoinGeckoCategories struct{}

func (cgc *CoinGeckoCategories) GetCategories() ([]Coin, error) {
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
		categories[i].Name = fmt.Sprintf("ID: %s", category.Id) // Store the value to the slice
		categories[i].Name = fmt.Sprintf("Coins: %d", category.CoinCount)
		categories[i].Name = fmt.Sprintf("Markets: %d\n", category.MarketCount)

		log.Printf("ID: %s, Coins: %d, Markets: %d\n", category.Id, category.CoinCount, category.MarketCount)
	}
	return categories, nil
}
