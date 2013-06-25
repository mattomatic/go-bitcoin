package btce

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
	return getChannel(o.Bids, common.Bid)
}

func (o *OrderBook) GetAsks() <-chan common.Order {
	return getChannel(o.Asks, common.Ask)
}

func getChannel(orders []Order, side common.Side) chan common.Order {
	ch := make(chan common.Order)
	go func() {
		defer close(ch)
		for i := 0; i < len(orders); i++ {
			orders[i].side = side
			ch <- &orders[i]
		}
	}()
	return ch
}
