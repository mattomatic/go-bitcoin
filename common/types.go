package common

import (
	"time"
)

type (
	Volume   float64
	Price    float64
	Fee      float64
	Amount   float64
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

type Description interface {
	GetExchange() Exchange
	GetSymbol() Symbol
}

type Trade interface {
	Description
	GetTimestamp() time.Time
	GetVolume() Volume
	GetPrice() Price
	GetTradeId() TradeId
}

type Ticker interface {
	GetExchange() Exchange
	GetSymbol() Symbol
	GetBid() Price
	GetAsk() Price
}

type Order interface {
	Description
	GetVolume() Volume
	GetPrice() Price
	GetFee() Fee
	GetSide() Side
}

type OrderBook interface {
	GetBids() <-chan Order
	GetAsks() <-chan Order
}

type Diff Order
