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
	Side     string
)

const (
	Bid Side = "Bid"
	Ask Side = "Ask"
)

const (
	USD Currency = "USD"
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
	GetSide() Side
}

type OrderBook interface {
	GetBids() <-chan Order
	GetAsks() <-chan Order
}

type Diff Order
