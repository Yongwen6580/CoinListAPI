package api

import (
	"testing"
)

func TestList(t *testing.T) {
	cg := CoinGecko{}
	coins, err := cg.List()
	if err != nil {
		t.Errorf("error calling List(): %v", err)
	}
	if len(coins) == 0 {
		t.Errorf("expected more than 0 coins, got %v", len(coins))
	}
}

func TestGetTrending(t *testing.T) {
	cg := CoinGecko{}
	topCoin, err := cg.GetTrending()
	if err != nil {
		t.Errorf("error calling GetTrending(): %v", err)
	}
	if len(topCoin) == 0 {
		t.Errorf("expected more than 0 top coins, got %v", len(topCoin))
	}
}

func TestGetTokenPrice(t *testing.T) {
	cg := CoinGecko{}
	price, err := cg.GetTokenPrice("bitcoin")
	if err != nil {
		t.Errorf("error calling GetTokenPrice(%v): %v", "bitcoin", err)
	}
	if price.Usd <= 0 {
		t.Errorf("expected price of %v to be greater than 0, got %.2v", "bitcoin", price.Usd)
		_, err := cg.List()
		if err != nil {
			t.Errorf("call to List failed!")
		}
	}
}
