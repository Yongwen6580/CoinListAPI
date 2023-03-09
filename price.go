package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Price struct {
}

// GetTokenPrice
/*
takes a tokenId string as a parameter and returns a pointer to a TokenPrice struct and an error (if any).
*/
func (cgp *Price) GetTokenPrice(tokenId string) (*Coin, error) {

	//url := fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=usd", tokenId)
	//resp, err := http.Get(url)
	//if err != nil {
	//	return nil, err
	//}

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
