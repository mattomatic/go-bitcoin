package bitstamp

import (
	"time"
)

type OrderBook struct {
	Timestamp string  `json:"timestamp"`
	Bids      []Order `json:"bids"`
	Asks      []Order `json:"asks"`
}

func (o *OrderBook) GetTimestamp() time.Time {
	return getTimestamp(o.Timestamp)
}
