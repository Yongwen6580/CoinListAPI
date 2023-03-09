package main

type Coin struct {
	Id              string  `json:"id"`
	Name            string  `json:"name"`
	Symbol          string  `json:"symbol"`
	Usd             float64 `json:"usd"`
	CoinCount       int     `json:"coin_count"`
	MarketCount     int     `json:"market_count"`
	Category        string  `json:"category"`
	MarketCap       float64 `json:"market_cap"`
	MarketCapChange float64 `json:"market_cap_change_24h"`
	Volume          float64 `json:"total_volume"`
	Dominance       float64 `json:"market_cap_percentage"`
	Market          int     `json:"market_data"`
}
