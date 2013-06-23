package bitstamp

import (
	"encoding/json"
	"time"
)

const (
	TickerUrl    = "https://bitstamp.net/api/ticker"
	OrderBookUrl = "https://www.bitstamp.net/api/order_book/"
)

const (
	SleepInterval = time.Second
)

func GetTickerChannel() <-chan *Ticker {
	ch := make(chan *Ticker)

	go func() {
		for {
			time.Sleep(SleepInterval)
			ch <- getTicker()
		}
	}()

	return ch
}

func GetOrderBookChannel() <-chan *OrderBook {
	ch := make(chan *OrderBook)

	go func() {
		for {
			time.Sleep(SleepInterval)
			ch <- getOrderBook()
		}
	}()

	return ch
}

func getTicker() *Ticker {
	bytes := httpRequest(TickerUrl)
	ticker := &Ticker{}

	err := json.Unmarshal(bytes, ticker)

	if err != nil {
		panic(err.Error())
	}

	return ticker
}

func getOrderBook() *OrderBook {
	bytes := httpRequest(OrderBookUrl)
	book := &OrderBook{}

	err := json.Unmarshal(bytes, book)

	if err != nil {
		panic(err.Error())
	}

	return book
}
