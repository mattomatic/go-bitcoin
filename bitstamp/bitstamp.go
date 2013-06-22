package bitstamp

import (
	"encoding/json"
)

const (
	TickerUrl    = "https://bitstamp.net/api/ticker"
	OrderBookUrl = "https://www.bitstamp.net/api/order_book/"
)

func GetTicker() *Ticker {
	bytes := httpRequest(TickerUrl)
	ticker := &Ticker{}

	err := json.Unmarshal(bytes, ticker)

	if err != nil {
		panic(err.Error())
	}

	return ticker
}

func GetOrderBook() *OrderBook {
	bytes := httpRequest(OrderBookUrl)
	book := &OrderBook{}

	err := json.Unmarshal(bytes, book)

	if err != nil {
		panic(err.Error())
	}

	return book
}
