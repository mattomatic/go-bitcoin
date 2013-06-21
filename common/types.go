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

type Trade interface {
	Exchange() Exchange
	Symbol() Symbol
	Timestamp() time.Time
	Volume() Volume
	Price() Price
	Currency() Currency
	TradeId() TradeId
}

type Ticker interface {
	Exchange() Exchange
	Symbol() Symbol
	Timestamp() time.Time
	Bid() Price
	Ask() Price
	High() Price
	Low() Price
	Volume() Volume
}
