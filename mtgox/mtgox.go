package mtgox

const (
	ExchangeId = "MTGOX"
)

func GetDepthDiffChannel() <-chan *Depth {
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

func GetOrderBookChannel() <-chan *OrderBook {
	depths := GetDepthDiffChannel()
	book := newOrderBook()
	ch := make(chan *OrderBook)

	go func() {
		defer close(ch)
		for {
			depth := <-depths
			book.handleDepth(depth)
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
