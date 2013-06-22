package bitstamp

import (
	"github.com/mattomatic/go-bitcoin/common"
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

func (o *OrderBook) GetBids() []common.Order {
	// just to speed things along
	bids := make([]common.Order, len(o.Bids))
	for i, b := range o.Bids {
		bids[i] = &b
	}
	return bids
}

func (o *OrderBook) GetAsks() []common.Order {
	// just to speed things along
	asks := make([]common.Order, len(o.Asks))
	for i, a := range o.Asks {
		asks[i] = &a
	}
	return asks
}
