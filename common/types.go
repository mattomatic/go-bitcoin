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
type Timestamp time.Time

type Trade interface {
	Exchange() Exchange
	Symbol() Symbol
	Timestamp() Timestamp
	Volume() Volume
	Price() Price
	Currency() Currency
	TradeId() TradeId
}
