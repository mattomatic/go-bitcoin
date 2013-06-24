package campbx

import (
	"encoding/json"
	"fmt"
	"github.com/mattomatic/go-bitcoin/common"
	"time"
)

const (
	ExchangeId   = "CAMPBX"
	TickerUrl    = "http://campbx.com/api/xticker.php"
	OrderBookUrl = "http://CampBX.com/api/xdepth.php"
	PollInterval = time.Second
)

func GetTickerChannel() <-chan *Ticker {
	ch := make(chan *Ticker)
	ticker := &Ticker{}

	go func() {
		defer close(ch)
		for {
			time.Sleep(PollInterval)

			bytes := httpRequest(TickerUrl)
			err := json.Unmarshal(bytes, ticker)

			if err != nil {
				panic(err.Error())
			}

			ch <- ticker
		}
	}()

	return ch
}

func GetOrderBookChannel() <-chan *OrderBook {
	ch := make(chan *OrderBook)
	book := &OrderBook{}

	go func() {
		defer close(ch)

		for {
			fmt.Println("polling")
			time.Sleep(PollInterval)

			bytes := httpRequest(OrderBookUrl)
			err := json.Unmarshal(bytes, book)

			if err != nil {
				panic(err.Error())
			}

			ch <- book
		}
	}()

	return ch
}

func GetDepthDiffChannel() <-chan common.DepthDiff {
	ch := make(chan common.DepthDiff)
	books := GetOrderBookChannel()

	go func() {
		defer close(ch)
		old := newOrderBook()

		for new := range books {
			for diff := range common.GenerateDiffs(old, new) {
				ch <- diff
			}

			old = new
		}
	}()

	return ch
}
