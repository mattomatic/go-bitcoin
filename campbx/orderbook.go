package campbx

import (
	"github.com/mattomatic/go-bitcoin/common"
)

type OrderBook struct {
	Bids []Order `json:"Bids"`
	Asks []Order `json:"Asks"`
}

func newOrderBook() *OrderBook {
	return &OrderBook{make([]Order, 0), make([]Order, 0)}
}

func (o *OrderBook) GetBids() <-chan common.Order {
	return getChannelForward(o.Bids, common.Bid)
}

func (o *OrderBook) GetAsks() <-chan common.Order {
	// campbx asks are sorted highest to lowest for some reason.
	return getChannelReverse(o.Asks, common.Ask)
}

func getChannelForward(orders []Order, side common.Side) chan common.Order {
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

func getChannelReverse(orders []Order, side common.Side) chan common.Order {
	ch := make(chan common.Order)
	go func() {
		defer close(ch)
		for i := len(orders) - 1; i >= 0; i-- {
			orders[i].side = side
			ch <- &orders[i]
		}
	}()
	return ch
}
