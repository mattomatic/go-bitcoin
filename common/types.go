package common

import (
	"time"
)

type (
	Volume   float64
	Price    float64
	Exchange string
	Currency string
	Symbol   string
	TradeId  string
	Side     int
)

const (
	Bid Side = iota
	Ask
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
	GetBid() Price
	GetAsk() Price
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

type DepthDiff interface {
	GetExchange() Exchange
	GetSymbol() Symbol
	GetVolume() Volume
	GetPrice() Price
	GetSide() Side
}
