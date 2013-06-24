package bitstamp

import (
	"github.com/mattomatic/go-bitcoin/common"
)

type OrderBook struct {
	Bids []Order `json:"bids"`
	Asks []Order `json:"asks"`
}

func newOrderBook() *OrderBook {
	return &OrderBook{make([]Order, 0), make([]Order, 0)}
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
