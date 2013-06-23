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

func (o *OrderBook) GetBids() <-chan common.Order {
	return getChannel(o.Bids)
}

func (o *OrderBook) GetAsks() <-chan common.Order {
	return getChannel(o.Asks)
}

func getChannel(orders []Order) chan common.Order {
	ch := make(chan common.Order)
	go func() {
		defer close(ch)
		for i := 0; i < len(orders); i++ {
			ch <- &orders[i]
		}
	}()
	return ch
}
