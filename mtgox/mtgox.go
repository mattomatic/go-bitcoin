package mtgox

const (
	ExchangeId = "MTGOX"
)

func GetDepthChannel() <-chan *Depth {
	connection := connect()
	connection.subscribe("depth")
	feeds := connection.poll()

	ch := make(chan *Depth)

	for {
		feed := <-feeds
		depth := &feed.message.(*DepthFeed).Depth
		ch <- depth
	}
}

func GetOrderBookChannel() <-chan *OrderBook {
	depths := GetDepthChannel()
	book := newOrderBook()
	ch := make(chan *OrderBook)

	for {
		depth := <-depths
		book.handleDepth(depth)
		ch <- book
	}
}

func GetTickerChannel() <-chan *Ticker {
	connection := connect()
	connection.subscribe("ticker")
	feeds := connection.poll()

	ch := make(chan *Ticker)

	for {
		feed := <-feeds
		ticker := &feed.message.(*TickerFeed).Ticker
		ch <- ticker
	}
}

func GetTradeChannel() <-chan *Trade {
	connection := connect()
	connection.subscribe("trade")
	feeds := connection.poll()

	ch := make(chan *Trade)

	for {
		feed := <-feeds
		trade := &feed.message.(*TradeFeed).Trade
		ch <- trade
	}
}
