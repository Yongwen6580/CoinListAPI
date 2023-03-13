package method

type Coin struct {
	Id              string  `json:"id"`
	Name            string  `json:"name"`
	Symbol          string  `json:"symbol"`
	Usd             float64 `json:"usd"`
	Category        string  `json:"category"`
	MarketCap       float64 `json:"market_cap"`
	MarketCapChange float64 `json:"market_cap_change_24h"`
	Volume          float64 `json:"total_volume"`
	Market          int     `json:"market_data"`
}
