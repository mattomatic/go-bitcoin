package bitstamp

import (
	"encoding/json"
	"github.com/mattomatic/go-bitcoin/common"
	"time"
)

const (
	ExchangeId   = "BITSTAMP"
	TickerUrl    = "https://bitstamp.net/api/ticker"
	OrderBookUrl = "https://www.bitstamp.net/api/order_book/"
	PollInterval = time.Second
)

func GetTickerChannel() <-chan *Ticker {
	ch := make(chan *Ticker)

	go func() {
		for {
			time.Sleep(PollInterval)
			ch <- getTicker()
		}
	}()

	return ch
}

func GetOrderBookChannel() <-chan *OrderBook {
	ch := make(chan *OrderBook)

	go func() {
		for {
			time.Sleep(PollInterval)
			ch <- getOrderBook()
		}
	}()

	return ch
}

func GetDepthDiffChannel() <-chan common.DepthDiff {
	ch := make(chan common.DepthDiff)
	books := GetOrderBookChannel()

	go func() {
		defer close(ch)
		old := newOrderBook() // start with an empty book

		for new := range books {
			for diff := range common.GenerateDiffs(old, new) {
				ch <- diff
			}

			old = new
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
