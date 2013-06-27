package btce

import (
	"github.com/mattomatic/go-bitcoin/algo"
	"github.com/mattomatic/go-bitcoin/common"
	"time"
)

const (
	ExchangeId   = "BTCE"
	ExchangeFee  = .0020
	OrderBookUrl = "https://btc-e.com/api/2/btc_usd/depth"
	PollInterval = time.Millisecond * 500
)

func GetOrderBookChannel() <-chan *OrderBook {
	ch := make(chan *OrderBook)

	go func() {
		defer close(ch)
		failed := false

		for {
			time.Sleep(PollInterval)
			book, err := getOrderBook()

			if failed && err != nil {
				panic("consecutive errors!")
			}

			failed = err != nil

			if failed {
				continue
			}

			ch <- book
		}
	}()

	return ch
}

func GetDiffChannel() <-chan common.Diff {
	ch := make(chan common.Diff)
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
