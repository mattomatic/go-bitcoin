package common

import (
	"time"
)

type Volume float64
type Price float64
type Exchange string
type Currency string
type Symbol string
type TradeId string
type FeedType int

const (
	TickerFeed FeedType = iota
	OrderBookFeed
	TradeFeed
	DepthFeed
)

type Trade interface {
	GetExchange() Exchange
	GetSymbol() Symbol
	GetTimestamp() time.Time
	GetVolume() Volume
	GetPrice() Price
	GetCurrency() Currency
	GetTradeId() TradeId
}

type Ticker interface {
	GetExchange() Exchange
	GetSymbol() Symbol
	GetTimestamp() time.Time
	GetBid() Price
	GetAsk() Price
	GetHigh() Price
	GetLow() Price
	GetVolume() Volume
}

type Feed interface {
	GetType() FeedType
	GetTimestamp() time.Time
	GetMessage() interface{}
}

type Client interface {
	ToggleTickerFeeds()
	ToggleOrderBookFeeds()
	ToggleAsync()
	Channel() <-chan Feed
}

type Order interface {
	GetExchange() Exchange
	GetSymbol() Symbol
	GetVolume() Volume
	GetPrice() Price
}

type OrderBook interface {
	GetBids() <-chan Order
	GetAsks() <-chan Order
}
