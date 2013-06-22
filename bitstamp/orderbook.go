package bitstamp

import (
	"github.com/mattomatic/go-bitcoin/common"
	"time"
)

type OrderBook struct {
	Timestamp string         `json:"timestamp"`
	Bids      []common.Order `json:"bids"`
	Asks      []common.Order `json:"asks"`
}

func (o *OrderBook) GetTimestamp() time.Time {
	return getTimestamp(o.Timestamp)
}

func (o *OrderBook) GetBids() []common.Order {
	return o.Bids
}

func (o *OrderBook) GetAsks() []common.Order {
	return o.Asks
}
