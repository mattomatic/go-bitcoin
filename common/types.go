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
	DepthFeed
	TradeFeed
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

type Feed struct {
	Type      FeedType
	Timestamp time.Time
	Message   interface{}
}

type Client interface {
	ToggleTickerFeeds()
	ToggleAsync()
	Channel() <-chan *Feed
}
