package mtgox

import (
	"github.com/mattomatic/go-bitcoin/common"
)

const (
	ExchangeId  = "MTGOX"
	ExchangeFee = 0.0060
)

func GetDiffChannel() <-chan common.Diff {
	// channels are invariant so we have to wrap this
	depthdiffs := make(chan common.Diff)
	diffs := getDepthDiffChannel()

	go func() {
		defer close(depthdiffs)
		for diff := range diffs {
			depthdiffs <- diff
		}
	}()

	return depthdiffs
}

func getDepthDiffChannel() <-chan *Depth {
	connection := connect()
	connection.subscribe("depth")
	feeds := connection.poll()
	ch := make(chan *Depth)

	go func() {
		defer close(ch)
		for {
			feed := <-feeds
			depth := &feed.message.(*DepthFeed).Depth
			ch <- depth
		}
	}()

	return ch
}

func GetOrderBookChannel() <-chan common.OrderBook {
	depths := getDepthDiffChannel()
	book := common.NewBook()
	ch := make(chan common.OrderBook)

	go func() {
		defer close(ch)
		for {
			depth := <-depths
			book.ApplyDiff(depth)
			ch <- book
		}
	}()

	return ch
}

func GetTickerChannel() <-chan *Ticker {
	connection := connect()
	connection.subscribe("ticker")
	feeds := connection.poll()
	ch := make(chan *Ticker)

	go func() {
		defer close(ch)
		for {
			feed := <-feeds
			ticker := &feed.message.(*TickerFeed).Ticker
			ch <- ticker
		}
	}()

	return ch
}

func GetTradeChannel() <-chan *Trade {
	connection := connect()
	connection.subscribe("trade")
	feeds := connection.poll()
	ch := make(chan *Trade)

	go func() {
		defer close(ch)
		for {
			feed := <-feeds
			trade := &feed.message.(*TradeFeed).Trade
			ch <- trade
		}
	}()

	return ch
}
