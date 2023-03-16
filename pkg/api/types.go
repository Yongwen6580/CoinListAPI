package api

type Coin struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type Price struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Symbol string  `json:"symbol"`
	Usd    float64 `json:"usd"`
}

type Top struct {
	Id            string  `json:"id"`
	CoinID        int64   `json:"coin_id"`
	Name          string  `json:"name"`
	Symbol        string  `json:"symbol"`
	Price         float64 `json:"price_btc"`
	MarketCapRank int     `json:"market_cap_rank"`
}
