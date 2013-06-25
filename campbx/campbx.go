package campbx

import (
	"encoding/json"
	"github.com/mattomatic/go-bitcoin/algo"
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
		old := newOrderBook()

		for new := range books {
			for diff := range algo.GenerateDiffs(old, new) {
				ch <- diff
			}

			old = new
		}
	}()

	return ch
}
