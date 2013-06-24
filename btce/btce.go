package btce

import (
	"github.com/mattomatic/go-bitcoin/algo"
	"github.com/mattomatic/go-bitcoin/common"
	"time"
)

const (
	ExchangeId   = "BTCE"
	OrderBookUrl = "https://btc-e.com/api/2/btc_usd/depth"
	PollInterval = time.Millisecond * 500
)

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
			for diff := range algo.GenerateDiffs(old, new) {
				ch <- diff
			}

			old = new
		}
	}()

	return ch
}
